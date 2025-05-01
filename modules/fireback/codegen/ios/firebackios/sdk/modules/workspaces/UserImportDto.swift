import Foundation
struct UserImportPassports : codable {
    var value: String? = nil
    var password: String? = nil
struct UserImportAddress : codable {
    var street: String? = nil
    var zipCode: String? = nil
    var city: String? = nil
    var country: String? = nil
struct UserImportDto : Codable {
    var avatar: String? = nil
    var passports: [UserImportPassports]? = nil
    var person: PersonEntity? = nil
    // var personId: String? = nil
    var address: UserImportAddress? = nil
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
class UserImportDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var avatar: String? = nil
  @Published var avatarErrorMessage: String? = nil
  func getDto() -> UserImportDto {
      var dto = UserImportDto()
    dto.avatar = self.avatar
      return dto
  }
}
