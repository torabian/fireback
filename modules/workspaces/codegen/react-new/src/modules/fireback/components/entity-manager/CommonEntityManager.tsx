import { httpErrorHanlder } from "../../hooks/api";
import { Toast } from "../../hooks/toast";
import { useCommonEntityManager } from "../../hooks/useCommonEntityManager";
import { Formik, FormikHelpers, FormikProps } from "formik";
import { useContext, useEffect } from "react";
import { KeyboardAction } from "../../definitions/definitions";
import { useBackButton, useCommonCrudActions } from "../action-menu/ActionMenu";
import { QueryErrorView } from "../error-view/QueryError";
import { usePageTitle } from "../page-title/PageTitle";
import { IResponse } from "../../definitions/JSONStyle";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

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
  customClass?: string;
  onCreateTitle?: string;
  onCancel?: () => void;
  beforeSubmit?: (data: T) => T;
  onSuccessPatchOrPost?: (response: IResponse<any>) => void;
  onFinishUriResolver?: (response: IResponse<any>, locale: string) => string;
}

export interface DtoEntity<T, V = null> {
  data?: Partial<T> | null;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
  enabledFields?: Partial<V>;
  onSuccess?: (response: IResponse<T>) => void;
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
  forceEdit,
  onlyOnRoot,
  customClass,
  beforeSubmit,
  onSuccessPatchOrPost,
}: CommonEntityManagerProps<any>) => {
  const { router, isEditing, locale, formik, t } = useCommonEntityManager<
    Partial<any>
  >({
    data,
  });

  useBackButton(onCancel, KeyboardAction.CommonBack);
  const { selectedUrw } = useContext(RemoteQueryContext);
  usePageTitle((isEditing || forceEdit ? onEditTitle : onCreateTitle) || "");

  const { query: getQuery } = getSingleHook;

  useEffect(() => {
    if (getQuery.data?.data) {
      formik.current?.setValues({
        ...getQuery.data.data,
      });
    }
  }, [getQuery.data]);

  useEffect(() => {
    formik.current?.setSubmitting(
      postHook?.mutation.isLoading || patchHook?.mutation.isLoading
    );
  }, [postHook?.isLoading, patchHook?.isLoading]);

  const onSubmit = (values: Partial<any>, d: FormikHelpers<Partial<any>>) => {
    if (beforeSubmit) {
      values = beforeSubmit(values);
    }

    const op =
      isEditing || forceEdit
        ? patchHook?.submit(values, d)
        : postHook?.submit(values, d);

    op.then((response: any) => {
      if (response.data?.uniqueId) {
        if (onSuccessPatchOrPost) {
          onSuccessPatchOrPost(response);
        } else if (onFinishUriResolver) {
          router.goBackOrDefault(onFinishUriResolver(response, locale));
        } else {
          Toast("Done", { type: "success" });
        }
      }
    }).catch((err) => httpErrorHanlder(err, t));
  };

  const formWorking =
    getSingleHook?.query?.isLoading ||
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
    return <div>{t.onlyOnRoot}</div>;
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
          {/* <ErrorsView errors={form.errors} /> */}
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
              <Form isEditing={isEditing} form={form} />
            )}
            <button type="submit" className="d-none" />
          </fieldset>
        </form>
      )}
    </Formik>
  );
};
