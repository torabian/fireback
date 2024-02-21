{{ template "tsimport" . }}

// In this section we have sub entities related to this object
{{ range .children }}
export class {{ .FullName }} extends BaseDto {
  {{ template "definitionrow" .CompleteFields }}
}
{{ end }}

// Class body

export type {{ .e.DtoName }}Keys =
  keyof typeof {{ .e.DtoName }}.Fields;

export class {{ .e.DtoName }} extends BaseDto {

  {{ template "definitionrow" .e.CompleteFields }}
  
  {{ if eq .ctx.Ts.IncludeStaticField true }}
    {{ template "staticfield" . }}
  {{ end }}

  {{ if eq .ctx.Ts.IncludeFirebackDef true }}
    {{ template "fbdefinition" . }}
  {{ end }}
}
