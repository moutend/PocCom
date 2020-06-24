#pragma once

#include <IPocCom/IPocCom.h>
#include <cpplogger/cpplogger.h>
#include <mutex>
#include <windows.h>

#include "context.h"
#include "loggingthread.h"
#include "uiathread.h"
#include "wineventthread.h"

#include <UIAutomationClient.h>

class CPocCom : public IPocCom {
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

  // IPocCom methods
  STDMETHODIMP Start(LPWSTR loggerURL, LOGLEVEL level);
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

  CPocCom();
  ~CPocCom();

private:
  LONG mReferenceCount;
  ITypeInfo *mTypeInfo;

  bool mIsActive = false;
  std::mutex mMutex;

  LoggingContext *mLoggingCtx = nullptr;
  AutomationContext *mAutomationCtx = nullptr;

  HANDLE mLoggingThread = nullptr;
  HANDLE mUIAThread = nullptr;
  HANDLE mWindowsEventThread = nullptr;
};

class CPocComFactory : public IClassFactory {
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
