import { usePageTitle } from "@/components/page-title/PageTitle";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { Formik, FormikHelpers, FormikProps } from "formik";
import { useT } from "@/hooks/useT";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormRichText } from "@/components/forms/form-richtext/FormRichText";
import { FormButton } from "@/components/forms/form-button/FormButton";

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
  const t = useT();
  const { router, uniqueId, queryClient, isEditing, locale, formik } =
    useCommonEntityManager<Partial<MailTemplateEditor>>({
      data,
    });

  usePageTitle(
    isEditing ? t.wokspaces.createNewWorkspace : t.wokspaces.editWorkspae
  );
  const onSubmit = (
    values: Partial<MailTemplateEditor>,
    d: FormikHelpers<Partial<MailTemplateEditor>>
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
            label={t.wokspaces.title}
          />
          <FormRichText
            onChange={(val) => form.setFieldValue("body", val)}
            value={form.values.body}
            label={t.wokspaces.body}
          />
          <FormButton
            label={t.wokspaces.resetToDefault}
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
