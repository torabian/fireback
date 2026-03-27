import { FormikProps } from "formik";
import { useEffect, useRef, useState } from "react";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useT } from "../../hooks/useT";
import { ClassicSigninActionReqDto } from "../../sdk/modules/abac/AbacActionsDto";
import { useCheckPassportMethodsActionQuery } from "../../sdk/modules/abac/CheckPassportMethods";
import {
  AuthAvailableMethods,
  AuthMethod,
  useTemporaryParamOptions,
} from "./auth.common";

export const usePresenter = () => {
  const t = useT();
  const { locale } = useLocale();
  const { push } = useRouter();
  const formik = useRef<FormikProps<
    Partial<ClassicSigninActionReqDto>
  > | null>();

  const query = useCheckPassportMethodsActionQuery({});

  useTemporaryParamOptions(["redirect_temporary", "workspace_type_id"]);

  const [availableOptions, setAvailableOptions] =
    useState<AuthAvailableMethods>(undefined);

  const totalAvailableMethods = availableOptions
    ? Object.values(availableOptions).filter(Boolean).length
    : undefined;

  const methodData = query.data?.data?.item;

  const onSelect = (value: AuthMethod, canGoBack = true) => {
    switch (value) {
      case AuthMethod.Email:
        push(`/${locale}/selfservice/email`, undefined, {
          canGoBack,
        });
        break;
      case AuthMethod.Phone:
        push(`/${locale}/selfservice/phone`, undefined, {
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
    t,
    formik,
    onSelect,
    availableOptions,
    passportMethodsQuery: query,
    isLoadingMethods: query.isLoading,
    totalAvailableMethods,
  };
};
