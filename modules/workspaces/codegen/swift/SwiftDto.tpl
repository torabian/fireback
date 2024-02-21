import Foundation

{{ range .children }}
struct {{ .FullName }} : codable {
    {{ template "definitionrow" .e.CompleteFields }}
{{ end }}

struct {{ .e.DtoName }} : Codable {
    {{ template "definitionrow" .e.CompleteFields }}


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