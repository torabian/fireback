import Foundation
struct UserSessionDto : Codable {
    var passport: PassportEntity? = nil
    // var passportId: String? = nil
    var token: String? = nil
    var exchangeKey: String? = nil
    var userWorkspaces: [UserWorkspaceEntity]? = nil
    var userWorkspacesListId: [String]? = nil
    var user: UserEntity? = nil
    // var userId: String? = nil
    var userId: String? = nil
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
class UserSessionDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var token: String? = nil
  @Published var tokenErrorMessage: String? = nil
  @Published var exchangeKey: String? = nil
  @Published var exchangeKeyErrorMessage: String? = nil
  @Published var userId: String? = nil
  @Published var userIdErrorMessage: String? = nil
  func getDto() -> UserSessionDto {
      var dto = UserSessionDto()
    dto.token = self.token
    dto.exchangeKey = self.exchangeKey
    dto.userId = self.userId
      return dto
  }
}
