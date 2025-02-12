import { EntityFormProps } from "{{ .FirebackUiDir }}/definitions/definitions";
import { RemoteQueryContext } from "{{ .SdkDir }}/core/react-tools";
import { useContext } from "react";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { FormText } from "{{ .FirebackUiDir }}/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "{{ .FirebackUiDir }}/components/forms/form-select/FormEntitySelect3";
import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

export const {{ .Template }}Form = ({
  form,
  isEditing,
}: EntityFormProps<{{ .Template }}Entity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
      {{ range .e.CompleteFields }}
        
      {{ if or (eq .Type "one") (eq .Type "many2many")  }}
        <FormEntitySelect
          {{ if eq .Type "many2many"}}
          multiple
          {{ end }}
          formEffect={ { form, field: {{ $.Template }}Entity.Fields.{{ .Name }}$ } }
          useQuery={useGet{{ .TargetWithoutEntityPlural}}}
          label={s.{{ $.templates }}.{{ .Name }} }
          hint={s.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "string") (eq .Type "text")  }}
        <FormText
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={s.{{ $.templates }}.{{ .Name }} }
          hint={s.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "int64") (eq .Type "float64")  }}
        <FormText
          type="number"
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={s.{{ $.templates }}.{{ .Name }} }
          hint={s.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "date") }}
        <FormDate
          value={values.{{ .Name }} }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
          errorMessage={errors.{{ .Name }} }
          label={s.{{ $.templates }}.{{ .Name }} }
          hint={s.{{ $.templates }}.{{ .Name }}Hint}
        />
      {{ else if or (eq .Type "daterange") }}
        <FormDate
          value={values.{{ .Name }}Start }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}Start, value, false)}
          errorMessage={errors.{{ .Name }}Start }
          label={s.{{ $.templates }}.{{ .Name }}Start }
          hint={s.{{ $.templates }}.{{ .Name }}StartHint}
        />
        <FormDate
          value={values.{{ .Name }}End }
          onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}End, value, false)}
          errorMessage={errors.{{ .Name }}End }
          label={s.{{ $.templates }}.{{ .Name }}End }
          hint={s.{{ $.templates }}.{{ .Name }}EndHint}
        />
      {{ else }}
        {/*
          <FormText
            type="?"
            value={values.{{ .Name }} }
            onChange={(value) => setFieldValue({{ $.Template }}Entity.Fields.{{ .Name }}, value, false)}
            errorMessage={errors.{{ .Name }} }
            label={s.{{ $.templates }}.{{ .Name }} }
            hint={s.{{ $.templates }}.{{ .Name }}Hint}
          />
         */}
      {{ end }}

      {{ end }}
    </>
  );
};
