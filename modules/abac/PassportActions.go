package abac

import "github.com/torabian/fireback/modules/fireback"

var PASSPORT_METHOD_EMAIL = "email"
var PASSPORT_METHOD_PHONE = "phone"

func GetUserByPassport2(value string) (*UserEntity, *PassportEntity, *fireback.IError) {
	passport := &PassportEntity{}
	err := fireback.GetDbRef().Where(&PassportEntity{Value: value}).First(passport).Error

	if err != nil {
		return nil, nil, fireback.GormErrorToIError(err)
	}

	user := &UserEntity{}
	fireback.GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.String}).First(user)

	return user, passport, nil
}

func GetUserByPassport(value string) (*UserEntity, *PassportEntity, error) {
	passport := &PassportEntity{}
	fireback.GetDbRef().Where(&PassportEntity{Value: value}).First(passport)

	user := &UserEntity{}
	fireback.GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.String}).First(user)

	return user, passport, nil
}

// func SendEmailFromTemplate(templateName string, dto *ForgetPasswordEntity) error {
// 	var config = GetAppConfig()

// 	cfg := config.MailTemplates.ForgetPasswordRequest

// 	if !cfg.Enabled {
// 		fmt.Println("Email server is not available, so this won't work. config.MailTemplates.ForgetPasswordRequest.Enabled = false")
// 		return errors.New(PassportMessageCode.EmailIsNotConfigured)
// 	}

// 	t, err := template.ParseFS(templates.PassportEmailTemplates, templateName)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var tpl bytes.Buffer
// 	err = t.Execute(&tpl, dto)
// 	if err != nil {
// 		panic(err)
// 	}

// 	result := tpl.String()

// 	return notification.SendMail(
// 		notification.EmailMessageContent{
// 			ToName:  dto.User.FirstName + " " + dto.User.LastName,
// 			ToEmail: dto.Passport.Value,
// 			Content: result,
// 			Subject: cfg.Subject,
// 		},
// 	)

// }

// func SigninUserWithGoogleClaim(user UserEntity, claim GoogleAuthClaim) (string, error) {

// 	fmt.Println(claim)
// 	tokenString := x(32)

// 	GetDbRef().Create(&Token{
// 		Hash:       tokenString,
// 		UserId:     user.UniqueId,
// 		UniqueId:   UUID(),
// 		ValidUntil: time.Now().Add(time.Hour * time.Duration(12)),
// 	})

// 	return tokenString, nil

// }
