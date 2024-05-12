import Combine
import Foundation 

{{ range .m.Actions }}

    {{ if .In.Fields }}

{{ template "extractInlineEnums" (arr .Upper "ActionReqDto" .In.Fields) }}
struct {{ .Upper }}ActionReqDto : Codable {
    {{ template "definitionrow" .In.Fields }}

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



class {{ .Upper }}ActionReqDtoVm: ObservableObject {

    {{ template "viewModelFieldItem" .In.Fields }}
 
    func getDto() -> {{ .Upper }}ActionReqDto {
        var dto = {{ .Upper }}ActionReqDto()
        {{ template "viewModelFieldFnItem" .In.Fields }}
        return dto
    }
}

    {{ end }}

    {{ if .Out.Fields }}

struct {{ .Upper }}ActionResDto : Codable {
    {{ template "definitionrow" .Out.Fields }}

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

    {{ end }}

{{ end }}