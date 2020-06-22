#pragma once

#include <ICoreServer/ICoreServer.h>
#include <cpplogger/cpplogger.h>
#include <mutex>
#include <windows.h>

#include "context.h"
#include "logloop.h"
#include "uialoop.h"
#include "wineventloop.h"

#include <UIAutomationClient.h>

class CCoreServer : public ICoreServer {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IDispatch methods
  STDMETHODIMP GetTypeInfoCount(UINT *pctinfo);
  STDMETHODIMP GetTypeInfo(UINT iTInfo, LCID lcid, ITypeInfo **ppTInfo);
  STDMETHODIMP GetIDsOfNames(REFIID riid, LPOLESTR *rgszNames, UINT cNames,
                             LCID lcid, DISPID *rgDispId);
  STDMETHODIMP Invoke(DISPID dispIdMember, REFIID riid, LCID lcid, WORD wFlags,
                      DISPPARAMS *pDispParams, VARIANT *pVarResult,
                      EXCEPINFO *pExcepInfo, UINT *puArgErr);

  // ICoreServer methods
  STDMETHODIMP Start();
  STDMETHODIMP Stop();
  STDMETHODIMP SetMSAAEventHandler(IAEventHandler handleFunc);
  STDMETHODIMP SetUIAEventHandler(IUIEventHandler handleFunc);
  STDMETHODIMP
  GetIUIAutomationElement(TreeWalkerDirection direction,
                          IUIAutomationElement *pRootElement,
                          IUIAutomationElement **ppElement);
  STDMETHODIMP UpdateTreeWalker();
  STDMETHODIMP ElementFromHandle(UIA_HWND hwnd,
                                 IUIAutomationElement **ppElement);

  CCoreServer();
  ~CCoreServer();

private:
  LONG mReferenceCount;
  ITypeInfo *mTypeInfo;

  bool mIsActive = false;
  std::mutex mMutex;

  LogLoopContext *mLogLoopCtx = nullptr;
  AutomationContext *mAutomationCtx = nullptr;

  HANDLE mLogLoopThread = nullptr;
  HANDLE mUIAThread = nullptr;
  HANDLE mWindowsEventThread = nullptr;
};

class CCoreServerFactory : public IClassFactory {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IClassFactory methods
  STDMETHODIMP CreateInstance(IUnknown *pUnkOuter, REFIID riid,
                              void **ppvObject);
  STDMETHODIMP LockServer(BOOL fLock);
};

extern void LockModule(BOOL bLock);
extern BOOL CreateRegistryKey(HKEY hKeyRoot, LPTSTR lpszKey, LPTSTR lpszValue,
                              LPTSTR lpszData);
extern void GetGuidString(REFGUID rguid, LPTSTR lpszGuid);
