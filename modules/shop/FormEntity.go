package shop

import (
	"github.com/torabian/fireback/modules/workspaces"
	"github.com/urfave/cli"
)

func FormFieldActionBatchCreateFn(dtos []*FormFields, query workspaces.QueryDSL) ([]*FormFields, *workspaces.IError) {
	if dtos != nil && len(dtos) > 0 {
		items := []*FormFields{}
		for _, item := range dtos {
			s, err := FormFieldsActionCreate(item, query)
			if err != nil {
				return nil, err
			}
			items = append(items, s)
		}
		return items, nil
	}
	return dtos, nil
}

var Analise cli.Command = cli.Command{
	Name:  "analise",
	Usage: "Analise the structure of a eav model",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:     "file",
			Usage:    "File address",
			Required: true,
		},
	},
	Action: func(c *cli.Context) error {

		form := EavModel{}
		workspaces.ReadYamlFile[EavModel](c.String("file"), &form)

		query := workspaces.CommonCliQueryDSLBuilderAuthorize(c, nil)
		form.StoreInDatabase(query)

		return nil
	},
}

func init() {
	FormCliCommands = append(FormCliCommands, Analise)
}

func FormActionCreate(
	dto *FormEntity, query workspaces.QueryDSL,
) (*FormEntity, *workspaces.IError) {

	return FormActionCreateFn(dto, query)
}

func FormActionUpdate(
	query workspaces.QueryDSL,
	fields *FormEntity,
) (*FormEntity, *workspaces.IError) {
	return FormActionUpdateFn(query, fields)
}
