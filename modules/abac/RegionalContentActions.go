package abac

import (
	"bytes"
	"context"
	"log"
	"strings"
	"text/template"

	"github.com/torabian/emi/emigo"
	abacdefs "github.com/torabian/fireback/modules/abac/defs"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"
	"gorm.io/gorm"
)

// regionalContent's old security block was { writeOnRoot: true }: the old generated code
// only set AllowOnRoot: true on Create/Update/Delete, leaving Query/Get plain -
// preserved here exactly.
var regionalContentPerms = NewCrudPermissionSet("root.manage", "regional-content", "regional content")
var PERM_ROOT_REGIONAL_CONTENT = regionalContentPerms.Wildcard
var PERM_ROOT_REGIONAL_CONTENT_QUERY = regionalContentPerms.Query
var PERM_ROOT_REGIONAL_CONTENT_CREATE = regionalContentPerms.Create
var PERM_ROOT_REGIONAL_CONTENT_UPDATE = regionalContentPerms.Update
var PERM_ROOT_REGIONAL_CONTENT_DELETE = regionalContentPerms.Delete
var ALL_REGIONAL_CONTENT_PERMISSIONS = regionalContentPerms.All

var RegionalContentActions = NewEntityActionsBundle[abacdefs.RegionalContentEntity]()

func RegionalContentBrowseAction(c abacdefs.RegionalContentBrowseActionRequest) (*abacdefs.RegionalContentBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := RegionalContentActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.RegionalContentBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func RegionalContentGetAction(c abacdefs.RegionalContentGetActionRequest) (*abacdefs.RegionalContentGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := RegionalContentActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.RegionalContentGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func RegionalContentCreateAction(c abacdefs.RegionalContentCreateActionRequest) (*abacdefs.RegionalContentCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_CREATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	// content/region/languageId/keyGroup are validate:"required" (see
	// RegionalContentDto.go) but nothing was actually enforcing it - same fix as
	// EmailProviderCreateAction/TableViewSizingCreateAction/PassportCreateAction.
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	entity := &abacdefs.RegionalContentEntity{
		Content:    c.Body.Content,
		Region:     c.Body.Region,
		Title:      c.Body.Title,
		LanguageId: c.Body.LanguageId,
		KeyGroup:   c.Body.KeyGroup,
		// Without this, the row is invisible to RegionalContentBrowseAction's
		// workspace-scoped query (see GetSqlContext) - same workspace-stamping fix as
		// EmailConfirmationCreateAction/PassportCreateAction/AppMenuCreateAction.
		WorkspaceId: emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := RegionalContentActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.RegionalContentCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func RegionalContentUpdateAction(c abacdefs.RegionalContentUpdateActionRequest) (*abacdefs.RegionalContentUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_UPDATE}, AllowOnRoot: true})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &abacdefs.RegionalContentEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.Content.Get(); ok {
		fields.Content = *v
	}
	if v, ok := c.Body.Region.Get(); ok {
		fields.Region = *v
	}
	if v, ok := c.Body.Title.Get(); ok {
		fields.Title = *v
	}
	if v, ok := c.Body.LanguageId.Get(); ok {
		fields.LanguageId = *v
	}
	if v, ok := c.Body.KeyGroup.Get(); ok {
		fields.KeyGroup = *v
	}
	updated, err2 := RegionalContentActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &abacdefs.RegionalContentUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func RegionalContentAwareDeletePreviewAction(c abacdefs.RegionalContentAwareDeletePreviewActionRequest) (*abacdefs.RegionalContentAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	uniqueIds := abacdefs.RegionalContentAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := abacdefs.RegionalContentEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.RegionalContentAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func RegionalContentAwareDeleteAction(c abacdefs.RegionalContentAwareDeleteActionRequest) (*abacdefs.RegionalContentAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_REGIONAL_CONTENT_DELETE}, AllowOnRoot: true}); err != nil {
		return nil, err
	}
	if err2 := abacdefs.RegionalContentEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &abacdefs.RegionalContentAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration RegionalContentEntity.go ---

type RegionContentKey string

const (
	SMS_OTP   RegionContentKey = "SMS_OTP"
	EMAIL_OTP RegionContentKey = "EMAIL_OTP"
)

type RegionalContentRequest struct {
	Region           string
	LanguageId       string
	RegionContentKey RegionContentKey
}

func RegionContentKeys() []string {
	return []string{string(EMAIL_OTP), string(SMS_OTP)}
}

func CompileContent(x *abacdefs.RegionalContentEntity, data map[string]string) (string, error) {
	if x.Content == "" {
		return "", nil
	}

	// Create a template and parse the template string
	tmpl, err := template.New("regionalContent").Parse(x.Content)
	if err != nil {
		return "", err
	}

	// Create a buffer to capture the template output
	var tplOutput bytes.Buffer

	// Execute the template and capture the output into the buffer
	err = tmpl.Execute(&tplOutput, data)
	if err != nil {
		return "", err
	}

	// Convert the buffer content to a string
	return tplOutput.String(), nil

}

var DefaultOtpForEmailMessageTitle string = `
Code: {{ .Otp }}
`

var DefaultOtpForEmailMessage string = `
Use the following code for single time authorization

{{ .Otp }}
`

func QuickGetOtpMessage(q fireback.QueryDSL, field RegionContentKey) *abacdefs.RegionalContentEntity {
	if result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
		LanguageId:       q.Language,
		Region:           "any",
		RegionContentKey: field,
	}, q); err != nil || result == nil {

		log.Default().Println("For otp, the default content has been used. Make sure you update the regional content, you can customize it for different users, regions, and languages")

		return &abacdefs.RegionalContentEntity{
			Content:  DefaultOtpForEmailMessage,
			Title:    DefaultOtpForEmailMessageTitle,
			UniqueId: "~in-binary-default-content",
		}
	} else {
		return result
	}
}

