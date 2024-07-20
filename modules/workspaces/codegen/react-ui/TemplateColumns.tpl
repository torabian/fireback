import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

export const columns = (t: typeof strings) => [
  {{ range .e.CompleteFields }}

  {{ if or (eq .Type "object") (eq .Type "array") (eq .Type "many2many") (eq .Type "one") }}
  {
    name: {{ $.Template }}Entity.Fields.{{ .Name }}$,
    title: t.{{ $.templates}}.{{ .Name }},
    getCellValue: (entity: {{ $.Template }}Entity) => entity.uniqueId,
    width: 100,
  },
  {{ else }}
  {
    name: {{ $.Template }}Entity.Fields.{{ .Name }},
    title: t.{{ $.templates}}.{{ .Name }},
    width: 100,
  },
  {{ end }}

  {{ end }}
];
