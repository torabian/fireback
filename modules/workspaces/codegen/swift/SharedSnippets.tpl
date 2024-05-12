{{ define "extractInlineEnums" }}
  {{ $name := index . 0 }}
  {{ $affix := index . 1 }}
  {{ $fields := index . 2 }}

  {{ range $fields }}
    {{ if eq .Type "enum" }}
enum {{$name}}{{$affix}}{{ .PublicName }} : Codable {
{{ range .OfType}}
  case {{ .Key }}
{{ end }}
}
    {{ end }}

    {{ if or (eq .Type "object") (eq .Type "array")}}
      {{ template "extractInlineEnums" (arr $name $affix .ComputedFields) }}
    {{ end }}
  {{ end }}
{{ end }}
{{ define "definitionrow" }}
  {{ $fields := index . 0 }}
  {{ $groupPrefix := index . 1 }}
  {{ range $fields }}

  {{ if eq .Type "array" }}
    var {{ .Name }}: [{{ .ComputedType }}]? = nil
  {{ else if eq .Type "many2many" }}
    var {{ .Name }}: [{{ .ComputedType }}]? = nil
  {{ else if eq .Type "enum" }}
    var {{ .Name }}: {{ $groupPrefix }}{{ .ComputedType }}? = nil
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

{{ define "rpcActionCommon" }}
  {{/* Common url building for the rpc */}}

  var prefix = ""
  if let api_url = ProcessInfo.processInfo.environment["api_url"] {
    prefix = api_url
  }

  let url = URL(string: prefix + "{{ .Url }}")!

  {{ range .UrlParams}}
  url = url.replace("{{ .}}", with: "{{ .}}")
  {{ end }}

  var request = URLRequest(url: url)

{{ end }}