func CompileTitle(x *abacdefs.RegionalContentEntity, data map[string]string) (string, error) {
	if x.Title == "" {
		return "", nil
	}

	// Create a template and parse the template string
	tmpl, err := template.New("regionalContent").Parse(x.Title)
	if err != nil {
		return "", err
	}

	// Create a buffer to capture the template output
	var tplOutput bytes.Buffer

	// Execute the template and capture the output into the buffer
	err = tmpl.Execute(&tplOutput, data)
	if err != nil {
		return "", err
	}

	// Convert the buffer content to a string
	return tplOutput.String(), nil

}

func ResolveRegionalContentTemplate(dto *RegionalContentRequest, q fireback.QueryDSL) (*abacdefs.RegionalContentEntity, *fireback.IError) {

	key := string(dto.RegionContentKey)
	var item abacdefs.RegionalContentEntity
	condition := &abacdefs.RegionalContentEntity{LanguageId: dto.LanguageId, Region: dto.Region, KeyGroup: key}

	if err := fireback.GetRef(q).
		Debug().
		Model(&abacdefs.RegionalContentEntity{}).
		Where(condition).
		First(&item).Error; err != nil {

		// If looking for a key in other than english and we do not have, let's get the english one instead
		// It's better to send templates in English than sending an error
		if condition.LanguageId != "en" && err == gorm.ErrRecordNotFound {
			condition.LanguageId = "en"
			condition.Region = "any"

			if err2 := fireback.GetRef(q).
				Debug().
				Model(&abacdefs.RegionalContentEntity{}).
				Where(condition).
				First(&item).Error; err2 != nil {
				return nil, fireback.GormErrorToIError(err2)
			}
		} else {
			return nil, fireback.GormErrorToIError(err)
		}
	}

	return &item, nil
}

var RegionalContentGetCmd cli.Command = cli.Command{

	Name:  "get",
	Usage: "Gets a template by region",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "region",
			Usage:    "Set the region or language code (examples: any, asia/*, europe/*, pl, fa)",
			Value:    "any",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "key",
			Usage:    "The key code for template (" + strings.Join(RegionContentKeys(), ", ") + ")",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "lang",
			Usage:    "The language code (fa, en, ...)",
			Required: true,
			Value:    "en",
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		// f := fireback.CommonCliQueryDSLBuilder(c)

		// result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
		// 	LanguageId:       c.String("lang"),
		// 	Region:           c.String("region"),
		// 	RegionContentKey: RegionContentKey(c.String(("key"))),
		// }, f)
		// fireback.HandleActionInCli(c, result, err, map[string]map[string]string{})

		return nil
	},
}
