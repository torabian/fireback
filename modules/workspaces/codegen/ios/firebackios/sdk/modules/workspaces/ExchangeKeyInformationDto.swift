import Foundation
struct ExchangeKeyInformationDto : Codable {
    var key: String? = nil
    var visibility: String? = nil
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
class ExchangeKeyInformationDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var key: String? = nil
  @Published var keyErrorMessage: String? = nil
  @Published var visibility: String? = nil
  @Published var visibilityErrorMessage: String? = nil
  func getDto() -> ExchangeKeyInformationDto {
      var dto = ExchangeKeyInformationDto()
    dto.key = self.key
    dto.visibility = self.visibility
      return dto
  }
}
