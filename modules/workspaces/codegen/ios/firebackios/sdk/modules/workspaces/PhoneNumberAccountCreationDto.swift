import Foundation
struct PhoneNumberAccountCreationDto : Codable {
    var phoneNumber: String? = nil
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
class PhoneNumberAccountCreationDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var phoneNumber: String? = nil
  @Published var phoneNumberErrorMessage: String? = nil
  func getDto() -> PhoneNumberAccountCreationDto {
      var dto = PhoneNumberAccountCreationDto()
    dto.phoneNumber = self.phoneNumber
      return dto
  }
}
