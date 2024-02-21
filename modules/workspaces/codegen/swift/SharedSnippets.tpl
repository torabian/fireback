
{{ define "definitionrow" }}
  {{ range . }}

  {{ if eq .Type "array" }}
    var {{ .Name }}: [{{ .ComputedType }}]? = nil
  {{ else if eq .Type "many2many" }}
    var {{ .Name }}: [{{ .ComputedType }}]? = nil
  {{ else }}
    var {{ .Name }}: {{ .ComputedType }} = nil
  {{ end }}

  {{ if eq .Type "one" }}
    // var {{ .PrivateName }}Id: String? = nil
  {{ end }}
  {{ if eq .Type "many2many" }}
    var {{ .PrivateName }}ListId: [String]? = nil
  {{ end }}

  {{ end }}
{{ end }}

{{ define "viewModelFieldFnItem" }}
 {{ range . }}

    {{ if or (eq .Type "string") (eq .Type "int64") (eq .Type "bool") }}
    dto.{{ .Name }} = self.{{ .Name }}
    {{ end }}

{{ end }}

{{ end }}

{{ define "viewModelFieldItem" }}
  {{ range . }}
    {{ if or (eq .Type "string") (eq .Type "int64") (eq .Type "bool") }}

  @Published var {{ .Name }}: {{ .ComputedType }} = nil
  @Published var {{ .Name }}ErrorMessage: {{ .ComputedType }} = nil

      {{ end }}
  {{ end }}
{{ end }}

{{ define "viewModelRow" }}
  {{ $e := index . 0 }}
  {{ $fields := index . 1 }}

  // improve the fields here
  {{ template "viewModelFieldItem" $fields }}
  
  func getDto() -> {{ $e.ObjectName }} {
      var dto = {{ $e.ObjectName }}()

      {{ template "viewModelFieldFnItem" $fields }}
     
      return dto
  }

{{ end }}


{{ define "swiftViewModel" }}
{{ $e := index . 0 }}
{{ $fields := index . 1 }}

class {{ $e.ObjectName }}ViewModel: ObservableObject {
  {{ template "viewModelRow" (arr $e $fields) }}
    
}
{{ end }}