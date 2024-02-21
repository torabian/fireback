import Foundation

{{ range .children }}
class {{ .FullName }} : Codable, Identifiable {
  {{ template "definitionrow" .CompleteFields }}
}
{{ end }}

class {{ .e.EntityName }} : Codable, Identifiable {
    {{ template "definitionrow" .e.CompleteFields }}
}

{{ template "swiftViewModel" (arr .e .e.CompleteFields)}}
