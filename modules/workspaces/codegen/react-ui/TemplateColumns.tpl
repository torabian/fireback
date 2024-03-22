import { enTranslations } from "@/translations/en";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export const columns = (t: typeof enTranslations) => [
  {
    name: {{ .Template }}Entity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  
  {{ range .e.CompleteFields }}

  {{ if or (.Type eq "object") (.Type eq "array") (.Type eq "many2many") (.Type eq "one") }}
  {
    name: {{ $.Template }}Entity.Fields.{{ .Name }}$,
    title: t.{{ $.templates}}.{{ .Name }},
    getCellValue: (entity: {{ $.Template }}Entity) => entity.uniqueId
    width: 100,
  },
  {{ else }}
  {
    name: {{ $.Template }}Entity.Fields.{{ .Name }},
    title: t.{{ $.templates}}.{{ .Name }},
    width: 100,
  },
  {{ end }}
];
