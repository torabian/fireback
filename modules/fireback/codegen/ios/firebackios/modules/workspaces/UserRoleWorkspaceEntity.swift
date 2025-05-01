class UserRoleWorkspaceEntity : Codable {
    private var _user: UserEntity?
    var `User`: UserEntity? {
        set { _user = newValue }
        get { return _user }
    }
    private var _role: RoleEntity?
    var `Role`: RoleEntity? {
        set { _role = newValue }
        get { return _role }
    }
    private var _workspace: WorkspaceEntity?
    var `Workspace`: WorkspaceEntity? {
        set { _workspace = newValue }
        get { return _workspace }
    }
}