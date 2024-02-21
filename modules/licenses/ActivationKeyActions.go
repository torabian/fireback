package licenses

import (
	"fmt"

	"github.com/schollz/progressbar/v3"
	"pixelplux.com/fireback/modules/workspaces"
)

func ActivationKeyActionCreate(
	dto *ActivationKeyEntity, query workspaces.QueryDSL,
) (*ActivationKeyEntity, *workspaces.IError) {
	return ActivationKeyActionCreateFn(dto, query)
}

func ActivationKeyActionUpdate(
	query workspaces.QueryDSL,
	fields *ActivationKeyEntity,
) (*ActivationKeyEntity, *workspaces.IError) {
	return ActivationKeyActionUpdateFn(query, fields)
}

/**
*	Generates activation key, which could be used while activating the software
* 	Useful when you want to distribute key on paper in the shops for example
**/
func LicenseActionSeederActivationKey(query workspaces.QueryDSL, series string, count int, length int, planId string) {

	successInsert := 0
	failureInsert := 0

	bar := progressbar.Default(int64(count))

	if series == "" {
		series = workspaces.UUID()
	}

	for i := 1; i <= count; i++ {
		used := int64(0)
		entity := &ActivationKeyEntity{
			UniqueId: workspaces.GenerateRandomKey(length),
			Series:   &series,
			Used:     &used,
			PlanId:   &planId,
		}

		_, err := ActivationKeyActionCreate(entity, query)
		if err == nil {
			successInsert++
		} else {
			failureInsert++
		}

		bar.Add(1)
		// time.Sleep(1 * time.Millisecond)
	}
	fmt.Println("Done generateion", series, count, length, planId)

	fmt.Println("Success", successInsert, "Failure", failureInsert)
}
