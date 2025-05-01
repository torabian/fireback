import Foundation
class PublicJoinKeyEntity : Codable, Identifiable {
    var role: RoleEntity? = nil
    // var roleId: String? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
}
class PublicJoinKeyEntityViewModel: ObservableObject {
  // improve the fields here
  func getDto() -> PublicJoinKeyEntity {
      var dto = PublicJoinKeyEntity()
      return dto
  }
}