#pragma once

#include <UIAutomationClient.h>
#include <oleauto.h>
#include <windows.h>

#include "context.h"

class FocusChangeEventHandler : public IUIAutomationFocusChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationFocusChangedEventHandler methods
  STDMETHODIMP HandleFocusChangedEvent(IUIAutomationElement *pSender);

  FocusChangeEventHandler(AutomationContext *ctx);

private:
  LONG mRefCount;
  AutomationContext *mAutomationCtx;
  LONG mLeft;
  LONG mTop;
  LONG mRight;
  LONG mBottom;
};

class PropertyChangeEventHandler
    : public IUIAutomationPropertyChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationPropertyChangedEventHandler methods
  STDMETHODIMP HandlePropertyChangedEvent(IUIAutomationElement *sender,
                                          PROPERTYID propertyId,
                                          VARIANT newValue);

  PropertyChangeEventHandler(AutomationContext *ctx);

private:
  LONG mRefCount;
  AutomationContext *mAutomationCtx;
};

class AutomationEventHandler : public IUIAutomationEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationEventHandler methods
  STDMETHODIMP HandleAutomationEvent(IUIAutomationElement *sender,
                                     EVENTID eventId);

  AutomationEventHandler(AutomationContext *ctx);

private:
  LONG mRefCount;
  AutomationContext *mAutomationCtx;
};

class StructureChangeEventHandler
    : public IUIAutomationStructureChangedEventHandler {
public:
  // IUnknown methods
  STDMETHODIMP QueryInterface(REFIID riid, void **ppvObject);
  STDMETHODIMP_(ULONG) AddRef();
  STDMETHODIMP_(ULONG) Release();

  // IUIAutomationStructureChangedEventHandler methods
  STDMETHODIMP HandleStructureChangedEvent(IUIAutomationElement *pSender,
                                           enum StructureChangeType changeType,
                                           SAFEARRAY *runtimeId);

  StructureChangeEventHandler(AutomationContext *ctx);

private:
  LONG mRefCount;
  AutomationContext *mAutomationCtx;
};
