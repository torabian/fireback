class CommonProfileEntity : Codable {
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
    private var _phoneNumber: String?
    var `PhoneNumber`: String? {
        set { _phoneNumber = newValue }
        get { return _phoneNumber }
    }
    private var _email: String?
    var `Email`: String? {
        set { _email = newValue }
        get { return _email }
    }
    private var _company: String?
    var `Company`: String? {
        set { _company = newValue }
        get { return _company }
    }
    private var _street: String?
    var `Street`: String? {
        set { _street = newValue }
        get { return _street }
    }
    private var _houseNumber: String?
    var `HouseNumber`: String? {
        set { _houseNumber = newValue }
        get { return _houseNumber }
    }
    private var _zipCode: String?
    var `ZipCode`: String? {
        set { _zipCode = newValue }
        get { return _zipCode }
    }
    private var _city: String?
    var `City`: String? {
        set { _city = newValue }
        get { return _city }
    }
    private var _gender: String?
    var `Gender`: String? {
        set { _gender = newValue }
        get { return _gender }
    }
}