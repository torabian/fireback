import Foundation
struct EmailOtpResponseDto : Codable {
    var request: ForgetPasswordEntity? = nil
    // var requestId: String? = nil
    var userSession: UserSessionDto? = nil
    // var userSessionId: String? = nil
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
class EmailOtpResponseDtoViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> EmailOtpResponseDto {
      var dto = EmailOtpResponseDto()
      return dto
  }
}
