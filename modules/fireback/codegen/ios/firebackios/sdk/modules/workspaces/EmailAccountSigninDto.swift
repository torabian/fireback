import Foundation
struct EmailAccountSigninDto : Codable {
    var email: String? = nil
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
class EmailAccountSigninDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var email: String? = nil
  @Published var emailErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  func getDto() -> EmailAccountSigninDto {
      var dto = EmailAccountSigninDto()
    dto.email = self.email
    dto.password = self.password
      return dto
  }
}
