package workspaces

var PassportMessageCode = newPassportMessageCode()

func newPassportMessageCode() *passportMessageCodeType {
	return &passportMessageCodeType{
		UserDoesNotExist:       "UserDoesNotExist",
		AlreadyConfirmed:       "AlreadyConfirmed",
		EmailNotFound:          "EmailNotFound",
		InvitationExpired:      "InvitationExpired",
		PasswordRequired:       "PasswordRequired",
		PassportNotAvailable:   "PassportNotAvailable",
		ResetNotFound:          "ResetNotFound",
		PASSPORT_NOT_FOUND:     "PASSPORT_NOT_FOUND",
		OTARequestBlockedUntil: "OTARequestBlockedUntil",
		EmailIsNotConfigured:   "EmailIsNotConfigured",
		OtpCodeInvalid:         "OtpCodeInvalid",
	}
}

type passportMessageCodeType struct {
	UserDoesNotExist       string
	AlreadyConfirmed       string
	EmailNotFound          string
	PassportNotAvailable   string
	InvitationExpired      string
	PasswordRequired       string
	PASSPORT_NOT_FOUND     string
	ResetNotFound          string
	OTARequestBlockedUntil string
	EmailIsNotConfigured   string
	OtpCodeInvalid         string
}
