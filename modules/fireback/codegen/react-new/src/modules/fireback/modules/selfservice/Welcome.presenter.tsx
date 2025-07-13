import { FormikProps } from "formik";
import { useEffect, useRef, useState } from "react";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useT } from "../../hooks/useT";
import { useGetPassportsAvailableMethods } from "../../sdk/modules/abac/useGetPassportsAvailableMethods";
import {
  AuthAvailableMethods,
  AuthMethod,
  useStoreRedirectParam,
} from "./auth.common";
import {
  CheckPassportMethodsActionResDto,
  ClassicSigninActionReqDto,
} from "../../sdk/modules/abac/AbacActionsDto";

export const usePresenter = () => {
  const t = useT();
  const { locale } = useLocale();
  const { push } = useRouter();
  const formik = useRef<FormikProps<
    Partial<ClassicSigninActionReqDto>
  > | null>();

  const { query: passportMethodsQuery } = useGetPassportsAvailableMethods({
    unauthorized: true,
  });

  useStoreRedirectParam("redirect_temporary");

  const [availableOptions, setAvailableOptions] =
    useState<AuthAvailableMethods>(undefined);

  const totalAvailableMethods = availableOptions
    ? Object.values(availableOptions).filter(Boolean).length
    : undefined;

  const methodData: CheckPassportMethodsActionResDto =
    passportMethodsQuery.data?.data;

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

    const newData = {
      email: methodData.email || false,
      google: methodData.google || false,
      facebook: methodData.facebook || false,
      phone: methodData.phone || false,
      googleOAuthClientKey: methodData.googleOAuthClientKey,
      facebookAppId: (methodData as any).facebookAppId,
    };

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
    passportMethodsQuery,
    isLoadingMethods: passportMethodsQuery.isLoading,
    totalAvailableMethods,
  };
};
