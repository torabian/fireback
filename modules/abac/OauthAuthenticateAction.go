package abac

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
)

// Supported OAuth providers
const (
	ProviderGoogle   = "google"
	ProviderFacebook = "facebook"
)

// TokenInfo is reused for both Google and Facebook responses
type TokenInfo struct {
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"` // Facebook doesn't provide this, but keep for compatibility
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

	if passport == nil {

		firstName, lastName := SplitName(info.Name, info.Email)
		res, err := completeClassicSignupProcess(&ClassicSignupActionReqDto{
			Value:     info.Email,
			Type:      PASSPORT_METHOD_EMAIL,
			FirstName: firstName,
			LastName:  lastName,
		}, q, nil, nil, func(ue *UserEntity, re *RoleEntity, we *WorkspaceEntity, pe *PassportEntity) {
			pe.ThirdPartyVerifier = provider
			we.Name = "My workspace"
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

// authenticateWithFacebook verifies the Facebook access token and returns user info
func authenticateWithFacebook(accessToken string, q fireback.QueryDSL) (*OauthAuthenticateActionResDto, *fireback.IError) {
	url := fmt.Sprintf("https://graph.facebook.com/me?fields=email,name,picture&access_token=%s", accessToken)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return nil, fireback.Create401Error(&AbacMessages.InvalidToken, []string{})
	}
	defer resp.Body.Close()

	var fbData struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Picture struct {
			Data struct {
				URL string `json:"url"`
			} `json:"data"`
		} `json:"picture"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&fbData); err != nil || fbData.Email == "" {
		return nil, fireback.Create401Error(&AbacMessages.FailedToDecodeGoogle, []string{})
	}

	tokenInfo := TokenInfo{
		Name:    fbData.Name,
		Picture: fbData.Picture.Data.URL,
		Email:   fbData.Email,
	}

	return continueAuthenticationViaOAuthEmail(tokenInfo, ProviderFacebook, q)
}

func init() {
	OauthAuthenticateActionImp = OauthAuthenticateAction
}

func SplitName(fullName, email string) (firstName, lastName string) {
	// 1. Try to split full name
	parts := strings.Fields(fullName)
	if len(parts) >= 2 {
		return parts[0], strings.Join(parts[1:], " ")
	}
	if len(parts) == 1 {
		return parts[0], "n/a"
	}

	// 2. Try to extract name from email
	if email != "" {
		local := strings.SplitN(email, "@", 2)[0]
		if strings.Contains(local, ".") {
			p := strings.SplitN(local, ".", 2)
			return p[0], p[1]
		}
		return local, "n/a"
	}

	// 3. Absolute fallback
	return "user", "n/a"
}
