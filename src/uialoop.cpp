#include <UIAutomationClient.h>
#include <cpplogger/cpplogger.h>
#include <iostream>
#include <oaidl.h>

#include "context.h"
#include "uiahandler.h"
#include "uialoop.h"
#include "util.h"

extern Logger::Logger *Log;

static HWND hForegroundWindow{};
static IUIAutomationCondition *pTrueCondition{};

static IUIAutomationElementArray *pFound{};
static IUIAutomationElement *pForegroundElement{};
static IUIAutomationElement *pFocusElement{};

static PROPERTYID itemIndexPropertyId{};
static PROPERTYID itemCountPropertyId{};

HRESULT updateForegroundElement(IUIAutomation *pUIAutomation) {
  HWND hw{};
  HRESULT hr{};

  hw = GetForegroundWindow();

  if (hw == hForegroundWindow) {
    return S_OK;
  }

  hForegroundWindow = hw;

  SafeRelease(&pFocusElement);
  SafeRelease(&pForegroundElement);
  SafeRelease(&pFound);

  if (hForegroundWindow != nullptr) {
    hr = pUIAutomation->ElementFromHandle(hForegroundWindow,
                                          &pForegroundElement);
  } else {
    hr = pUIAutomation->GetRootElement(&pForegroundElement);
  }

  return hr;
}

HRESULT findFocusElement(IUIAutomation *pUIAutomation, SAFEARRAY *r1) {
  HRESULT hr{};

  if (r1 == nullptr) {
    return E_FAIL;
  }
  if (pFound == nullptr) {
    hr =
        pForegroundElement->FindAll(TreeScope_Subtree, pTrueCondition, &pFound);
  }
  if (FAILED(hr)) {
    return hr;
  }

  int length{};

  hr = pFound->get_Length(&length);

  if (FAILED(hr)) {
    return hr;
  }

  SAFEARRAY *r2{};
  BOOL areSame{};
  IUIAutomationElement *pElement{};

  for (int i = 0; i < length; i++) {
    hr = pFound->GetElement(i, &pElement);

    if (FAILED(hr)) {
      continue;
    }

    hr = pElement->GetRuntimeId(&r2);

    if (FAILED(hr)) {
      continue;
    }

    hr = pUIAutomation->CompareRuntimeIds(r1, r2, &areSame);

    if (FAILED(hr)) {
      continue;
    }
    if (!areSame) {
      continue;
    }

    pFocusElement = pElement;

    return S_OK;
  }

  return hr;
}
HRESULT registerPropertyId(const std::wstring &propertyGUID,
                           const std::wstring &propertyName,
                           UIAutomationType propertyType,
                           PROPERTYID *propertyId) {
  HRESULT hr{};
  GUID guid{};

  hr = CLSIDFromString(propertyGUID.c_str(), &guid);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call CLSIDFromString", GetCurrentThreadId(),
              __LONGFILE__);
    return hr;
  }

  IUIAutomationRegistrar *pRegistrar{};

  hr = CoCreateInstance(CLSID_CUIAutomationRegistrar, nullptr,
                        CLSCTX_INPROC_SERVER, IID_IUIAutomationRegistrar,
                        reinterpret_cast<void **>(&pRegistrar));

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call CoCreateInstance", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

  UIAutomationPropertyInfo info = {guid, propertyName.c_str(), propertyType};

  hr = pRegistrar->RegisterProperty(&info, propertyId);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call RegisterProperty", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

CLEANUP:
  SafeRelease(&pRegistrar);

  return S_OK;
}

DWORD WINAPI uiaThread(LPVOID context) {
  Log->Info(L"Start UI automation thread", GetCurrentThreadId(), __LONGFILE__);

  HRESULT hr{};
  bool isActive{true};

  hr = CoInitializeEx(nullptr, COINIT_MULTITHREADED);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call CoInitializeEx", GetCurrentThreadId(),
              __LONGFILE__);
    return hr;
  }

  AutomationContext *ctx = static_cast<AutomationContext *>(context);

  if (ctx == nullptr) {
    Log->Fail(L"Failed to obtain context", GetCurrentThreadId(), __LONGFILE__);
    return E_FAIL;
  }

  hr = CoCreateInstance(__uuidof(CUIAutomation8), nullptr, CLSCTX_INPROC_SERVER,
                        __uuidof(IUIAutomation5),
                        reinterpret_cast<void **>(&(ctx->UIAutomation)));

  if (FAILED(hr)) {
    Log->Fail(L"Failed to create an instance of IUIAutomation",
              GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }

  hr = ctx->UIAutomation->CreateTrueCondition(&pTrueCondition);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call IUIAutomation::CreateTrueCondition()",
              GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }

  FocusChangeEventHandler *pFocusChangeEventHandler =
      new FocusChangeEventHandler(ctx);

  hr = ctx->UIAutomation->AddFocusChangedEventHandler(ctx->BaseCacheRequest,
                                                      pFocusChangeEventHandler);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to call IUIAutomation::AddFocusChangedEventHandler",
              GetCurrentThreadId(), __LONGFILE__);
    goto CLEANUP;
  }

  hr = registerPropertyId(L"{92A053DA-2969-4021-BF27-514CFC2E4A69}",
                          L"ItemIndex", UIAutomationType_Int,
                          &itemIndexPropertyId);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to register ItemIndex property", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

  hr = registerPropertyId(L"{ABBF5C45-5CCC-47b7-BB4E-87CB87BBD162}",
                          L"ItemCount", UIAutomationType_Int,
                          &itemCountPropertyId);

  if (FAILED(hr)) {
    Log->Fail(L"Failed to register ItemCount property", GetCurrentThreadId(),
              __LONGFILE__);
    goto CLEANUP;
  }

  Log->Info(L"Wait for quit event", GetCurrentThreadId(), __LONGFILE__);

  while (isActive) {
    HANDLE waitArray[2] = {ctx->QuitEvent, ctx->FocusEvent};
    DWORD waitResult = WaitForMultipleObjects(2, waitArray, FALSE, INFINITE);

    switch (waitResult) {
    case WAIT_OBJECT_0 + 0: // ctx->QuitEvent
      isActive = false;
      break;
    case WAIT_OBJECT_0 + 1: // ctx->FocusEvent
      hr = updateForegroundElement(ctx->UIAutomation);

      if (FAILED(hr)) {
        Log->Warn(L"Failed to update foreground element.", GetCurrentThreadId(),
                  __LONGFILE__);
        break;
      }

      hr = findFocusElement(ctx->UIAutomation, ctx->FocusElementRuntimeId);

      if (FAILED(hr)) {
        Log->Warn(L"Failed to find focus element.", GetCurrentThreadId(),
                  __LONGFILE__);
        break;
      }

      Log->Info(L"Found focus element.", GetCurrentThreadId(), __LONGFILE__);

      break;
    }
  }

CLEANUP:

  Log->Info(L"Cleanup", GetCurrentThreadId(), __LONGFILE__);

  SafeRelease(&pTrueCondition);
  SafeRelease(&(ctx->UIAutomation));

  CoUninitialize();

  Log->Info(L"End UI automation thread", GetCurrentThreadId(), __LONGFILE__);

  return S_OK;
}
