import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { TemplateEntity } from "src/sdk/xsdk";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { TemplateEntityFields } from "src/sdk/xsdk/modules/xmodule/xtypefields";
 
export const TemplateForm = ({
  form,
  isEditing,
}: EntityFormProps<TemplateEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <% for (let field of fields) { let name = field.name ; %>
        <% if (field.fbType === "one" || field.fbType === "array") {  %>
          <FormEntitySelect3
            <%-field.fbType === 'array' ? 'multiple' : '' %>
            formEffect={{ form, field: TemplateEntityFields.<%- name %>$ }}
            useQuery={<%- castGetQueryFromGolangType(field.type) %>}
            label={t.templates.<%- name %>}
            hint={t.templates.<%- name %>Hint}
          />
        <% } else if (field.type.includes("string") || field.type.includes("bool")) { %>
          <FormText
            value={values.<%- name %>}
            onChange={(value) => setFieldValue(TemplateEntityFields.<%- name %>, value, false)}
            errorMessage={errors.<%- name %>}
            label={t.templates.<%- name %>}
            hint={t.templates.<%- name %>Hint}
          />

        <% } else if (field.type.includes("int64") || field.type.includes("float64")) { %>
          <FormText
            type="number"
            value={values.<%- name %>}
            onChange={(value) => setFieldValue(TemplateEntityFields.<%- name %>, value, false)}
            errorMessage={errors.<%- name %>}
            label={t.templates.<%- name %>}
            hint={t.templates.<%- name %>Hint}
          />

        <% } else { %>

          {/*
          Unkown field: <%- field.type %>
          Name: <%- field.jsonField %>
          <FormText
            value={values.<%- name %>}
            onChange={(value) => setFieldValue(TemplateEntityFields.<%- name %>, value, false)}
            errorMessage={errors.<%- name %>}
            label={t.templates.<%- name %>}
            hint={t.templates.<%- name %>Hint}
          />
          
          */}
        <% } %>
      <% } %>
       
    </>
  );
};
