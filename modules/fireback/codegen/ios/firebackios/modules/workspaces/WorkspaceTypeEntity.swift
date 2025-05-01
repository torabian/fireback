class WorkspaceTypeEntity : Codable {
    private var _title: String?
    var `Title`: String? {
        set { _title = newValue }
        get { return _title }
    }
    private var _description: String?
    var `Description`: String? {
        set { _description = newValue }
        get { return _description }
    }
    private var _slug: String?
    var `Slug`: String? {
        set { _slug = newValue }
        get { return _slug }
    }
    private var _role: RoleEntity?
    var `Role`: RoleEntity? {
        set { _role = newValue }
        get { return _role }
    }
}