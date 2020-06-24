#include <cpplogger/cpplogger.h>
#include <mutex>
#include <shlwapi.h>
#include <windows.h>

#include "PocCom.h"
#include "util.h"

extern Logger::Logger *Log;
extern wchar_t *logServerAddr;

const TCHAR ProgIDStr[] = TEXT("ScreenReaderX.PocCom");
LONG LockCount{};
HINSTANCE PocComDLLInstance{};
TCHAR PocComCLSIDStr[256]{};
TCHAR LibraryIDStr[256]{};

// CPocCom
CPocCom::CPocCom() : mReferenceCount(1), mTypeInfo(nullptr) {
  ITypeLib *pTypeLib{};
  HRESULT hr{};

  LockModule(true);

  hr = LoadRegTypeLib(LIBID_PocComLib, 1, 0, 0, &pTypeLib);

  if (SUCCEEDED(hr)) {
    pTypeLib->GetTypeInfoOfGuid(IID_IPocCom, &mTypeInfo);
    pTypeLib->Release();
  }
}

CPocCom::~CPocCom() { LockModule(false); }

STDMETHODIMP CPocCom::QueryInterface(REFIID riid, void **ppvObject) {
  *ppvObject = nullptr;

  if (IsEqualIID(riid, IID_IUnknown) || IsEqualIID(riid, IID_IDispatch) ||
      IsEqualIID(riid, IID_IPocCom)) {
    *ppvObject = static_cast<IPocCom *>(this);
  } else {
    return E_NOINTERFACE;
  }

  AddRef();

  return S_OK;
}

STDMETHODIMP_(ULONG) CPocCom::AddRef() {
  return InterlockedIncrement(&mReferenceCount);
}

STDMETHODIMP_(ULONG) CPocCom::Release() {
  if (InterlockedDecrement(&mReferenceCount) == 0) {
    delete this;

    return 0;
  }

  return mReferenceCount;
}

STDMETHODIMP CPocCom::GetTypeInfoCount(UINT *pctinfo) {
  *pctinfo = mTypeInfo != nullptr ? 1 : 0;

  return S_OK;
}

STDMETHODIMP CPocCom::GetTypeInfo(UINT iTInfo, LCID lcid, ITypeInfo **ppTInfo) {
  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  mTypeInfo->AddRef();
  *ppTInfo = mTypeInfo;

  return S_OK;
}

STDMETHODIMP CPocCom::GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames,
                                    UINT cNames, LCID lcid, DISPID *rgDispId) {
  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  return mTypeInfo->GetIDsOfNames(rgszNames, cNames, rgDispId);
}

STDMETHODIMP CPocCom::Invoke(DISPID dispIdMember, REFIID riid, LCID lcid,
                             WORD wFlags, DISPPARAMS *pDispParams,
                             VARIANT *pVarResult, EXCEPINFO *pExcepInfo,
                             UINT *puArgErr) {
  void *p = static_cast<IPocCom *>(this);

  if (mTypeInfo == nullptr) {
    return E_NOTIMPL;
  }

  return mTypeInfo->Invoke(p, dispIdMember, wFlags, pDispParams, pVarResult,
                           pExcepInfo, puArgErr);
}

