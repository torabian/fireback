class RoleEntity : Codable {
    private var _name: String?
    var `Name`: String? {
        set { _name = newValue }
        get { return _name }
    }
    private var _capabilities: [[CapabilityEntity]]?
    var `Capabilities`: [[CapabilityEntity]]? {
        set { _capabilities = newValue }
        get { return _capabilities }
    }
}