import Foundation
struct OtpAuthenticateDto : Codable {
    var value: String? = nil
    var otp: String? = nil
    var type: String? = nil
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
class OtpAuthenticateDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var otp: String? = nil
  @Published var otpErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  func getDto() -> OtpAuthenticateDto {
      var dto = OtpAuthenticateDto()
    dto.value = self.value
    dto.otp = self.otp
    dto.type = self.type
    dto.password = self.password
      return dto
  }
}
