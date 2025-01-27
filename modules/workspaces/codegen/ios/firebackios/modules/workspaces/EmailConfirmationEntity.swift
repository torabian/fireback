class EmailConfirmationEntity : Codable {
    private var _user: UserEntity?
    var `User`: UserEntity? {
        set { _user = newValue }
        get { return _user }
    }
    private var _status: String?
    var `Status`: String? {
        set { _status = newValue }
        get { return _status }
    }
    private var _email: String?
    var `Email`: String? {
        set { _email = newValue }
        get { return _email }
    }
    private var _key: String?
    var `Key`: String? {
        set { _key = newValue }
        get { return _key }
    }
    private var _expiresAt: String?
    var `ExpiresAt`: String? {
        set { _expiresAt = newValue }
        get { return _expiresAt }
    }
}