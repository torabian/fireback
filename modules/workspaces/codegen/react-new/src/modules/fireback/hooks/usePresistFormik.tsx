import { useEffect, useRef } from "react";
import { useRouter } from "./useRouter";
import { FormikProps } from "formik";
import qs from "qs";

export const usePersistFormik = (form: FormikProps<unknown>) => {
  const router = useRouter();

  const unlock = useRef(false);

  useEffect(() => {
    if (unlock.current) {
      const query = qs.stringify(form.values, { addQueryPrefix: true });
      router.replace(query, query, {});
      console.log("triggered");
    }
  }, [form.values]);

  useEffect(() => {
    if (window.location.toString().indexOf("?") !== -1) {
      const params = qs.parse(
        window.location
          .toString()
          .substring(window.location.toString().indexOf("?")),
        {
          ignoreQueryPrefix: true,
        }
      );
      console.log("Params", params);
      form.setValues(params);
    }
    unlock.current = true;
  }, []);
};
