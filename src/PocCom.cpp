#include <cpplogger/cpplogger.h>
#include <mutex>
#include <shlwapi.h>
#include <windows.h>

#include "CoreServer.h"
#include "util.h"

extern Logger::Logger *Log;

const TCHAR ProgIDStr[] = TEXT("ScreenReaderX.CoreServer");
LONG LockCount{};
HINSTANCE CoreServerDLLInstance{};
TCHAR CoreServerCLSIDStr[256]{};
TCHAR LibraryIDStr[256]{};

// CCoreServer
CCoreServer::CCoreServer() : mReferenceCount(1), mTypeInfo(nullptr) {
  ITypeLib *pTypeLib{};
  HRESULT hr{};

  LockModule(true);

  hr = LoadRegTypeLib(LIBID_CoreServerLib, 1, 0, 0, &pTypeLib);

  if (SUCCEEDED(hr)) {
    pTypeLib->GetTypeInfoOfGuid(IID_ICoreServer, &mTypeInfo);
    pTypeLib->Release();
  }
}

CCoreServer::~CCoreServer() { LockModule(false); }

STDMETHODIMP CCoreServer::QueryInterface(REFIID riid, void **ppvObject) {
  *ppvObject = nullptr;

  if (IsEqualIID(riid, IID_IUnknown) || IsEqualIID(riid, IID_IDispatch) ||
      IsEqualIID(riid, IID_ICoreServer)) {
    *ppvObject = static_cast<ICoreServer *>(this);
  } else {
    return E_NOINTERFACE;
  }

  AddRef();

  return S_OK;
}

STDMETHODIMP_(ULONG) CCoreServer::AddRef() {
  return InterlockedIncrement(&mReferenceCount);
}

STDMETHODIMP_(ULONG) CCoreServer::Release() {
  if (InterlockedDecrement(&mReferenceCount) == 0) {
    delete this;

    return 0;
  }

  return mReferenceCount;
}

STDMETHODIMP CCoreServer::GetTypeInfoCount(UINT *pctinfo) {
  *pctinfo = mTypeInfo != nullptr ? 1 : 0;

  return S_OK;
}

STDMETHODIMP CCoreServer::GetTypeInfo(UINT iTInfo, LCID lcid,
                                      ITypeInfo **ppTInfo) {
  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  mTypeInfo->AddRef();
  *ppTInfo = mTypeInfo;

  return S_OK;
}

STDMETHODIMP CCoreServer::GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames,
                                        UINT cNames, LCID lcid,
                                        DISPID *rgDispId) {
  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  return mTypeInfo->GetIDsOfNames(rgszNames, cNames, rgDispId);
}

STDMETHODIMP CCoreServer::Invoke(DISPID dispIdMember, REFIID riid, LCID lcid,
                                 WORD wFlags, DISPPARAMS *pDispParams,
                                 VARIANT *pVarResult, EXCEPINFO *pExcepInfo,
                                 UINT *puArgErr) {
  void *p = static_cast<ICoreServer *>(this);

  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  return mTypeInfo->Invoke(p, dispIdMember, wFlags, pDispParams, pVarResult,
                           pExcepInfo, puArgErr);
}

