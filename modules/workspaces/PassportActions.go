package workspaces

var PASSPORT_METHOD_EMAIL = "email"
var PASSPORT_METHOD_PHONE = "phone"

func PassportActionCreate(
	dto *PassportEntity, query QueryDSL,
) (*PassportEntity, *IError) {
	return PassportActionCreateFn(dto, query)
}

func PassportActionUpdate(
	query QueryDSL,
	fields *PassportEntity,
) (*PassportEntity, *IError) {
	return PassportActionUpdateFn(query, fields)
}

func GetUserByPassport2(value string) (*UserEntity, *PassportEntity, *IError) {
	passport := &PassportEntity{}
	err := GetDbRef().Where(&PassportEntity{Value: value}).First(passport).Error

	if err != nil {
		return nil, nil, GormErrorToIError(err)
	}

	user := &UserEntity{}
	GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.String}).First(user)

	return user, passport, nil
}

func GetUserByPassport(value string) (*UserEntity, *PassportEntity, error) {
	passport := &PassportEntity{}
	GetDbRef().Where(&PassportEntity{Value: value}).First(passport)

	user := &UserEntity{}
	GetDbRef().Where(&UserEntity{UniqueId: passport.UserId.String}).First(user)

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
