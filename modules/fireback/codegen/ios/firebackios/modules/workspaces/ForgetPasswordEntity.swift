class ForgetPasswordEntity : Codable {
    private var _user: UserEntity?
    var `User`: UserEntity? {
        set { _user = newValue }
        get { return _user }
    }
    private var _passport: PassportEntity?
    var `Passport`: PassportEntity? {
        set { _passport = newValue }
        get { return _passport }
    }
    private var _status: String?
    var `Status`: String? {
        set { _status = newValue }
        get { return _status }
    }
    private var _validUntil: String?
    var `ValidUntil`: String? {
        set { _validUntil = newValue }
        get { return _validUntil }
    }
    private var _blockedUntil: String?
    var `BlockedUntil`: String? {
        set { _blockedUntil = newValue }
        get { return _blockedUntil }
    }
    private var _secondsToUnblock: Int?
    var `SecondsToUnblock`: Int? {
        set { _secondsToUnblock = newValue }
        get { return _secondsToUnblock }
    }
    private var _otp: String?
    var `Otp`: String? {
        set { _otp = newValue }
        get { return _otp }
    }
    private var _recoveryAbsoluteUrl: String?
    var `RecoveryAbsoluteUrl`: String? {
        set { _recoveryAbsoluteUrl = newValue }
        get { return _recoveryAbsoluteUrl }
    }
}