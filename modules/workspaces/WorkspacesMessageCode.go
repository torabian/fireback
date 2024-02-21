package workspaces

var WorkspacesMessageCode = newWorkspacesMessageCode()

func newWorkspacesMessageCode() *workspaceReqmsg {
	return &workspaceReqmsg{
		InviteToWorkspaceMailSenderMissing: "InviteToWorkspaceMailSenderMissing",
		UserWhichHasThisTokenDoesNotExist:  "UserWhichHasThisTokenDoesNotExist",
		ProvideTokenInAuthorization:        "ProvideTokenInAuthorization",
		UserNotFoundOrDeleted:              "UserNotFoundOrDeleted",
		SelectWorkspaceId:                  "SelectWorkspaceId",
		PassportNotAvailable:               "PassportNotAvailable",
		GsmConfigurationIsNotAvailable:     "GsmConfigurationIsNotAvailable",
		EmailConfigurationIsNotAvailable:   "EmailConfigurationIsNotAvailable",
		PassportUserNotAvailable:           "PassportUserNotAvailable",
	}
}

type workspaceReqmsg struct {
	InviteToWorkspaceMailSenderMissing string
	UserWhichHasThisTokenDoesNotExist  string
	ProvideTokenInAuthorization        string
	UserNotFoundOrDeleted              string
	GsmConfigurationIsNotAvailable     string
	EmailConfigurationIsNotAvailable   string
	SelectWorkspaceId                  string
	PassportNotAvailable               string
	PassportUserNotAvailable           string
}
