class WorkspaceConfigEntity : Codable {
    private var _disablePublicWorkspaceCreation: Int?
    var `DisablePublicWorkspaceCreation`: Int? {
        set { _disablePublicWorkspaceCreation = newValue }
        get { return _disablePublicWorkspaceCreation }
    }
}