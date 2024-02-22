import Foundation
struct CronjobConfigDto : Codable {
    var expression: String? = nil
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
class CronjobConfigDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var expression: String? = nil
  @Published var expressionErrorMessage: String? = nil
  func getDto() -> CronjobConfigDto {
      var dto = CronjobConfigDto()
    dto.expression = self.expression
      return dto
  }
}
