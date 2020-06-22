#include <UIAutomationClient.h>
#include <UIAutomationCore.h>
#include <cpplogger/cpplogger.h>
#include <oleauto.h>

#include <strsafe.h>

#include "types.h"
#include "uiahandler.h"
#include "util.h"

extern Logger::Logger *Log;

FocusChangeEventHandler::FocusChangeEventHandler(AutomationContext *ctx)
    : mAutomationCtx(ctx), mLeft(0), mTop(0), mRight(0), mBottom(0) {}

ULONG FocusChangeEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG FocusChangeEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT FocusChangeEventHandler::QueryInterface(REFIID riid,
                                                void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface = static_cast<IUIAutomationFocusChangedEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationFocusChangedEventHandler))
    *ppInterface = static_cast<IUIAutomationFocusChangedEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
FocusChangeEventHandler::HandleFocusChangedEvent(
    IUIAutomationElement *pSender) {
  HRESULT hr{};

  if (pSender == nullptr) {
    return S_OK;
  }

  Log->Info(L"Called HandleFocusChangedEvent()", GetCurrentThreadId(),
            __LONGFILE__);

  RECT rectangle{};

  hr = pSender->get_CurrentBoundingRectangle(&rectangle);

  if (FAILED(hr)) {
    return hr;
  }

  LONG l = mLeft;
  LONG t = mTop;
  LONG r = mRight;
  LONG b = mBottom;

  mLeft = rectangle.left;
  mTop = rectangle.top;
  mRight = rectangle.right;
  mBottom = rectangle.bottom;

  bool isSameBoundingRectangle = l == rectangle.left && t == rectangle.top &&
                                 r == rectangle.right && b == rectangle.bottom;

  if (isSameBoundingRectangle) {
    return S_OK;
  }
  if (mAutomationCtx != nullptr &&
      mAutomationCtx->IUIEventHandleFunc != nullptr) {
    mAutomationCtx->IUIEventHandleFunc(UIA_AutomationFocusChangedEventId,
                                       pSender);
  }

  SAFEARRAY *runtimeId{};

  hr = pSender->GetRuntimeId(&runtimeId);

  if (FAILED(hr)) {
    Log->Warn(L"Failed to call IUIAutomationElement::GetRuntimeId()",
              GetCurrentThreadId(), __LONGFILE__);
    return hr;
  }
  if (runtimeId == nullptr) {
    return S_OK;
  }

  hr = SafeArrayCopy(runtimeId, &(mAutomationCtx->FocusElementRuntimeId));

  if (FAILED(hr)) {
    Log->Warn(L"Failed to call SafeArrayCopy()", GetCurrentThreadId(),
              __LONGFILE__);
    return hr;
  }
  if (!SetEvent(mAutomationCtx->FocusEvent)) {
    Log->Fail(L"Failed to send event", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  return S_OK;
}

PropertyChangeEventHandler::PropertyChangeEventHandler(AutomationContext *ctx)
    : mAutomationCtx(ctx) {}

ULONG PropertyChangeEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG PropertyChangeEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT PropertyChangeEventHandler::QueryInterface(REFIID riid,
                                                   void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface =
        static_cast<IUIAutomationPropertyChangedEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationPropertyChangedEventHandler))
    *ppInterface =
        static_cast<IUIAutomationPropertyChangedEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
PropertyChangeEventHandler::HandlePropertyChangedEvent(
    IUIAutomationElement *pSender, PROPERTYID propertyId, VARIANT newValue) {
  if (pSender == nullptr) {
    return S_OK;
  }

  Log->Info(L"Called HandlePropertyChangedEvent()", GetCurrentThreadId(),
            __LONGFILE__);

  if (mAutomationCtx != nullptr &&
      mAutomationCtx->IUIEventHandleFunc != nullptr) {
    mAutomationCtx->IUIEventHandleFunc(UIA_AutomationPropertyChangedEventId,
                                       pSender);
  }

  return S_OK;
}

AutomationEventHandler::AutomationEventHandler(AutomationContext *ctx)
    : mAutomationCtx(ctx) {}

ULONG AutomationEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG AutomationEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT AutomationEventHandler::QueryInterface(REFIID riid,
                                               void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface = static_cast<IUIAutomationEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationEventHandler))
    *ppInterface = static_cast<IUIAutomationEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
AutomationEventHandler::HandleAutomationEvent(IUIAutomationElement *pSender,
                                              EVENTID eventId) {
  if (pSender == nullptr) {
    return S_OK;
  }

  Log->Info(L"Called HandleAutomationEvent()", GetCurrentThreadId(),
            __LONGFILE__);

  if (mAutomationCtx != nullptr &&
      mAutomationCtx->IUIEventHandleFunc != nullptr) {
    mAutomationCtx->IUIEventHandleFunc(static_cast<INT64>(eventId), pSender);
  }

  return S_OK;
}

StructureChangeEventHandler::StructureChangeEventHandler(AutomationContext *ctx)
    : mAutomationCtx(ctx) {}

ULONG StructureChangeEventHandler::AddRef() {
  ULONG ret = InterlockedIncrement(&mRefCount);
  return ret;
}

ULONG StructureChangeEventHandler::Release() {
  ULONG ret = InterlockedDecrement(&mRefCount);
  if (ret == 0) {
    delete this;
    return 0;
  }
  return ret;
}

HRESULT StructureChangeEventHandler::QueryInterface(REFIID riid,
                                                    void **ppInterface) {
  if (riid == __uuidof(IUnknown))
    *ppInterface =
        static_cast<IUIAutomationStructureChangedEventHandler *>(this);
  else if (riid == __uuidof(IUIAutomationStructureChangedEventHandler))
    *ppInterface =
        static_cast<IUIAutomationStructureChangedEventHandler *>(this);
  else {
    *ppInterface = nullptr;
    return E_NOINTERFACE;
  }
  this->AddRef();
  return S_OK;
}

HRESULT
StructureChangeEventHandler::HandleStructureChangedEvent(
    IUIAutomationElement *pSender, enum StructureChangeType changeType,
    SAFEARRAY *runtimeId) {
  if (pSender == nullptr) {
    return S_OK;
  }

  Log->Info(L"Called HandleStructureChangedEvent()", GetCurrentThreadId(),
            __LONGFILE__);

  if (mAutomationCtx != nullptr &&
      mAutomationCtx->IUIEventHandleFunc != nullptr) {
    mAutomationCtx->IUIEventHandleFunc(0, pSender);
  }

  return S_OK;
}
