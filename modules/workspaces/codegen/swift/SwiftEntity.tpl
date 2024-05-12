import Foundation

{{ range .children }}
{{ template "extractInlineEnums" (arr .FullName "" .CompleteFields) }}

class {{ .FullName }} : Codable, Identifiable {
  {{ $px := printf "%s%s" .FullName "" }}
  {{ template "definitionrow" (arr .CompleteFields $px) }}
}
{{ end }}

{{ template "extractInlineEnums" (arr .e.EntityName "" .e.CompleteFields) }}
class {{ .e.EntityName }} : Codable, Identifiable {
  {{ $px := printf "%s%s" .e.EntityName "" }}
  {{ template "definitionrow" (arr .e.CompleteFields $px) }}
}

{{ template "swiftViewModel" (arr .e .e.CompleteFields)}}
