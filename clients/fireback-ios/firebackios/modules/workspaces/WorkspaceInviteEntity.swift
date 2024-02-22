class WorkspaceInviteEntity : Codable {
    private var _passportMode: String?
    var `PassportMode`: String? {
        set { _passportMode = newValue }
        get { return _passportMode }
    }
    private var _coverLetter: String?
    var `CoverLetter`: String? {
        set { _coverLetter = newValue }
        get { return _coverLetter }
    }
    private var _targetUserLocale: String?
    var `TargetUserLocale`: String? {
        set { _targetUserLocale = newValue }
        get { return _targetUserLocale }
    }
    private var _email: String?
    var `Email`: String? {
        set { _email = newValue }
        get { return _email }
    }
    private var _workspace: WorkspaceEntity?
    var `Workspace`: WorkspaceEntity? {
        set { _workspace = newValue }
        get { return _workspace }
    }
    private var _firstName: String?
    var `FirstName`: String? {
        set { _firstName = newValue }
        get { return _firstName }
    }
    private var _lastName: String?
    var `LastName`: String? {
        set { _lastName = newValue }
        get { return _lastName }
    }
    private var _inviteeUserId: String?
    var `InviteeUserId`: String? {
        set { _inviteeUserId = newValue }
        get { return _inviteeUserId }
    }
    private var _phoneNumber: String?
    var `PhoneNumber`: String? {
        set { _phoneNumber = newValue }
        get { return _phoneNumber }
    }
    private var _used: Bool?
    var `Used`: Bool? {
        set { _used = newValue }
        get { return _used }
    }
    private var _role: RoleEntity?
    var `Role`: RoleEntity? {
        set { _role = newValue }
        get { return _role }
    }
}