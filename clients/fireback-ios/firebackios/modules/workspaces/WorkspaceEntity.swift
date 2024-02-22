class WorkspaceEntity : Codable {
    private var _description: String?
    var `Description`: String? {
        set { _description = newValue }
        get { return _description }
    }
    private var _name: String?
    var `Name`: String? {
        set { _name = newValue }
        get { return _name }
    }
}