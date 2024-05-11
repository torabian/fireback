import Foundation
struct AuthResultDto : Codable {
    var workspaceId: String? = nil
    var userRoleWorkspacePermissions: [UserRoleWorkspacePermissionDto]? = nil
    var userRoleWorkspacePermissionsListId: [String]? = nil
    var internalSql: String? = nil
    var userId: String? = nil
    var userHas: [String]? = nil
    var workspaceHas: [String]? = nil
    var user: UserEntity? = nil
    // var userId: String? = nil
    var accessLevel: UserAccessLevelDto? = nil
    // var accessLevelId: String? = nil
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
class AuthResultDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var workspaceId: String? = nil
  @Published var workspaceIdErrorMessage: String? = nil
  @Published var internalSql: String? = nil
  @Published var internalSqlErrorMessage: String? = nil
  @Published var userId: String? = nil
  @Published var userIdErrorMessage: String? = nil
  func getDto() -> AuthResultDto {
      var dto = AuthResultDto()
    dto.workspaceId = self.workspaceId
    dto.internalSql = self.internalSql
    dto.userId = self.userId
      return dto
  }
}
