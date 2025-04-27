{{ define "kotlinimport" }}
{{ range $key, $value := .kotlinimports }}
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


{{ define "viewmodeltype" }}{{ if .Module }}com.fireback.modules.{{ end }}{{ if eq .ComputedType "int" }} {{ "Integer" }} {{ else if eq .ComputedType "float" }} Float {{ else }} {{ .ComputedType }} {{ end }}{{ end }}


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

{{ define "castToModel" }}
   
{{ end }}

{{ define "kotlinClassContent" }}
  {{ template "definitionrow" .CompleteFields }}
  public static class VM extends ViewModel {
    // Fields to work with as form field (dto)
    {{ template "viewmodelrow" .CompleteFields }}

    // Handling error message for each field
    {{ template "viewModelMessageRow" .CompleteFields }}

  }
{{ end }}

{{ define "applyExceptionOnViewModel" }}
public void applyException(Throwable e) {
    if (!(e instanceof ResponseErrorException)) {
        return;
    }
    ResponseErrorException responseError = (ResponseErrorException) e;

    // @todo on fireback: This needs to be recursive.
    responseError.error.errors.forEach(item -> {
        {{ range . }}
        if (item.location != null && item.location.equals("{{ .Name }}")) {
            this.set{{.PublicName}}Msg(item.messageTranslated);
        }
        {{ end }}
    });
}
{{ end }}


{{ define "viewModelMessageRow" }}

    {{ range . }}
    // upper: {{ .PublicName }} {{ .Name }}
    private MutableLiveData<String> {{ .Name }}Msg = new MutableLiveData<>();
    public MutableLiveData<String> get{{ .PublicName }}Msg() {
        return {{ .Name }}Msg;
    }

    public void set{{ .PublicName }}Msg(String v) {
        {{ .Name }}Msg.setValue(v);
    }
    
    {{ end }}

{{ end }}


