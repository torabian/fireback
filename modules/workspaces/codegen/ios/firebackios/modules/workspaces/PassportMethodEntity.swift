class PassportMethodEntity : Codable {
    private var _name: String?
    var `Name`: String? {
        set { _name = newValue }
        get { return _name }
    }
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
    private var _region: String?
    var `Region`: String? {
        set { _region = newValue }
        get { return _region }
    }
}