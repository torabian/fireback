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

	// err := GetDbRef().Transaction(func(tx *gorm.DB) error {

	// 	userUniqueId := UUID()
	// 	user = &UserEntity{

	// 		UniqueId: userUniqueId,
	// 		Person: &PersonEntity{
	// 			FirstName: dto.FirstName,
	// 			LastName:  dto.LastName,
	// 		},
	// 	}
	// 	err := tx.Create(&user).Error

	// 	if err != nil {
	// 		return err
	// 	}

	// 	passport, err = CreateEmailPassportForUser(userUniqueId, *dto.Email, *dto.Password, tx)

	// 	if err != nil {
	// 		if strings.Contains(err.Error(), "UNIQUE constraint") {
	// 			return errors.New(PassportMessageCode.PassportNotAvailable)
	// 		}
	// 		return err
	// 	}

	// 	// var roleId string
	// 	// var workspaceId string

	// 	if dto.PublicJoinKeyId == nil {
	// 		root := "root"
	// 		workspace := &WorkspaceEntity{
	// 			Name:     dto.FirstName,
	// 			UniqueId: UUID(),
	// 			ParentId: &root,
	// 		}

	// 		workspaceErr := tx.Create(&workspace).Error

	// 		if workspaceErr != nil {
	// 			return workspaceErr
	// 		}

	// 		roleName := "Administrator"
	// 		capabilities := []*CapabilityEntity{
	// 			{UniqueId: "root/*"},
	// 		}

	// 		// Get the workspace type id, and assign that role to the user instead of creating one for him
	// 		if dto.WorkspaceTypeId != nil {
	// 			q2 := query
	// 			q2.Tx = tx
	// 			q2.Query = "slug = " + *dto.WorkspaceTypeId
	// 			q2.Deep = true
	// 			items, _, errWt := WorkspaceTypeActionQuery(q2)
	// 			var wt *WorkspaceTypeEntity = nil

	// 			if errWt != nil {
	// 				return errWt
	// 			}

	// 			if len(items) == 0 || items[0] == nil {
	// 				return errors.New("this method of signup is no longer available")
	// 			}

	// 			wt = items[0]

	// 			if wt.Role == nil {
	// 				return errors.New("this method of signup is no longer available")
	// 			}

	// 			roleId = *wt.RoleId

	// 		} else {

	// 			roleD := &RoleEntity{
	// 				UniqueId:     UUID(),
	// 				Name:         &roleName,
	// 				WorkspaceId:  &workspace.UniqueId,
	// 				Capabilities: capabilities,
	// 			}

	// 			rolex, err2 := RoleActionCreate(roleD, QueryDSL{Tx: tx})
	// 			if err2 != nil {
	// 				return err
	// 			}

	// 			roleId = rolex.UniqueId
	// 		}

	// 		workspaceId = workspace.UniqueId
	// 	} else {

	// 		joinEntity := &PublicJoinKeyEntity{}
	// 		err := tx.Where("unique_id = ?", dto.PublicJoinKeyId).First(joinEntity).Error

	// 		if err != nil {
	// 			return err
	// 		}
	// 		roleId = *joinEntity.RoleId
	// 		workspaceId = *joinEntity.WorkspaceId
	// 	}

	// 	// linker := UserRoleWorkspaceEntity{
	// 	// 	User:        user,
	// 	// 	RoleId:      &roleId,
	// 	// 	UniqueId:    UUID(),
	// 	// 	WorkspaceId: &workspaceId,
	// 	// }

	// 	return tx.Create(&linker).Error

	// })

	// if err != nil {
	// 	return nil, nil, "", GormErrorToIError(err)
	// }

	// // Getting the token and user information should happend outside of transaction
	// user, token, err = SigninUserWithEmailAndPassword(*dto.Email, *dto.Password)

	// if err != nil {
	// 	return nil, nil, "", GormErrorToIError(err)
	// }

	return user, passport, token, nil
}
