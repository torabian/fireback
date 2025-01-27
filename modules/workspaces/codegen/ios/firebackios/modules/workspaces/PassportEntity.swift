class PassportEntity : Codable {
    private var _type: String?
    var `Type`: String? {
        set { _type = newValue }
        get { return _type }
    }
    private var _value: String?
    var `Value`: String? {
        set { _value = newValue }
        get { return _value }
    }
    private var _password: String?
    var `Password`: String? {
        set { _password = newValue }
        get { return _password }
    }
    private var _confirmed: Bool?
    var `Confirmed`: Bool? {
        set { _confirmed = newValue }
        get { return _confirmed }
    }
    private var _accessToken: String?
    var `AccessToken`: String? {
        set { _accessToken = newValue }
        get { return _accessToken }
    }
}