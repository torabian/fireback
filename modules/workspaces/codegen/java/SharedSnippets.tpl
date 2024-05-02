{{ define "javaimport" }}
{{ range $key, $value := .javaimports }}
{{ if and ($value.Items) ($key) }}
import {{ $key}}.*;
{{ end }}
{{ end }}
{{ end }}

{{ define "definitionrow" }}

    {{ range . }}
    public {{ if .Module }}com.fireback.modules.{{ .Module }}.{{ end }}{{ .ComputedType }} {{ .Name }};
    {{ end }}

{{ end }}


{{ define "viewmodeltype" }}{{ if eq .ComputedType "int" }} {{ "Integer" }} {{ else if eq .ComputedType "float" }} Float {{ else }} {{ .ComputedType }} {{ end }}{{ end }}


{{ define "viewmodelrow" }}

    {{ range . }}
    // upper: {{ .PublicName }} {{ .Name }}
    private MutableLiveData<{{ template "viewmodeltype" . }}> {{ .Name }} = new MutableLiveData<>();
    public MutableLiveData<{{ template "viewmodeltype" . }}> get{{ .PublicName }}() {
        return {{ .Name }};
    }

    public void set{{ .PublicName }}({{ template "viewmodeltype" . }} v) {
        {{ .Name }}.setValue(v);
    }
    
    {{ end }}

{{ end }}


