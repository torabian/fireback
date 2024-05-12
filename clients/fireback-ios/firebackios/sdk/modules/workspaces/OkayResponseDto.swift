import Foundation
struct OkayResponseDto : Codable {
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
class OkayResponseDtoViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> OkayResponseDto {
      var dto = OkayResponseDto()
      return dto
  }
}
