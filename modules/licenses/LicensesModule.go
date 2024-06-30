package licenses

import (
	"embed"
	"fmt"
	"log"

	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

//go:embed *Module3.yml
var Module2Definitions embed.FS

func LicensesModuleSetup() *workspaces.ModuleProvider {
	module := &workspaces.ModuleProvider{
		Name:        "licenses",
		Definitions: &Module2Definitions,
	}

	module.ProvidePermissionHandler(ALL_LICENSE_PERMISSIONS)

	module.ProvideMockImportHandler(func() {
		// Import some fake products and plans
		LicensableProductImportMocks()
		ProductPlanImportMocks()

		// Generate keys based on all plans, these are not activated
		f := workspaces.QueryDSL{Deep: true, ItemsPerPage: 100, StartIndex: 0, WorkspaceId: "system"}
		items, count, _ := ProductPlanActionQuery(f)
		fmt.Println("Product plans count:", count, items)
		for _, productPlan := range items {
			LicenseActionSeederActivationKey(f, "2023-01-01", 10, 25, productPlan.UniqueId)

			Email := "demo@user.com"
			Owner := "Ali Torabi"
			MachineId := "XXXA-ADWW-W9289"

			// Also issue a license from this plan id
			if license, err := LicenseActionFromPlanId(&LicenseFromPlanIdDto{
				Email:     &Email,
				Owner:     &Owner,
				MachineId: &MachineId,
			}, workspaces.QueryDSL{UserId: "system", WorkspaceId: "system", UniqueId: productPlan.UniqueId}); err == nil {
				fmt.Println(license.SignedLicense)
			} else {
				log.Fatalln(err)
			}
		}

	})

	module.ProvideMockWriterHandler(func(languages []string) {
		LicensableProductWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
		ProductPlanWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
		ActivationKeyWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
		LicenseWriteQueryMock(workspaces.MockQueryContext{Languages: languages})
	})

	module.Actions = [][]workspaces.Module2Action{
		GetActivationKeyModule2Actions(),
		GetProductPlanModule2Actions(),
		GetLicensableProductModule2Actions(),
		GetLicenseModule2Actions(),
	}

	module.ProvideEntityHandlers(func(dbref *gorm.DB) error {

		return dbref.AutoMigrate(
			&LicenseEntity{},
			&LicensePermissions{},
			&LicensableProductEntity{},
			&ActivationKeyEntity{},
			&ProductPlanEntity{},
			&ProductPlanEntityPolyglot{},
			&ProductPlanPermissions{},
			&LicensableProductEntityPolyglot{},
		)
	})

	module.ProvideCliHandlers([]cli.Command{
		LicenseCliFn(),
	})

	return module
}
