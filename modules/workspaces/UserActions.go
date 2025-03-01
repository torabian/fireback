package workspaces

/*
This file holds in memory some temporary solution to tokens that need to be used upon redirection
They can be stored, and accessed only once, so third party can get the token in the other domain
*/

type exchangeItem struct {
	token string
}

var exchangePool map[string]exchangeItem = map[string]exchangeItem{}

func PutTokenInExchangePool(token string) string {
	uniqueId := UUID()

	exchangePool[uniqueId] = exchangeItem{
		token: token,
	}

	return uniqueId
}

func GetTokenFromExchangePoolAction(query QueryDSL) (*ExchangeKeyInformationDto, *IError) {
	item := exchangePool[query.UniqueId]
	token := item.token
	delete(exchangePool, query.UniqueId)

	if token == "" {
		return nil, Create401Error(&WorkspacesMessages.InvalidExchangeKey, []string{})
	}

	return &ExchangeKeyInformationDto{Key: token}, nil
}

/*
1- Maybe user does not exists in the system, so we create a new user, and give him a google passport
2- Maybe user has an account, and has linked google account, so we just act as signin
3- Maybe user email is registered, but it's not linked with the account. Now we give error,
4- To first login and then come back.

type GoogleAuthClaim struct {
	Aud           string `json:"aud"`
	Azp           string `json:"azp"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Exp           uint   `json:"exp"`
	FamilyName    string `json:"family_name"`
	GivenName     string `json:"given_name"`
	Iat           uint   `json:"iat"`
	Iss           string `json:"iss"`
	Jti           string `json:"jti"`
	Locale        string `json:"locale"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	Sub           string `json:"sub"`
}
BROKEN!!
*/
// func SigninUserWithGoogle(claim GoogleAuthClaim, accessToken string) (UserEntity, string, error) {

// 	u := UUID()

// 	var p2 models.Passport

// 	dbref.Preload("User").Where("type = ? and email = ?", models.PassportTypes.Google, claim.Email).First(&p2)

// 	var user UserEntity

// 	if p2.UniqueId == "" {
// 		uid := u.String()
// 		user = UserEntity{
// 			UniqueId:  uid,
// 			FirstName: claim.GivenName,
// 			Lastname:  claim.FamilyName,
// 			Photo:     "googleoauth2://" + claim.Picture,
// 			// Passports: []models.Passport{
// 			// 	{Email: claim.Email, Password: "", Type: models.PassportTypes.Google, AccessToken: accessToken, UserID: uid},
// 			// },
// 		}

// 		err := dbref.Create(&user).Error

// 		if err != nil {
// 			return UserEntity{}, "", err
// 		}

// 		// 3. Create the workspace for him
// 		// 4. Assign his owner role
// 		CreateWorkspace(user.UniqueId, claim.Email+" workspace")

// 	} else {

// 		fmt.Println("++++++")

// 		dbref.Model(&models.Passport{}).Where(models.Passport{Type: "Google", UserID: p2.UserID}).Updates(models.Passport{AccessToken: accessToken})
// 	}

// 	// 5. Log him in
// 	token, _ := SigninUserWithGoogleClaim(user, claim)

// 	return user, token, nil
// }

// func GetUserPreferences(user *UserEntity) map[string]interface{} {
// 	body := map[string]interface{}{}

// 	var items []*Preference
// 	GetDbRef().Where("user_id = ?", user.UniqueId).Find(&items)

// 	for _, item := range items {
// 		if item.ValueType == "float" {
// 			body[item.ItemKey], _ = strconv.Atoi(item.Value)
// 		} else if item.ValueType == "boolean" && item.Value == "true" {
// 			body[item.ItemKey] = true
// 		} else if item.ValueType == "boolean" && item.Value == "false" {
// 			body[item.ItemKey] = false
// 		} else {
// 			body[item.ItemKey] = item.Value
// 		}

// 	}

// 	return body
// }

// func PatchUserPreferences(user *UserEntity, data map[string]interface{}) {

// 	for k, v := range data {

// 		valueType := "string"

// 		switch v.(type) {

// 		case int:
// 			valueType = "number"
// 		case float64:
// 			valueType = "float"
// 		case bool:
// 			valueType = "boolean"
// 		}

// 		if v == "true" {
// 			valueType = "boolean"
// 		}

// 		if v == "false" {
// 			valueType = "boolean"
// 		}

// 		if GetDbRef().Model(&Preference{}).Where(&Preference{UserID: &user.UniqueId, ItemKey: k}).Update("value", v).RowsAffected == 0 {
// 			GetDbRef().Create(&Preference{
// 				User:      user,
// 				ItemKey:   k,
// 				ValueType: valueType,
// 				Value:     fmt.Sprint(v),
// 			})
// 		}

// 	}

// }

func GetUserFromToken(tokenString string) (*UserEntity, error) {

	var item TokenEntity

	if err := GetDbRef().Preload("User").Where(RealEscape("token = ?", tokenString)).First(&item).Error; err != nil {
		return &UserEntity{}, err
	}

	user, _ := UserActionGetOne(QueryDSL{UniqueId: item.UserId.String})
	return user, nil
}

func UserActionCreate(
	dto *UserEntity, query QueryDSL,
) (*UserEntity, *IError) {
	query.WorkspaceId = "root"
	return UserActionCreateFn(dto, query)
}

func UserActionUpdate(
	query QueryDSL,
	fields *UserEntity,
) (*UserEntity, *IError) {
	return UserActionUpdateFn(query, fields)
}

// func UserActionQuery(query QueryDSL) ([]*UserEntity, *QueryResultMeta, error) {

// 	result, qrm, err := UnsafeQuerySqlFromFs[UserEntity](
// 		&queries.QueriesFs, "queryUsers", query,
// 	)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return result, qrm, err
// }
