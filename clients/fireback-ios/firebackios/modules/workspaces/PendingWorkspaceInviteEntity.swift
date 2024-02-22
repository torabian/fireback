class PendingWorkspaceInviteEntity : Codable {
    private var _value: String?
    var `Value`: String? {
        set { _value = newValue }
        get { return _value }
    }
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
    private var _coverLetter: String?
    var `CoverLetter`: String? {
        set { _coverLetter = newValue }
        get { return _coverLetter }
    }
    private var _workspaceName: String?
    var `WorkspaceName`: String? {
        set { _workspaceName = newValue }
        get { return _workspaceName }
    }
    private var _role: RoleEntity?
    var `Role`: RoleEntity? {
        set { _role = newValue }
        get { return _role }
    }
}