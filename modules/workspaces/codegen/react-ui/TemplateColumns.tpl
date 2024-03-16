import { enTranslations } from "@/translations/en";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export const columns = (t: typeof enTranslations) => [
  {
    name: {{ .Template }}Entity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  
  {{ range .e.CompleteFields }}

  {
    name: {{ $.Template }}Entity.Fields.{{ .Name }},
    title: t.{{ $.templates}}.{{ .Name }},
    width: 100,
  },    

  {{ end }}
];
