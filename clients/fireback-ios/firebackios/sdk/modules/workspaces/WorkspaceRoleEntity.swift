import Foundation
class WorkspaceRoleEntity : Codable, Identifiable {
    var userWorkspace: UserWorkspaceEntity? = nil
    // var userWorkspaceId: String? = nil
    var role: RoleEntity? = nil
    // var roleId: String? = nil
}
class WorkspaceRoleEntityViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> WorkspaceRoleEntity {
      var dto = WorkspaceRoleEntity()
      return dto
  }
}