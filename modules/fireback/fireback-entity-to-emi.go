package fireback

import (
	"fmt"

	"github.com/torabian/emi/lib/core"
)

// This file converts entities from fireback into the emi dto and actions
// Fireback won't generate rpc code anymore, therefor this will be an option to give some
// control over actions being generated.

type EntityEmiData struct {
	Actions []*FirebackEmiAction `json:"actions"`
	Dtos    []*core.EmiDto       `json:"dtos"`
}

func ConvertEntityToEmi(entity Module3Entity) EntityEmiData {
	res := EntityEmiData{}

	refDto := core.EmiDto{}
	var walk func(fields []*Module3Field) []*core.EmiField
	walk = func(fields []*Module3Field) []*core.EmiField {

		emiFields := []*core.EmiField{}
		for _, field := range fields {
			newField := &core.EmiField{
				Name: field.Name,
				Type: core.FieldType(field.Type),
			}

			if len(field.Fields) > 0 {
				newField.Fields = walk(field.Fields)
			}

			emiFields = append(emiFields, newField)
		}

		return emiFields
	}

	refDto.Name = entity.Name
	refDto.Fields = walk(entity.Fields)
	// res.Dtos = append(res.Dtos, &refDto)

	{
		action := FirebackEmiAction{
			EmiAction: core.EmiAction{

				Name:   fmt.Sprintf("create%v", entity.Upper()),
				Url:    fmt.Sprintf("/%v", entity.DashedName()),
				Method: "post",
				In: &core.EmiActionBody{
					Fields: refDto.Fields,
				},
				Out: &core.EmiActionBody{
					Fields: refDto.Fields,
				},
			},
		}
		res.Actions = append(res.Actions, &action)
	}

	{
		action := FirebackEmiAction{
			EmiAction: core.EmiAction{

				Name:   fmt.Sprintf("update%v", entity.Upper()),
				Url:    fmt.Sprintf("/%v", entity.DashedName()),
				Method: "patch",
			},
		}
		res.Actions = append(res.Actions, &action)
	}
	{
		action := FirebackEmiAction{
			EmiAction: core.EmiAction{

				Name:   fmt.Sprintf("query%v", ToUpper(entity.PluralName())),
				Url:    fmt.Sprintf("/%v", entity.DashedName()),
				Method: "get",
			},
		}
		res.Actions = append(res.Actions, &action)
	}
	{
		action := FirebackEmiAction{
			EmiAction: core.EmiAction{

				Name:   fmt.Sprintf("delete%v", entity.Upper()),
				Url:    fmt.Sprintf("/%v", entity.DashedName()),
				Method: "delete",
			},
		}
		res.Actions = append(res.Actions, &action)
	}
	{
		action := FirebackEmiAction{
			EmiAction: core.EmiAction{

				Name:   fmt.Sprintf("delete%v", ToUpper(entity.PluralName())),
				Url:    fmt.Sprintf("/%v", entity.DashedName()),
				Method: "delete",
			},
		}
		res.Actions = append(res.Actions, &action)
	}
	return res
}
