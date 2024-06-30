class UserEntity : Codable {
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
    private var _photo: String?
    var `Photo`: String? {
        set { _photo = newValue }
        get { return _photo }
    }
}