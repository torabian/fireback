class PublicJoinKeyEntity : Codable {
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