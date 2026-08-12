import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { Formik, type FormikHelpers, type FormikProps } from "formik";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { FormRichText } from "@fireback/ui-core/components/forms/form-richtext/FormRichText";
import { FormButton } from "@fireback/ui-core/components/forms/form-button/FormButton";

export interface MailTemplateEditor {
  title?: string;
  body?: string;
  defaultBody?: string;
  defaultTitle?: string;
}

interface DtoEntity<T> {
  data?: T | null;
  setInnerRef?: (ref: FormikProps<Partial<T>>) => void;
}

export const MailTemplateEntityManager = ({
  data,
  setInnerRef,
}: DtoEntity<MailTemplateEditor>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<MailTemplateEditor>>({
      data,
    });

  usePageTitle(isEditing ? s.createNewWorkspace : s.editWorkspae);
  const onSubmit = (
    values: Partial<MailTemplateEditor>,
    d: FormikHelpers<Partial<MailTemplateEditor>>,
  ) => {};

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
      {(form: FormikProps<Partial<MailTemplateEditor>>) => (
        <form onSubmit={(e) => e.preventDefault()}>
          <FormText
            value={form.values.title}
            onChange={(val) => form.setFieldValue("title", val)}
            label={s.title}
          />
          <FormRichText
            onChange={(val) => form.setFieldValue("body", val)}
            value={form.values.body}
            label={s.body}
          />
          <FormButton
            label={s.resetToDefault}
            onClick={() =>
              form.setValues({
                ...form.values,
                body: form.values.defaultBody,
                title: form.values.defaultTitle,
              })
            }
          />
        </form>
      )}
    </Formik>
  );
};
