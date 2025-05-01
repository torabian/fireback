package abac

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/torabian/fireback/modules/fireback"
)

// Supported OAuth providers
const (
	ProviderGoogle   = "google"
	ProviderFacebook = "facebook"
)

// GoogleTokenInfo represents the response from Google's token validation API
type TokenInfo struct {
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Scope         string `json:"scope"`
	Exp           string `json:"exp"`
	ExpiresIn     string `json:"expires_in"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Access_type   string `json:"access_type"`
}

// OauthAuthenticateAction authenticates a user via OAuth
func OauthAuthenticateAction(
	req *OauthAuthenticateActionReqDto,
	q fireback.QueryDSL) (*OauthAuthenticateActionResDto, *fireback.IError) {

	switch req.Service {
	case ProviderGoogle:
		return authenticateWithGoogle(req.Token, q)
	case ProviderFacebook:
		return authenticateWithFacebook(req.Token, q)
	default:
		return nil, fireback.Create401Error(&AbacMessages.UnsupportedOAuth, []string{})
	}
}

func continueAuthenticationViaOAuthEmail(info TokenInfo, provider string, q fireback.QueryDSL) (*OauthAuthenticateActionResDto, *fireback.IError) {

	if err := validateValueFormat(info.Email); err != nil {
		return nil, err
	}

	ClearPassportValue(&info.Email)

	passport := findPassport(info.Email, q)

	// from here we devide the work flow to existing and non exists passport
	if passport == nil {
		res, err := completeClassicSignupProcess(&ClassicSignupActionReqDto{
			Value:     info.Email,
			Type:      PASSPORT_METHOD_EMAIL,
			FirstName: "Guest",
			LastName:  "User",
		}, q, nil, nil, func(ue *UserEntity, re *RoleEntity, we *WorkspaceEntity, pe *PassportEntity) {

			// This is important
			pe.ThirdPartyVerifier = provider
		})

		if err != nil {
			return nil, err
		}

		return &OauthAuthenticateActionResDto{
			Session: res.Session,
		}, nil
	} else {
		session := &UserSessionDto{}

		if _, err := fetchUserAndPassToSession(info.Email, session, q); err != nil {
			return nil, err
		}

		if err := applyUserTokenAndWorkspacesToToken(session, q); err != nil {
			return nil, err
		}

		return &OauthAuthenticateActionResDto{
			Session: session,
		}, nil
	}
}

// authenticateWithGoogle verifies the Google access token and returns user info
func authenticateWithGoogle(accessToken string, q fireback.QueryDSL) (*OauthAuthenticateActionResDto, *fireback.IError) {
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fireback.Create401Error(&AbacMessages.InvalidToken, []string{})
	}
	defer resp.Body.Close()

	var tokenInfo TokenInfo
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil || tokenInfo.Email == "" {
		return nil, fireback.Create401Error(&AbacMessages.FailedToDecodeGoogle, []string{})
	}

	return continueAuthenticationViaOAuthEmail(tokenInfo, ProviderGoogle, q)
}

func authenticateWithFacebook(accessToken string, q fireback.QueryDSL) (*OauthAuthenticateActionResDto, *fireback.IError) {
	// TODO: Implement Facebook token validation
	return nil, nil
}

func init() {
	OauthAuthenticateActionImp = OauthAuthenticateAction
}
