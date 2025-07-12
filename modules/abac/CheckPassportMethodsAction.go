package abac

import "github.com/torabian/fireback/modules/fireback"

func init() {
	// Override the implementation with our actual code.
	CheckPassportMethodsActionImp = CheckPassportMethodsAction
}

func CheckPassportMethodsAction(q fireback.QueryDSL) (*CheckPassportMethodsActionResDto, *fireback.IError) {
	state := &CheckPassportMethodsActionResDto{}

	// Get the workspacec configuration as well, for different reasons such as captcha info
	config, err2 := WorkspaceConfigActions.GetByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})
	if err2 != nil {
		if err2.HttpCode != 404 {
			return nil, err2
		}
	}

	if config != nil {
		state.EnabledRecaptcha2 = config.EnableRecaptcha2.Bool && config.Recaptcha2ClientKey != "" && config.Recaptcha2ServerKey != ""
		state.Recaptcha2ClientKey = config.Recaptcha2ClientKey
	}

	// This can be implemented at some point to detect from user IP or other services
	// that where he/she is. Then, we can narrow down the methods available for that specific region
	// For example, there is a global option to use phone number, but for Malaysian users we want
	// to use a different operator to send text messages
	userDetectedRegion := "global"

	// Known unsafe operation. We need all the records in the database, in order
	// to determine the best authentication option for the user.
	// configuration field are not being returned publicly, only the final state.
	stream, _, err := PassportMethodEntityStream(fireback.QueryDSL{})
	if err != nil {
		return nil, fireback.CastToIError(err)
	}

	for items := range stream {
		for _, item := range items {
			if item.Region == "" || item.Type == "" {
				continue
			}
			region := item.Region
			Type := item.Type

			// This logic has issue if the global is changed.
			// We need to always set the condition based on hierachy on real scenario
			// We only support global now, to keep the interactions compatible with future.
			if region != userDetectedRegion {
				continue
			}

			if Type == PassportMethodType.Email {
				state.Email = true
			}

			if Type == PassportMethodType.Phone {
				state.Phone = true
			}

			if Type == PassportMethodType.Google {
				state.Google = true
				state.GoogleOAuthClientKey = item.ClientKey
			}

			if Type == PassportMethodType.Facebook {
				state.Facebook = true
				state.FacebookAppId = item.ClientKey
			}
		}
	}

	return state, nil
}
