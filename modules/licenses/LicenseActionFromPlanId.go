package licenses

import (
	"fmt"
	"time"

	"github.com/torabian/fireback/modules/workspaces"
)

func LicenseActionFromPlanId(
	dto *LicenseFromPlanIdDto,
	query workspaces.QueryDSL,
) (*LicenseEntity, *workspaces.IError) {

	plan, err := ProductPlanActionGetOne(query)

	if err != nil {
		return nil, err
	}

	permissions := []LicenseContentPermission{}
	permissionsCloned := []*LicensePermissions{}

	if plan.Permissions != nil && len(plan.Permissions) > 0 {
		for _, item := range plan.Permissions {
			fmt.Println("Per", item)
			permissions = append(permissions, LicenseContentPermission{
				CapabilityId: *item.CapabilityId,
			})
		}

		for _, item := range plan.Permissions {
			permissionsCloned = append(permissionsCloned, &LicensePermissions{
				CapabilityId: item.CapabilityId,
				Capability:   item.Capability,
			})
		}
	}

	{
		doc := LicenseContent{
			Email:             *dto.Email,
			ValidityEndDate:   time.Now().Add(time.Hour * 24 * time.Duration(*plan.Duration)),
			WorkspaceId:       query.WorkspaceId,
			MachineId:         *dto.MachineId,
			Owner:             *dto.Owner,
			UserId:            query.UserId,
			ValidityStartDate: time.Now(),
			Permissions:       permissions,
		}

		fmt.Println(3, plan)
		fmt.Println(4, plan.Product)

		if plan.Product == nil || plan.Product.PrivateKey == nil || *plan.Product.PrivateKey == "" {
			return nil, &workspaces.IError{
				Message: LicensesMessageCode.PrivateKeyIsMissing,
			}
		}

		license, err := GenerateLicense(doc, *plan.Product.PrivateKey)

		if err != nil {
			return nil, workspaces.GormErrorToIError(err)
		}

		title := *plan.Product.Name + " - License from " +
			doc.ValidityStartDate.Format("2006-02-01") +
			" until " + doc.ValidityEndDate.Format("2006-02-01")

		license2, err := LicenseActionCreate(&LicenseEntity{
			ValidityStartDate: doc.ValidityStartDate,
			ValidityEndDate:   doc.ValidityEndDate,
			Name:              &title,
			SignedLicense:     &license,
			Permissions:       permissionsCloned,
		}, query)

		return license2, nil

	}
}
