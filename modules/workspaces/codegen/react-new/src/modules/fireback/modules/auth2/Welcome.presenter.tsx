import { FormikProps } from "formik";
import { useRef } from "react";
import { useLocale } from "../../hooks/useLocale";
import { useRouter } from "../../hooks/useRouter";
import { useT } from "../../hooks/useT";
import { ClassicSigninActionReqDto } from "../../sdk/modules/workspaces/WorkspacesActionsDto";
import { AuthMethod } from "./auth.common";

export const usePresenter = () => {
  const t = useT();
  const { locale } = useLocale();
  const { push } = useRouter();
  const formik = useRef<FormikProps<
    Partial<ClassicSigninActionReqDto>
  > | null>();

  const onSelect = (value: AuthMethod) => {
    switch (value) {
      case AuthMethod.Email:
        push(`/${locale}/auth/email`);
        break;
      case AuthMethod.Phone:
        push(`/${locale}/auth/phone`);
        break;
    }
  };

  return { t, formik, onSelect };
};
