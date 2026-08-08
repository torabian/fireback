import { httpErrorHanlder, mutationErrorsToFormik } from "../../hooks/api";
import { Toast } from "../../hooks/toast";
import { useCommonEntityManager } from "../../hooks/useCommonEntityManager";
import { useS } from "../../hooks/useS";
import { strings } from "../strings/translations";
import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useContext, useEffect, useRef, useState } from "react";

import { useBackButton, useCommonCrudActions } from "../action-menu/ActionMenu";
import { QueryErrorView } from "../error-view/QueryError";
import { usePageTitle } from "../page-title/PageTitle";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";
import { get, set } from "lodash";
import type { GResponse } from "../../../sdk/sdk/envelopes";
import { ErrorsView } from "../error-view/ErrorView";
import { KeyboardAction } from "../../hooks/useExportTools";

export interface CommonEntityManagerProps<T> {
  data?: T | null;
  Form?: any;
  getSingleHook?: any;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  postHook?: any;
  forceEdit?: boolean;
  disableOnGetFailed?: boolean;
  patchHook?: any;
  onlyOnRoot?: boolean;
  onEditTitle?: string;
  beforeSetValues?: (data: Partial<T>) => Partial<T>;
  customClass?: string;
  onCreateTitle?: string;
  onCancel?: () => void;
  beforeSubmit?: (data: T) => T;
  onSuccessPatchOrPost?: (response: GResponse<any>) => void;
  onFinishUriResolver?: (response: GResponse<any>, locale: string) => string;
}

export interface DtoEntity<T, V = null> {
  data?: Partial<T> | null;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  enabledFields?: Partial<V>;
  onSuccess?: (response: GResponse<T>) => void;
  showSubmit?: boolean;
  Form?: any;
}

export const CommonEntityManager = ({
  data,
  Form,
  getSingleHook,
  postHook,
  onCancel,
  onFinishUriResolver,
  disableOnGetFailed,
  patchHook,
  onCreateTitle,
  onEditTitle,
  setInnerRef,
  beforeSetValues,
  forceEdit,
  onlyOnRoot,
  customClass,
  beforeSubmit,
  onSuccessPatchOrPost,
}: CommonEntityManagerProps<any>) => {
  const [initialData, setInitialData] = useState();
  const { router, isEditing, locale, formik } = useCommonEntityManager<
    Partial<any>
  >({
    data,
  });
  const s = useS(strings);

  const touchedData = useRef({});

  useBackButton(onCancel, KeyboardAction.CommonBack);
  const { selectedUrw } = useContext(RemoteQueryContext);
  usePageTitle((isEditing || forceEdit ? onEditTitle : onCreateTitle) || "");

  const getQuery = getSingleHook;

  useEffect(() => {
    const item = getQuery?.data?.data?.item?.toJSON();
    if (item) {
      formik.current?.setValues(
        beforeSetValues ? beforeSetValues({ ...item }) : { ...item },
      );

      setInitialData(item);
    }
  }, [getQuery?.data]);

  useEffect(() => {
    formik.current?.setSubmitting(postHook?.isLoading || patchHook?.isLoading);
  }, [postHook?.isLoading, patchHook?.isLoading]);

  const onSubmit = (p: Partial<any>, d: FormikHelpers<Partial<any>>) => {
    let values: any = touchedData.current;
    values.uniqueId = p.uniqueId;
    if (beforeSubmit) {
      values = beforeSubmit(values);
    }

    const op =
      isEditing || forceEdit
        ? patchHook?.mutateAsync(JSON.stringify(values), d)
        : postHook?.mutateAsync(JSON.stringify(values), d);

    op.then((response: GResponse<unknown>) => {
      if (response.data?.item) {
        if (onSuccessPatchOrPost) {
          onSuccessPatchOrPost(response);
        } else if (onFinishUriResolver) {
          router.goBackOrDefault(onFinishUriResolver(response, locale));
        } else {
          Toast(s.components.done, { type: "success" });
        }
      }
      if (response.error) {
        formik.current.setErrors(
          mutationErrorsToFormik(response.error.toJSON()),
        );
      }
    }).catch((err) => httpErrorHanlder(err, s));
  };

  const formWorking =
    getSingleHook?.isLoading ||
    false ||
    postHook?.query?.isLoading ||
    false ||
    patchHook?.query?.isLoading ||
    false;

  useCommonCrudActions({
    // onCancel: onCancel,
    onSave() {
      formik.current?.submitForm();
    },
  });

  if (onlyOnRoot && selectedUrw.workspaceId !== "root") {
    return <div>{s.onlyOnRoot}</div>;
  }

  return (
    <Formik
      innerRef={(r) => {
        if (r) {
          formik.current = r;
          setInnerRef && setInnerRef(r);
        }
      }}
      initialValues={{}}
      onSubmit={onSubmit}
    >
      {(form: FormikProps<Partial<any>>) => (
        <form
          onSubmit={(e) => {
            e.preventDefault();
            form.submitForm();
          }}
          className={
            customClass == undefined
              ? "headless-form-entity-manager"
              : customClass
          }
        >
          {/* <pre>{JSON.stringify(form.values, null, 2)}</pre> */}
          <ErrorsView errors={form.errors} />
          <fieldset disabled={formWorking}>
            <div style={{ marginBottom: "15px" }}>
              <QueryErrorView
                query={
                  postHook?.mutation?.isError
                    ? postHook.mutation
                    : patchHook?.mutation?.isError
                      ? patchHook.mutation
                      : getSingleHook?.query?.isError
                        ? getSingleHook.query
                        : null
                }
              />
            </div>
            {disableOnGetFailed === true &&
            getSingleHook?.query?.isError ? null : (
              <Form
                isEditing={isEditing}
                initialData={initialData}
                form={{
                  ...form,

                  setValues: (
                    values: React.SetStateAction<any>,
                    shouldValidate?: boolean,
                  ) => {
                    for (const key in values) {
                      set(touchedData.current, key, values[key]);
                    }

                    return form.setValues(values);
                  },

                  setFieldValue: (
                    field: string,
                    value: any,
                    shouldValidate?: boolean,
                  ) => {
                    // In case of having a nested object, we touch the entire nested for safety.
                    // This is completely correct for json fields for example, but might not be
                    // most efficient for object or embed types in fireback.
                    if (field.includes(".")) {
                      const v = field.split(".")[0];
                      set(touchedData.current, v, get(form.values, v));
                    }

                    set(touchedData.current, field, value);

                    return form.setFieldValue(field, value, shouldValidate);
                  },
                }}
              />
            )}
            <button type="submit" className="d-none" />
          </fieldset>
        </form>
      )}
    </Formik>
  );
};
