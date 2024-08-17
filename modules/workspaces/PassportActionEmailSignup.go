package workspaces

import (
	"gorm.io/gorm"
)

func PassportAppendEmailToUser(dto *ClassicAuthDto, query QueryDSL) (*PassportEntity, *IError) {
	passwordHashed, _ := HashPassword(*dto.Password)

	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return nil, iError
	}

	entity := &PassportEntity{
		Value: dto.Value,
		// Confirmed: 1,
		Password: &passwordHashed,
		Type:     &PassportTypes.EmailPassword,
	}

	return PassportActionCreateFn(entity, query)

}

func PassportActionEmailSignup(dto *ClassicAuthDto, query QueryDSL) (*UserSessionDto, *IError) {

	session := &UserSessionDto{}
	if iError := UserWithEmailAndPasswordValidator(dto, false); iError != nil {
		return session, iError
	}

	user, passport, token, err := CreateUserWithEmailAndPassword(
		dto,
		query,
	)

	if dto.InviteId != nil {
		query.UserId = user.UniqueId
		AcceptInvitationAction(&AcceptInviteDto{
			InviteUniqueId: dto.InviteId,
		}, query)
	}

	if err != nil {
		return nil, err
	}

	// workspacesList := GetUserWorkspaces(user.UniqueId)

	// var config = GetAppConfig()
	// if config.MailTemplates.ConfirmMail.Enabled {
	// 	SendUserMailConfirmation(dto.Email, *user)
	// }

	session.Passport = passport
	session.User = user
	session.Token = &token
	// session.UserRoleWorkspaces = workspacesList
	ek := PutTokenInExchangePool(token)
	session.ExchangeKey = &ek

	return session, nil
}

func CreateEmailPassportForUser(userId string, email string, password string, tx *gorm.DB) (*PassportEntity, error) {
	passwordHashed, _ := HashPassword(password)

	var passport = &PassportEntity{
		Value:    &email,
		Password: &passwordHashed,
		Type:     &PassportTypes.EmailPassword,
		UserId:   &userId,
		UniqueId: UUID(),
	}

	if tx == nil {
		tx = GetDbRef()
	}

	if err := tx.Create(&passport).Error; err != nil {
		return nil, err
	} else {
		return passport, nil
	}
}

func CreateUserWithEmailAndPassword(dto *ClassicAuthDto, query QueryDSL) (*UserEntity, *PassportEntity, string, *IError) {
	var passport *PassportEntity
	var user *UserEntity
	var token string

	return user, passport, token, nil
}
