import Foundation
struct ResetEmailDto : Codable {
    var password: String? = nil
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
class ResetEmailDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  func getDto() -> ResetEmailDto {
      var dto = ResetEmailDto()
    dto.password = self.password
      return dto
  }
}