STDMETHODIMP CCoreServer::Start() {
  std::lock_guard<std::mutex> lock(mMutex);

  if (mIsActive) {
    Log->Warn(L"ICoreServer::Start() is already called", GetCurrentThreadId(),
              __LONGFILE__);

    return E_FAIL;
  }

  mIsActive = true;

  Log = new Logger::Logger(L"CoreServer", L"v0.1.0-develop", 4096);

  Log->Info(L"Called ICoreServer::Start()", GetCurrentThreadId(), __LONGFILE__);

  mLogLoopCtx = new LogLoopContext();

  mLogLoopCtx->QuitEvent =
      CreateEventEx(nullptr, nullptr, 0, EVENT_MODIFY_STATE | SYNCHRONIZE);

  if (mLogLoopCtx->QuitEvent == nullptr) {
    Log->Fail(L"Failed to create event", GetCurrentThreadId(), __LONGFILE__);

    return E_FAIL;
  }

  Log->Info(L"Create log loop thread", GetCurrentThreadId(), __LONGFILE__);

  mLogLoopThread = CreateThread(nullptr, 0, logLoop,
                                static_cast<void *>(mLogLoopCtx), 0, nullptr);

  if (mLogLoopThread == nullptr) {
    Log->Fail(L"Failed to create thread", GetCurrentThreadId(), __LONGFILE__);

    return E_FAIL;
  }

  mAutomationCtx = new AutomationContext();

  mAutomationCtx->QuitEvent =
      CreateEventEx(nullptr, nullptr, 0, EVENT_MODIFY_STATE | SYNCHRONIZE);

  if (mAutomationCtx->QuitEvent == nullptr) {
    Log->Fail(L"Failed to create event", GetCurrentThreadId(), __LONGFILE__);

    return E_FAIL;
  }

  mAutomationCtx->FocusEvent =
      CreateEventEx(nullptr, nullptr, 0, EVENT_MODIFY_STATE | SYNCHRONIZE);

  if (mAutomationCtx->FocusEvent == nullptr) {
    Log->Fail(L"Failed to create event", GetCurrentThreadId(), __LONGFILE__);

    return E_FAIL;
  }

  Log->Info(L"Create UI automation thread", GetCurrentThreadId(), __LONGFILE__);

  mUIAThread = CreateThread(nullptr, 0, uiaThread,
                            static_cast<void *>(mAutomationCtx), 0, nullptr);

  if (mUIAThread == nullptr) {
    Log->Fail(L"Failed to create thread", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  Log->Info(L"Create Windows event thread", GetCurrentThreadId(), __LONGFILE__);

  mWindowsEventThread =
      CreateThread(nullptr, 0, windowsEventThread,
                   static_cast<void *>(mAutomationCtx), 0, nullptr);

  if (mWindowsEventThread == nullptr) {
    Log->Fail(L"Failed to create thread", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  return S_OK;
}

STDMETHODIMP CCoreServer::Stop() {
  std::lock_guard<std::mutex> lock(mMutex);

  if (!mIsActive) {
    return E_FAIL;
  }

  Log->Info(L"Called ICoreServer::Stop()", GetCurrentThreadId(), __LONGFILE__);

  if (mUIAThread == nullptr) {
    goto END_UIA_CLEANUP;
  }
  if (!SetEvent(mAutomationCtx->QuitEvent)) {
    Log->Fail(L"Failed to send event", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  WaitForSingleObject(mUIAThread, INFINITE);
  SafeCloseHandle(&mUIAThread);
  SafeCloseHandle(&(mAutomationCtx->QuitEvent));

  Log->Info(L"Delete UI automation thread", GetCurrentThreadId(), __LONGFILE__);

END_UIA_CLEANUP:

  if (mWindowsEventThread == nullptr) {
    goto END_WINDOWS_EVENT_CLEANUP;
  }

  mAutomationCtx->IsActive = false;
  WaitForSingleObject(mWindowsEventThread, INFINITE);
  SafeCloseHandle(&mWindowsEventThread);

  delete mAutomationCtx;
  mAutomationCtx = nullptr;

  Log->Info(L"Delete Windows event thread", GetCurrentThreadId(), __LONGFILE__);

END_WINDOWS_EVENT_CLEANUP:

  if (mLogLoopThread == nullptr) {
    goto END_LOGLOOP_CLEANUP;
  }
  if (!SetEvent(mLogLoopCtx->QuitEvent)) {
    Log->Fail(L"Failed to send event", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  WaitForSingleObject(mLogLoopThread, INFINITE);
  SafeCloseHandle(&mLogLoopThread);
  SafeCloseHandle(&(mLogLoopCtx->QuitEvent));

  delete mLogLoopCtx;
  mLogLoopCtx = nullptr;

END_LOGLOOP_CLEANUP:

  mIsActive = false;

  return S_OK;
}

STDMETHODIMP CCoreServer::SetUIAEventHandler(IUIEventHandler handleFunc) {
  std::lock_guard<std::mutex> lock(mMutex);

  mAutomationCtx->IUIEventHandleFunc = handleFunc;

  return S_OK;
}

STDMETHODIMP CCoreServer::SetMSAAEventHandler(IAEventHandler handleFunc) {
  std::lock_guard<std::mutex> lock(mMutex);

  mAutomationCtx->IAEventHandleFunc = handleFunc;

  return S_OK;
}

STDMETHODIMP
CCoreServer::GetIUIAutomationElement(TreeWalkerDirection direction,
                                     IUIAutomationElement *pRootElement,
                                     IUIAutomationElement **ppElement) {
  std::lock_guard<std::mutex> lock(mMutex);

  if (ppElement == nullptr || mAutomationCtx->BaseTreeWalker == nullptr) {
    return E_FAIL;
  }

  Log->Info(L"Called ICoreServer::GetIUIAutomationElement()",
            GetCurrentThreadId(), __LONGFILE__);

  HRESULT hr{};

  switch (direction) {
  case TW_NEXT:
    hr = mAutomationCtx->BaseTreeWalker->GetNextSiblingElement(pRootElement,
                                                               ppElement);
    break;
  case TW_PREVIOUS:
    hr = mAutomationCtx->BaseTreeWalker->GetPreviousSiblingElement(pRootElement,
                                                                   ppElement);
    break;
  case TW_FIRST_CHILD:
    hr = mAutomationCtx->BaseTreeWalker->GetFirstChildElement(pRootElement,
                                                              ppElement);
    break;
  case TW_LAST_CHILD:
    hr = mAutomationCtx->BaseTreeWalker->GetLastChildElement(pRootElement,
                                                             ppElement);
    break;
  case TW_PARENT:
    hr = mAutomationCtx->BaseTreeWalker->GetParentElement(pRootElement,
                                                          ppElement);
    break;
  default:
    return S_OK;
  }

  return S_OK;
}

STDMETHODIMP
CCoreServer::UpdateTreeWalker() {
  std::lock_guard<std::mutex> lock(mMutex);

  Log->Info(L"Called ICoreServer::UpdateTreeWalker()", GetCurrentThreadId(),
            __LONGFILE__);

  SafeRelease(&(mAutomationCtx->BaseTreeWalker));

  HRESULT hr{};

  hr = mAutomationCtx->UIAutomation->get_RawViewWalker(
      &(mAutomationCtx->BaseTreeWalker));

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call IUIAutomation::get_RawViewWalker",
              GetCurrentThreadId(), __LONGFILE__);
  }

  return S_OK;
}

STDMETHODIMP
CCoreServer::ElementFromHandle(UIA_HWND hwnd,
                               IUIAutomationElement **ppElement) {
  std::lock_guard<std::mutex> lock(mMutex);

  Log->Info(L"Called ICoreServer::ElementFromHandle()", GetCurrentThreadId(),
            __LONGFILE__);

  HRESULT hr{};

  hr = mAutomationCtx->UIAutomation->ElementFromHandle(hwnd, ppElement);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call IUIAutomation::ElementFromHandle",
              GetCurrentThreadId(), __LONGFILE__);
  }

  return S_OK;
}

// CCoreServerFactory
STDMETHODIMP CCoreServerFactory::QueryInterface(REFIID riid, void **ppvObject) {
  *ppvObject = nullptr;

  if (IsEqualIID(riid, IID_IUnknown) || IsEqualIID(riid, IID_IClassFactory)) {
    *ppvObject = static_cast<IClassFactory *>(this);
  } else {
    return E_NOINTERFACE;
  }

  AddRef();

  return S_OK;
}

STDMETHODIMP_(ULONG) CCoreServerFactory::AddRef() {
  LockModule(true);

  return 2;
}

STDMETHODIMP_(ULONG) CCoreServerFactory::Release() {
  LockModule(false);

  return 1;
}

STDMETHODIMP CCoreServerFactory::CreateInstance(IUnknown *pUnkOuter,
                                                REFIID riid, void **ppvObject) {
  CCoreServer *p{};
  HRESULT hr{};

  *ppvObject = nullptr;

  if (pUnkOuter != nullptr) {
    return CLASS_E_NOAGGREGATION;
  }

  p = new CCoreServer();

  if (p == nullptr) {
    return E_OUTOFMEMORY;
  }

  hr = p->QueryInterface(riid, ppvObject);
  p->Release();

  return hr;
}

STDMETHODIMP CCoreServerFactory::LockServer(BOOL fLock) {
  LockModule(fLock);

  return S_OK;
}

// DLL Export
STDAPI DllCanUnloadNow(void) { return LockCount == 0 ? S_OK : S_FALSE; }

STDAPI DllGetClassObject(REFCLSID rclsid, REFIID riid, LPVOID *ppv) {
  static CCoreServerFactory serverFactory;
  HRESULT hr{};
  *ppv = nullptr;

  if (IsEqualCLSID(rclsid, CLSID_CoreServer)) {
    hr = serverFactory.QueryInterface(riid, ppv);
  } else {
    hr = CLASS_E_CLASSNOTAVAILABLE;
  }

  return hr;
}

STDAPI DllRegisterServer(void) {
  TCHAR szModulePath[256]{};
  TCHAR szKey[256]{};
  WCHAR szTypeLibPath[256]{};
  HRESULT hr{};
  ITypeLib *pTypeLib{};

  wsprintf(szKey, TEXT("CLSID\\%s"), CoreServerCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         TEXT("ScreenReaderX.CoreServer"))) {
    return E_FAIL;
  }

  GetModuleFileName(CoreServerDLLInstance, szModulePath,
                    sizeof(szModulePath) / sizeof(TCHAR));
  wsprintf(szKey, TEXT("CLSID\\%s\\InprocServer32"), CoreServerCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr, szModulePath)) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("CLSID\\%s\\InprocServer32"), CoreServerCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, TEXT("ThreadingModel"),
                         TEXT("Apartment"))) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("CLSID\\%s\\ProgID"), CoreServerCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         (LPTSTR)ProgIDStr)) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("%s"), ProgIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         TEXT("ScreenReaderX.CoreServer"))) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("%s\\CLSID"), ProgIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         (LPTSTR)CoreServerCLSIDStr)) {
    return E_FAIL;
  }

  GetModuleFileNameW(CoreServerDLLInstance, szTypeLibPath,
                     sizeof(szTypeLibPath) / sizeof(TCHAR));

  hr = LoadTypeLib(szTypeLibPath, &pTypeLib);

  if (SUCCEEDED(hr)) {
    hr = RegisterTypeLib(pTypeLib, szTypeLibPath, nullptr);

    if (SUCCEEDED(hr)) {
      wsprintf(szKey, TEXT("CLSID\\%s\\TypeLib"), CoreServerCLSIDStr);

      if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr, LibraryIDStr)) {
        pTypeLib->Release();

        return E_FAIL;
      }
    }

    pTypeLib->Release();
  }

  return S_OK;
}

