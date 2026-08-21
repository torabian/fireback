package resolve

import (
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
)

func GetUserFromToken(tokenString string) (*abacdefs.UserEntity, error) {

	var item abacdefs.TokenEntity

	if err := fireback.GetDbRef().Where(fireback.RealEscape("token = ?", tokenString)).First(&item).Error; err != nil {
		return &abacdefs.UserEntity{}, err
	}

	// Not workspace-scoped (see UserBrowseAction's own comment) - abacdefs.UserEntityActions.Get
	// is the entity's own generated, unscoped lookup.
	user, _ := abacdefs.UserEntityActions.Get(fireback.GetDbRef(), item.UserId.OrDefault(""))

	// HydrateUserPrimaryAddress(user)
	return user, nil
}
