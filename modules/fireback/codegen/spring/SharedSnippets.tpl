{{ define "javaimport" }}
{{ range $key, $value := .javaimports }}
{{ if and ($value.Items) ($key) }}
import {{ $key}}.*;
{{ end }}
{{ end }}
{{ end }}


{{ define "definitionrow" }}

    {{ range . }}
    public {{ if .Module }}com.fireback.modules.{{ end }}{{ .ComputedType }} {{ .Name }};
    {{ end }}

{{ end }}


{{ define "javaClassContent" }}
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  public String uniqueId;

  {{ template "definitionrow" .CompleteFields }}

{{ end }}


{{ define "dtoClassContent" }}

  {{ template "definitionrow" .CompleteFields }}

{{ end }}