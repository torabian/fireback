import Foundation
class UserWorkspaceEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
    var userPermissions: [String]? = nil
    var rolePermission: [UserRoleWorkspaceDto]? = nil
    var workspacePermissions: [String]? = nil
}
class UserWorkspaceEntityViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> UserWorkspaceEntity {
      var dto = UserWorkspaceEntity()
      return dto
  }
}