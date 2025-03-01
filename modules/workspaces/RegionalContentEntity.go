package workspaces

import (
	"bytes"
	"log"
	"strings"
	"text/template"

	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func RegionalContentActionCreate(
	dto *RegionalContentEntity, query QueryDSL,
) (*RegionalContentEntity, *IError) {
	return RegionalContentActionCreateFn(dto, query)
}
func RegionalContentActionUpdate(
	query QueryDSL,
	fields *RegionalContentEntity,
) (*RegionalContentEntity, *IError) {
	return RegionalContentActionUpdateFn(query, fields)
}

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

func (x *RegionalContentEntity) CompileContent(data map[string]string) (string, error) {
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

func QuickGetOtpMessage(q QueryDSL, field RegionContentKey) *RegionalContentEntity {
	if result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
		LanguageId:       q.Language,
		Region:           "any",
		RegionContentKey: field,
	}, q); err != nil || result == nil {

		log.Default().Println("For otp, the default content has been used. Make sure you update the regional content, you can customize it for different users, regions, and languages")

		return &RegionalContentEntity{
			Content:        DefaultOtpForEmailMessage,
			ContentExcerpt: &DefaultOtpForEmailMessage,
			Title:          DefaultOtpForEmailMessageTitle,
			UniqueId:       "~in-binary-default-content",
		}
	} else {
		return result
	}
}

func (x *RegionalContentEntity) CompileTitle(data map[string]string) (string, error) {
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

func ResolveRegionalContentTemplate(dto *RegionalContentRequest, q QueryDSL) (*RegionalContentEntity, *IError) {

	key := string(dto.RegionContentKey)
	var item RegionalContentEntity
	condition := &RegionalContentEntity{LanguageId: dto.LanguageId, Region: dto.Region, KeyGroup: key}

	if err := GetRef(q).
		Debug().
		Model(&RegionalContentEntity{}).
		Where(condition).
		First(&item).Error; err != nil {

		// If looking for a key in other than english and we do not have, let's get the english one instead
		// It's better to send templates in English than sending an error
		if condition.LanguageId != "en" && err == gorm.ErrRecordNotFound {
			condition.LanguageId = "en"
			condition.Region = "any"

			if err2 := GetRef(q).
				Debug().
				Model(&RegionalContentEntity{}).
				Where(condition).
				First(&item).Error; err2 != nil {
				return nil, GormErrorToIError(err2)
			}
		} else {
			return nil, GormErrorToIError(err)
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
	Action: func(c *cli.Context) error {
		f := CommonCliQueryDSLBuilder(c)

		result, err := ResolveRegionalContentTemplate(&RegionalContentRequest{
			LanguageId:       c.String("lang"),
			Region:           c.String("region"),
			RegionContentKey: RegionContentKey(c.String(("key"))),
		}, f)
		HandleActionInCli(c, result, err, map[string]map[string]string{})

		return nil
	},
}

func init() {
	RegionalContentCliCommands = append(RegionalContentCliCommands, RegionalContentGetCmd)
}
