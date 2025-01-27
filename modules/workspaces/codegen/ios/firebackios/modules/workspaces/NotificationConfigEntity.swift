class NotificationConfigEntity : Codable {
    private var _cascadeToSubWorkspaces: Bool?
    var `CascadeToSubWorkspaces`: Bool? {
        set { _cascadeToSubWorkspaces = newValue }
        get { return _cascadeToSubWorkspaces }
    }
    private var _forcedCascadeEmailProvider: Bool?
    var `ForcedCascadeEmailProvider`: Bool? {
        set { _forcedCascadeEmailProvider = newValue }
        get { return _forcedCascadeEmailProvider }
    }
    private var _generalEmailProvider: EmailProviderEntity?
    var `GeneralEmailProvider`: EmailProviderEntity? {
        set { _generalEmailProvider = newValue }
        get { return _generalEmailProvider }
    }
    private var _generalGsmProvider: GsmProviderEntity?
    var `GeneralGsmProvider`: GsmProviderEntity? {
        set { _generalGsmProvider = newValue }
        get { return _generalGsmProvider }
    }
    private var _inviteToWorkspaceContent: String?
    var `InviteToWorkspaceContent`: String? {
        set { _inviteToWorkspaceContent = newValue }
        get { return _inviteToWorkspaceContent }
    }
    private var _inviteToWorkspaceContentExcerpt: String?
    var `InviteToWorkspaceContentExcerpt`: String? {
        set { _inviteToWorkspaceContentExcerpt = newValue }
        get { return _inviteToWorkspaceContentExcerpt }
    }
    private var _inviteToWorkspaceContentDefault: String?
    var `InviteToWorkspaceContentDefault`: String? {
        set { _inviteToWorkspaceContentDefault = newValue }
        get { return _inviteToWorkspaceContentDefault }
    }
    private var _inviteToWorkspaceContentDefaultExcerpt: String?
    var `InviteToWorkspaceContentDefaultExcerpt`: String? {
        set { _inviteToWorkspaceContentDefaultExcerpt = newValue }
        get { return _inviteToWorkspaceContentDefaultExcerpt }
    }
    private var _inviteToWorkspaceTitle: String?
    var `InviteToWorkspaceTitle`: String? {
        set { _inviteToWorkspaceTitle = newValue }
        get { return _inviteToWorkspaceTitle }
    }
    private var _inviteToWorkspaceTitleDefault: String?
    var `InviteToWorkspaceTitleDefault`: String? {
        set { _inviteToWorkspaceTitleDefault = newValue }
        get { return _inviteToWorkspaceTitleDefault }
    }
    private var _inviteToWorkspaceSender: EmailSenderEntity?
    var `InviteToWorkspaceSender`: EmailSenderEntity? {
        set { _inviteToWorkspaceSender = newValue }
        get { return _inviteToWorkspaceSender }
    }
    private var _forgetPasswordContent: String?
    var `ForgetPasswordContent`: String? {
        set { _forgetPasswordContent = newValue }
        get { return _forgetPasswordContent }
    }
    private var _forgetPasswordContentExcerpt: String?
    var `ForgetPasswordContentExcerpt`: String? {
        set { _forgetPasswordContentExcerpt = newValue }
        get { return _forgetPasswordContentExcerpt }
    }
    private var _forgetPasswordContentDefault: String?
    var `ForgetPasswordContentDefault`: String? {
        set { _forgetPasswordContentDefault = newValue }
        get { return _forgetPasswordContentDefault }
    }
    private var _forgetPasswordContentDefaultExcerpt: String?
    var `ForgetPasswordContentDefaultExcerpt`: String? {
        set { _forgetPasswordContentDefaultExcerpt = newValue }
        get { return _forgetPasswordContentDefaultExcerpt }
    }
    private var _forgetPasswordTitle: String?
    var `ForgetPasswordTitle`: String? {
        set { _forgetPasswordTitle = newValue }
        get { return _forgetPasswordTitle }
    }
    private var _forgetPasswordTitleDefault: String?
    var `ForgetPasswordTitleDefault`: String? {
        set { _forgetPasswordTitleDefault = newValue }
        get { return _forgetPasswordTitleDefault }
    }
    private var _forgetPasswordSender: EmailSenderEntity?
    var `ForgetPasswordSender`: EmailSenderEntity? {
        set { _forgetPasswordSender = newValue }
        get { return _forgetPasswordSender }
    }
    private var _acceptLanguage: String?
    var `AcceptLanguage`: String? {
        set { _acceptLanguage = newValue }
        get { return _acceptLanguage }
    }
    private var _confirmEmailSender: EmailSenderEntity?
    var `ConfirmEmailSender`: EmailSenderEntity? {
        set { _confirmEmailSender = newValue }
        get { return _confirmEmailSender }
    }
    private var _confirmEmailContent: String?
    var `ConfirmEmailContent`: String? {
        set { _confirmEmailContent = newValue }
        get { return _confirmEmailContent }
    }
    private var _confirmEmailContentExcerpt: String?
    var `ConfirmEmailContentExcerpt`: String? {
        set { _confirmEmailContentExcerpt = newValue }
        get { return _confirmEmailContentExcerpt }
    }
    private var _confirmEmailContentDefault: String?
    var `ConfirmEmailContentDefault`: String? {
        set { _confirmEmailContentDefault = newValue }
        get { return _confirmEmailContentDefault }
    }
    private var _confirmEmailContentDefaultExcerpt: String?
    var `ConfirmEmailContentDefaultExcerpt`: String? {
        set { _confirmEmailContentDefaultExcerpt = newValue }
        get { return _confirmEmailContentDefaultExcerpt }
    }
    private var _confirmEmailTitle: String?
    var `ConfirmEmailTitle`: String? {
        set { _confirmEmailTitle = newValue }
        get { return _confirmEmailTitle }
    }
    private var _confirmEmailTitleDefault: String?
    var `ConfirmEmailTitleDefault`: String? {
        set { _confirmEmailTitleDefault = newValue }
        get { return _confirmEmailTitleDefault }
    }
}