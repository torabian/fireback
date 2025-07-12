import { useFormik } from "formik";
import { useRouter } from "../../hooks/useRouter";
import { useS } from "../../hooks/useS";
import { usePostPassportChangePassword } from "../../sdk/modules/abac/usePostPassportChangePassword";
import { strings } from "./strings/translations";

export const usePresenter = () => {
  const s = useS(strings);
  const { goBack, query } = useRouter();
  const { submit: changePassword, mutation } = usePostPassportChangePassword();

  const form = useFormik({
    initialValues: {},
    onSubmit: () => {
      alert("done");
    },
  });

  return {
    mutation,
    form,

    goBack,
    s,
  };
};
