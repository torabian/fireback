import Foundation
struct UserRoleWorkspacePermissionDto : Codable {
    var workspaceId: String? = nil
    var userId: String? = nil
    var roleId: String? = nil
    var capabilityId: String? = nil
    var type: String? = nil
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
class UserRoleWorkspacePermissionDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var workspaceId: String? = nil
  @Published var workspaceIdErrorMessage: String? = nil
  @Published var userId: String? = nil
  @Published var userIdErrorMessage: String? = nil
  @Published var roleId: String? = nil
  @Published var roleIdErrorMessage: String? = nil
  @Published var capabilityId: String? = nil
  @Published var capabilityIdErrorMessage: String? = nil
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  func getDto() -> UserRoleWorkspacePermissionDto {
      var dto = UserRoleWorkspacePermissionDto()
    dto.workspaceId = self.workspaceId
    dto.userId = self.userId
    dto.roleId = self.roleId
    dto.capabilityId = self.capabilityId
    dto.type = self.type
      return dto
  }
}
