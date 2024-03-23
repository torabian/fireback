import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
 
export const {{ .Template }}Form = ({
  form,
  isEditing,
}: EntityFormProps<{{ .Template }}Entity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      {{ range .e.CompleteFields }}
        
      {{ if or (eq .Type "one") (eq .Type "many2many")  }}
        <FormEntitySelect3
          {{ if eq .Type "many2many"}}
          multiple
          {{ end }}
          formEffect={ { form, field: {{ $.Template }}Entity.Fields.{{ .Name }}$ } }
          useQuery={useGet{{ .TargetWithoutEntityPlural}}}
          label={t.{{ $.templates }}.{{ .Name }} }
          hint={t.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "string") (eq .Type "text")  }}
        <FormText
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={t.{{ $.templates }}.{{ .Name }} }
          hint={t.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "int64") (eq .Type "float64")  }}
        <FormText
          type="number"
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={t.{{ $.templates }}.{{ .Name }} }
          hint={t.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "date") }}
        <FormDate
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={t.{{ $.templates }}.{{ .Name }} }
          hint={t.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "daterange") }}
        <FormDate
          value={values.{{ .Name }}Start }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}Start, value, false)}
          errorMessage={errors.{{ .Name }}Start }
          label={t.{{ $.templates }}.{{ .Name }}Start }
          hint={t.{{ $.templates }}.{{ .Name }}StartHint}
        />
        <FormDate
          value={values.{{ .Name }}End }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}End, value, false)}
          errorMessage={errors.{{ .Name }}End }
          label={t.{{ $.templates }}.{{ .Name }}End }
          hint={t.{{ $.templates }}.{{ .Name }}EndHint}
        />
      {{ else }}
        {/*
          <FormText
            type="?"
            value={values.{{ .Name }} }
            onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
            errorMessage={errors.{{ .Name }} }
            label={t.{{ $.templates }}.{{ .Name }} }
            hint={t.{{ $.templates }}.{{ .Name }}Hint}
          />
         */}
      {{ end }}

      {{ end }}
    </>
  );
};
