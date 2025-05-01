class TokenEntity : Codable {
    private var _user: UserEntity?
    var `User`: UserEntity? {
        set { _user = newValue }
        get { return _user }
    }
    private var _validUntil: String?
    var `ValidUntil`: String? {
        set { _validUntil = newValue }
        get { return _validUntil }
    }
}