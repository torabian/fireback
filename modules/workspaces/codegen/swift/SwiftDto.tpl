import Foundation

{{ range .children }}

{{ template "extractInlineEnums" (arr .FullName "" .CompleteFields) }}
struct {{ .FullName }} : codable {
    {{ $px := printf "%s%s" .FullName "ActionReqDto" }}
    {{ template "definitionrow" (arr .CompleteFields $px) }}
{{ end }}

{{ template "extractInlineEnums" (arr .e.DtoName "" .e.CompleteFields) }}
struct {{ .e.DtoName }} : Codable {
    {{ $px := printf "%s%s" .e.DtoName "ActionReqDto" }}
    {{ template "definitionrow" (arr .e.CompleteFields $px) }}


    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}

{{ template "swiftViewModel" (arr .e .e.CompleteFields)}}