class CapabilityEntity : Codable {
    private var _name: String?
    var `Name`: String? {
        set { _name = newValue }
        get { return _name }
    }
}