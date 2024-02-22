class EmailSenderEntity : Codable {
    private var _fromName: String?
    var `FromName`: String? {
        set { _fromName = newValue }
        get { return _fromName }
    }
    private var _fromEmailAddress: String?
    var `FromEmailAddress`: String? {
        set { _fromEmailAddress = newValue }
        get { return _fromEmailAddress }
    }
    private var _replyTo: String?
    var `ReplyTo`: String? {
        set { _replyTo = newValue }
        get { return _replyTo }
    }
    private var _nickName: String?
    var `NickName`: String? {
        set { _nickName = newValue }
        get { return _nickName }
    }
}