STDAPI DllUnregisterServer(void) {
  TCHAR szKey[256]{};

  wsprintf(szKey, TEXT("CLSID\\%s"), CoreServerCLSIDStr);
  SHDeleteKey(HKEY_CLASSES_ROOT, szKey);

  wsprintf(szKey, TEXT("%s"), ProgIDStr);
  SHDeleteKey(HKEY_CLASSES_ROOT, szKey);

  UnRegisterTypeLib(LIBID_CoreServerLib, 1, 0, LOCALE_NEUTRAL, SYS_WIN32);

  return S_OK;
}

BOOL WINAPI DllMain(HINSTANCE hInstance, DWORD dwReason, LPVOID lpReserved) {
  switch (dwReason) {
  case DLL_PROCESS_ATTACH:
    CoreServerDLLInstance = hInstance;
    DisableThreadLibraryCalls(hInstance);
    GetGuidString(CLSID_CoreServer, CoreServerCLSIDStr);
    GetGuidString(LIBID_CoreServerLib, LibraryIDStr);

    return TRUE;
  }

  return TRUE;
}

// Helper function
void LockModule(BOOL bLock) {
  if (bLock) {
    InterlockedIncrement(&LockCount);
  } else {
    InterlockedDecrement(&LockCount);
  }
}

BOOL CreateRegistryKey(HKEY hKeyRoot, LPTSTR lpszKey, LPTSTR lpszValue,
                       LPTSTR lpszData) {
  HKEY hKey{};
  LONG lResult{};
  DWORD dwSize{};

  lResult =
      RegCreateKeyEx(hKeyRoot, lpszKey, 0, nullptr, REG_OPTION_NON_VOLATILE,
                     KEY_WRITE, nullptr, &hKey, nullptr);

  if (lResult != ERROR_SUCCESS) {
    return FALSE;
  }
  if (lpszData != nullptr) {
    dwSize = (lstrlen(lpszData) + 1) * sizeof(TCHAR);
  } else {
    dwSize = 0;
  }

  RegSetValueEx(hKey, lpszValue, 0, REG_SZ, (LPBYTE)lpszData, dwSize);
  RegCloseKey(hKey);

  return TRUE;
}

void GetGuidString(REFGUID rguid, LPTSTR lpszGuid) {
  LPWSTR lpsz;
  StringFromCLSID(rguid, &lpsz);

#ifdef UNICODE
  lstrcpyW(lpszGuid, lpsz);
#else
  WideCharToMultiByte(CP_ACP, 0, lpsz, -1, lpszGuid, 256, nullptr, nullptr);
#endif

  CoTaskMemFree(lpsz);
}