STDMETHODIMP CPocCom::Start(LPWSTR logServerAddr, LOGLEVEL level) {
  std::lock_guard<std::mutex> lock(mMutex);

  if (mIsActive) {
    Log->Warn(L"IPocCom::Start() is already called", GetCurrentThreadId(),
              __LONGFILE__);

    return E_FAIL;
  }

  mIsActive = true;

  // Assumes that the `logServerAddr` is form of "hostname:port".
  LogServerAddr = new wchar_t[128]{};

  HRESULT hr =
      StringCbPrintfW(LogServerAddr, 256, L"http://%s/v1/log", logServerAddr);

  if (FAILED(hr)) {
    return E_FAIL;
  }

  Log = new Logger::Logger(L"PocCom", L"v0.1.0-develop", 4096);

  Log->Info(L"Called IPocCom::Start()", GetCurrentThreadId(), __LONGFILE__);

  mLoggingCtx = new LoggingContext();

  mLoggingCtx->QuitEvent =
      CreateEventEx(nullptr, nullptr, 0, EVENT_MODIFY_STATE | SYNCHRONIZE);

  if (mLoggingCtx->QuitEvent == nullptr) {
    Log->Fail(L"Failed to create event", GetCurrentThreadId(), __LONGFILE__);

    return E_FAIL;
  }

  Log->Info(L"Create logging thread", GetCurrentThreadId(), __LONGFILE__);

  mLoggingThread = CreateThread(nullptr, 0, loggingThread,
                                static_cast<void *>(mLoggingCtx), 0, nullptr);

  if (mLoggingThread == nullptr) {
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

STDMETHODIMP CPocCom::Stop() {
  std::lock_guard<std::mutex> lock(mMutex);

  if (!mIsActive) {
    return E_FAIL;
  }

  Log->Info(L"Called IPocCom::Stop()", GetCurrentThreadId(), __LONGFILE__);

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

  if (mLoggingThread == nullptr) {
    goto END_LOGLOOP_CLEANUP;
  }
  if (!SetEvent(mLoggingCtx->QuitEvent)) {
    Log->Fail(L"Failed to send event", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  WaitForSingleObject(mLoggingThread, INFINITE);
  SafeCloseHandle(&mLoggingThread);
  SafeCloseHandle(&(mLoggingCtx->QuitEvent));

  delete mLoggingCtx;
  mLoggingCtx = nullptr;

END_LOGLOOP_CLEANUP:

  mIsActive = false;

  return S_OK;
}

STDMETHODIMP CPocCom::SetUIAEventHandler(IUIEventHandler handleFunc) {
  std::lock_guard<std::mutex> lock(mMutex);

  mAutomationCtx->IUIEventHandleFunc = handleFunc;

  return S_OK;
}

STDMETHODIMP CPocCom::SetMSAAEventHandler(IAEventHandler handleFunc) {
  std::lock_guard<std::mutex> lock(mMutex);

  mAutomationCtx->IAEventHandleFunc = handleFunc;

  return S_OK;
}

STDMETHODIMP
CPocCom::GetIUIAutomationElement(TreeWalkerDirection direction,
                                 IUIAutomationElement *pRootElement,
                                 IUIAutomationElement **ppElement) {
  std::lock_guard<std::mutex> lock(mMutex);

  if (ppElement == nullptr || mAutomationCtx->BaseTreeWalker == nullptr) {
    return E_FAIL;
  }

  Log->Info(L"Called IPocCom::GetIUIAutomationElement()", GetCurrentThreadId(),
            __LONGFILE__);

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
CPocCom::UpdateTreeWalker() {
  std::lock_guard<std::mutex> lock(mMutex);

  Log->Info(L"Called IPocCom::UpdateTreeWalker()", GetCurrentThreadId(),
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
CPocCom::ElementFromHandle(UIA_HWND hwnd, IUIAutomationElement **ppElement) {
  std::lock_guard<std::mutex> lock(mMutex);

  Log->Info(L"Called IPocCom::ElementFromHandle()", GetCurrentThreadId(),
            __LONGFILE__);

  HRESULT hr{};

  hr = mAutomationCtx->UIAutomation->ElementFromHandle(hwnd, ppElement);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call IUIAutomation::ElementFromHandle",
              GetCurrentThreadId(), __LONGFILE__);
  }

  return S_OK;
}

// CPocComFactory
STDMETHODIMP CPocComFactory::QueryInterface(REFIID riid, void **ppvObject) {
  *ppvObject = nullptr;

  if (IsEqualIID(riid, IID_IUnknown) || IsEqualIID(riid, IID_IClassFactory)) {
    *ppvObject = static_cast<IClassFactory *>(this);
  } else {
    return E_NOINTERFACE;
  }

  AddRef();

  return S_OK;
}

STDMETHODIMP_(ULONG) CPocComFactory::AddRef() {
  LockModule(true);

  return 2;
}

STDMETHODIMP_(ULONG) CPocComFactory::Release() {
  LockModule(false);

  return 1;
}

STDMETHODIMP CPocComFactory::CreateInstance(IUnknown *pUnkOuter, REFIID riid,
                                            void **ppvObject) {
  CPocCom *p{};
  HRESULT hr{};

  *ppvObject = nullptr;

  if (pUnkOuter != nullptr) {
    return CLASS_E_NOAGGREGATION;
  }

  p = new CPocCom();

  if (p == nullptr) {
    return E_OUTOFMEMORY;
  }

  hr = p->QueryInterface(riid, ppvObject);
  p->Release();

  return hr;
}

STDMETHODIMP CPocComFactory::LockServer(BOOL fLock) {
  LockModule(fLock);

  return S_OK;
}

// DLL Export
STDAPI DllCanUnloadNow(void) { return LockCount == 0 ? S_OK : S_FALSE; }

STDAPI DllGetClassObject(REFCLSID rclsid, REFIID riid, LPVOID *ppv) {
  static CPocComFactory serverFactory;
  HRESULT hr{};
  *ppv = nullptr;

  if (IsEqualCLSID(rclsid, CLSID_PocCom)) {
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

  wsprintf(szKey, TEXT("CLSID\\%s"), PocComCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         TEXT("ScreenReaderX.PocCom"))) {
    return E_FAIL;
  }

  GetModuleFileName(PocComDLLInstance, szModulePath,
                    sizeof(szModulePath) / sizeof(TCHAR));
  wsprintf(szKey, TEXT("CLSID\\%s\\InprocServer32"), PocComCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr, szModulePath)) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("CLSID\\%s\\InprocServer32"), PocComCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, TEXT("ThreadingModel"),
                         TEXT("Apartment"))) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("CLSID\\%s\\ProgID"), PocComCLSIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         (LPTSTR)ProgIDStr)) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("%s"), ProgIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         TEXT("ScreenReaderX.PocCom"))) {
    return E_FAIL;
  }

  wsprintf(szKey, TEXT("%s\\CLSID"), ProgIDStr);

  if (!CreateRegistryKey(HKEY_CLASSES_ROOT, szKey, nullptr,
                         (LPTSTR)PocComCLSIDStr)) {
    return E_FAIL;
  }

  GetModuleFileNameW(PocComDLLInstance, szTypeLibPath,
                     sizeof(szTypeLibPath) / sizeof(TCHAR));

  hr = LoadTypeLib(szTypeLibPath, &pTypeLib);

  if (SUCCEEDED(hr)) {
    hr = RegisterTypeLib(pTypeLib, szTypeLibPath, nullptr);

    if (SUCCEEDED(hr)) {
      wsprintf(szKey, TEXT("CLSID\\%s\\TypeLib"), PocComCLSIDStr);

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

  wsprintf(szKey, TEXT("CLSID\\%s"), PocComCLSIDStr);
  SHDeleteKey(HKEY_CLASSES_ROOT, szKey);

  wsprintf(szKey, TEXT("%s"), ProgIDStr);
  SHDeleteKey(HKEY_CLASSES_ROOT, szKey);

  UnRegisterTypeLib(LIBID_PocComLib, 1, 0, LOCALE_NEUTRAL, SYS_WIN32);

  return S_OK;
}

BOOL WINAPI DllMain(HINSTANCE hInstance, DWORD dwReason, LPVOID lpReserved) {
  switch (dwReason) {
  case DLL_PROCESS_ATTACH:
    PocComDLLInstance = hInstance;
    DisableThreadLibraryCalls(hInstance);
    GetGuidString(CLSID_PocCom, PocComCLSIDStr);
    GetGuidString(LIBID_PocComLib, LibraryIDStr);

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
