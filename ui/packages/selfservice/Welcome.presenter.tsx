import { type FormikProps } from "formik";
import { useEffect, useRef, useState } from "react";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useCheckPassportMethodsActionQuery } from "@fireback/selfservice/sdk/abac/CheckPassportMethodsAction";
import {
  type AuthAvailableMethods,
  AuthMethod,
  useTemporaryParamOptions,
} from "./auth.common";
import type { ClassicSigninActionReq } from "@fireback/selfservice/sdk/abac/ClassicSigninAction";

export const usePresenter = () => {
  const { push } = useRouter();
  const formik = useRef<FormikProps<Partial<ClassicSigninActionReq>> | null>();

  const query = useCheckPassportMethodsActionQuery({});

  // Key must be "redirect" - it's both the ?redirect=... query/hash param name
  // useTemporaryParamOptions looks for *and* the sessionStorage key it's stashed
  // under (see its own implementation - the same string is used for both), which
  // is what auth.common.tsx's onComplete later reads back as `redirect2`. A
  // mismatched name here (this used to read "redirect_temporary", a param no
  // caller ever actually sent) meant a real `?redirect=` was silently never
  // captured - the login flow always fell through to the default route instead.
  useTemporaryParamOptions(["redirect", "workspace_type_id"]);

  const [availableOptions, setAvailableOptions] =
    useState<AuthAvailableMethods>(undefined);

  const totalAvailableMethods = availableOptions
    ? Object.values(availableOptions).filter(Boolean).length
    : undefined;

  const methodData = query.data?.data?.item;

  const onSelect = (value: AuthMethod, canGoBack = true) => {
    switch (value) {
      case AuthMethod.Email:
        push(`/selfservice/email`, undefined, {
          canGoBack,
        });
        break;
      case AuthMethod.Phone:
        push(`/selfservice/phone`, undefined, {
          canGoBack,
        });
        break;
    }
  };

  useEffect(() => {
    if (!methodData) {
      return;
    }

    // Extract the authentication methods here.
    // Make sure, you select only fields which are indicating an option,
    // because adding extra fields here might interfer with auto-selection later.
    const newData = {
      email: methodData.email,
      google: methodData.google,
      facebook: methodData.facebook,
      phone: methodData.phone,
      googleOAuthClientKey: methodData.googleOAuthClientKey,
      facebookAppId: (methodData as any).facebookAppId,
    };

    // If there is only a single method to login available
    const totalAvailableMethods = Object.values(newData).filter(Boolean).length;

    if (totalAvailableMethods === 1) {
      if (newData.email) {
        onSelect(AuthMethod.Email, false);
      }
      if (newData.phone) {
        onSelect(AuthMethod.Phone, false);
      }
      if (newData.google) {
        onSelect(AuthMethod.Google, false);
      }
      if (newData.facebook) {
        onSelect(AuthMethod.Facebook, false);
      }
    }

    setAvailableOptions(newData);
  }, [methodData]);

  return {
    formik,
    onSelect,
    availableOptions,
    passportMethodsQuery: query,
    isLoadingMethods: query.isLoading,
    totalAvailableMethods,
  };
};
