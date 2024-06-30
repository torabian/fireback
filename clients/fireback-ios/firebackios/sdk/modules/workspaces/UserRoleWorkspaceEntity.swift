import Foundation
class UserRoleWorkspaceEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var role: RoleEntity? = nil
    // var roleId: String? = nil
    var workspace: WorkspaceEntity? = nil
    // var workspaceId: String? = nil
}
class UserRoleWorkspaceEntityViewModel: ObservableObject {
    // improve the fields here
    func getDto() -> UserRoleWorkspaceEntity {
        var dto = UserRoleWorkspaceEntity()
        return dto
    }
}