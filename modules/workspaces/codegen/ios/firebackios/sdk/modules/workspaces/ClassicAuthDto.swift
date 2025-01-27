import Foundation
struct ClassicAuthDto : Codable {
    var value: String? = nil
    var password: String? = nil
    var firstName: String? = nil
    var lastName: String? = nil
    var inviteId: String? = nil
    var publicJoinKeyId: String? = nil
    var workspaceTypeId: String? = nil
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
class ClassicAuthDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  @Published var firstName: String? = nil
  @Published var firstNameErrorMessage: String? = nil
  @Published var lastName: String? = nil
  @Published var lastNameErrorMessage: String? = nil
  @Published var inviteId: String? = nil
  @Published var inviteIdErrorMessage: String? = nil
  @Published var publicJoinKeyId: String? = nil
  @Published var publicJoinKeyIdErrorMessage: String? = nil
  @Published var workspaceTypeId: String? = nil
  @Published var workspaceTypeIdErrorMessage: String? = nil
  func getDto() -> ClassicAuthDto {
      var dto = ClassicAuthDto()
    dto.value = self.value
    dto.password = self.password
    dto.firstName = self.firstName
    dto.lastName = self.lastName
    dto.inviteId = self.inviteId
    dto.publicJoinKeyId = self.publicJoinKeyId
    dto.workspaceTypeId = self.workspaceTypeId
      return dto
  }
}
