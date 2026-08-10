package messaging

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/torabian/emi/emigo"
	"github.com/torabian/fireback/modules/fireback"
	"github.com/torabian/fireback/modules/fireback/application"
	"github.com/urfave/cli/v3"

	medianasms "github.com/medianasms/go-rest-sdk"
)

var gsmProviderPerms = NewCrudPermissionSet("root.manage", "gsm-provider", "gsm provider")
var PERM_ROOT_GSM_PROVIDER = gsmProviderPerms.Wildcard
var PERM_ROOT_GSM_PROVIDER_QUERY = gsmProviderPerms.Query
var PERM_ROOT_GSM_PROVIDER_CREATE = gsmProviderPerms.Create
var PERM_ROOT_GSM_PROVIDER_UPDATE = gsmProviderPerms.Update
var PERM_ROOT_GSM_PROVIDER_DELETE = gsmProviderPerms.Delete
var ALL_GSM_PROVIDER_PERMISSIONS = gsmProviderPerms.All

var GsmProviderType = &xGsmProviderType{Url: "url", Terminal: "terminal", Mediana: "mediana"}

type xGsmProviderType struct {
	Url      string
	Terminal string
	Mediana  string
}

var GsmProviderActions = NewEntityActionsBundle[GsmProviderEntity]()

func GsmProviderBrowseAction(c GsmProviderBrowseActionRequest) (*GsmProviderBrowseActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_QUERY}})
	if err != nil {
		return nil, err
	}
	items, qrm, err2 := GsmProviderActions.Query(*query)
	if err2 != nil {
		return nil, err2
	}
	return &GsmProviderBrowseActionResponse{Payload: fireback.GResponseQuery(items, qrm, query)}, nil
}

func GsmProviderGetAction(c GsmProviderGetActionRequest) (*GsmProviderGetActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_QUERY}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	item, err2 := GsmProviderActions.GetOne(*query)
	if err2 != nil {
		return nil, err2
	}
	return &GsmProviderGetActionResponse{Payload: fireback.GResponseSingleItem(item)}, nil
}

func GsmProviderCreateAction(c GsmProviderCreateActionRequest) (*GsmProviderCreateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_CREATE}})
	if err != nil {
		return nil, err
	}
	if err2 := fireback.CommonStructValidatorPointer(&c.Body, false); err2 != nil {
		return nil, err2
	}
	entity := &GsmProviderEntity{
		ApiKey:           c.Body.ApiKey,
		MainSenderNumber: c.Body.MainSenderNumber,
		Type:             c.Body.Type,
		InvokeUrl:        c.Body.InvokeUrl,
		InvokeBody:       c.Body.InvokeBody,
		UserId:           emigo.NullableOf(query.UserId),
		WorkspaceId:      emigo.NullableOf(query.WorkspaceId),
	}
	created, err2 := GsmProviderActions.Create(entity, *query)
	if err2 != nil {
		return nil, err2
	}
	return &GsmProviderCreateActionResponse{Payload: fireback.GResponseSingleItem(created)}, nil
}

func GsmProviderUpdateAction(c GsmProviderUpdateActionRequest) (*GsmProviderUpdateActionResponse, error) {
	query, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_UPDATE}})
	if err != nil {
		return nil, err
	}
	query.UniqueId = c.Params.UniqueId
	fields := &GsmProviderEntity{UniqueId: c.Params.UniqueId}
	if v, ok := c.Body.ApiKey.Get(); ok {
		fields.ApiKey = *v
	}
	if v, ok := c.Body.MainSenderNumber.Get(); ok {
		fields.MainSenderNumber = *v
	}
	if v, ok := c.Body.Type.Get(); ok {
		fields.Type = *v
	}
	if v, ok := c.Body.InvokeUrl.Get(); ok {
		fields.InvokeUrl = *v
	}
	if v, ok := c.Body.InvokeBody.Get(); ok {
		fields.InvokeBody = *v
	}
	updated, err2 := GsmProviderActions.Update(*query, fields)
	if err2 != nil {
		return nil, err2
	}
	return &GsmProviderUpdateActionResponse{Payload: fireback.GResponseSingleItem(updated)}, nil
}

func GsmProviderAwareDeletePreviewAction(c GsmProviderAwareDeletePreviewActionRequest) (*GsmProviderAwareDeletePreviewActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_DELETE}}); err != nil {
		return nil, err
	}
	uniqueIds := GsmProviderAwareDeletePreviewActionQueryFromString(c.QueryParams.Encode()).UniqueIds
	preview, err2 := GsmProviderEntityActions.AwareDeletePreview(fireback.GetDbRef(), uniqueIds)
	if err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &GsmProviderAwareDeletePreviewActionResponse{Payload: fireback.GResponseSingleItem(preview)}, nil
}

func GsmProviderAwareDeleteAction(c GsmProviderAwareDeleteActionRequest) (*GsmProviderAwareDeleteActionResponse, error) {
	if _, err := fireback.ResolveActionContext(c, &fireback.SecurityModel{ActionRequires: []application.PermissionInfo{PERM_ROOT_GSM_PROVIDER_DELETE}}); err != nil {
		return nil, err
	}
	if err2 := GsmProviderEntityActions.AwareDelete(fireback.GetDbRef(), c.Body.UniqueIds); err2 != nil {
		return nil, fireback.GormErrorToIError(err2)
	}
	return &GsmProviderAwareDeleteActionResponse{Payload: fireback.GResponseSingleItem(struct{}{})}, nil
}

// --- Hand business logic recovered from the pre-migration GsmProviderEntity.go ---

var GsmProviderTestCmd cli.Command = cli.Command{

	Name:  "sms",
	Usage: "Sends the text message via gsm provider id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "id",
			Value:    "",
			Usage:    "Provider which you want to use for the message",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "message",
			Value:    "",
			Usage:    "Message content",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "to",
			Value:    "",
			Usage:    "Message recipient",
			Required: true,
		},
	},
	Action: func(ctx context.Context, c *cli.Command) error {
		// message := c.String("message")
		// result, err := GsmSendSMS(c.String("id"), message, []string{c.String("to")})
		// fireback.HandleActionInCli(c, result, err, map[string]map[string]string{})

		return nil
	},
}

func GsmSendSMS(providerId string, message string, recp []string) (*GsmSendSmsWithProviderActionRes, *fireback.IError) {

	if provider, err := GsmProviderActions.GetOne(fireback.QueryDSL{UniqueId: providerId}); err != nil {
		return nil, err
	} else {
		return provider.SendSms(message, recp)
	}
}

func (x *GsmProviderEntity) SendSms(message string, recp []string) (*GsmSendSmsWithProviderActionRes, *fireback.IError) {

	if x.Type == GsmProviderType.Url {
		if j, err := GsmSendSMSByHttpCall(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionRes{QueueId: j}, nil
		}
	} else if x.Type == GsmProviderType.Terminal {
		if j, err := GsmSendSMSByTerminal(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionRes{QueueId: j}, nil
		}
	} else if x.Type == GsmProviderType.Mediana {
		if j, err := GsmSendSMSByMediana(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionRes{QueueId: j}, nil
		}
	}

	fmt.Println(x.Json())
	return nil, fireback.Create401Error(&messagingMessages.SmsNotSent, []string{})
}

func GsmSendSMSByHttpCall(provider *GsmProviderEntity, message string, recp []string) (string, *fireback.IError) {
	fmt.Println("Sending sms using http call", provider.UniqueId)

	if provider.InvokeUrl == "" {
		return "", fireback.Create401Error(&messagingMessages.InvokeUrlMissing, []string{})
	}

	m, _ := json.MarshalIndent(recp, "", "  ")

	body := `{"apiKey":"{apiKey}","recipients":{recipients},"sender":"{sender}"}`

	if provider.InvokeBody != "" {
		body = provider.InvokeBody
	}

	if provider.ApiKey != "" {
		body = strings.ReplaceAll(body, "{apiKey}", provider.ApiKey)
	}
	body = strings.ReplaceAll(body, "{recipients}", string(m))
	if provider.MainSenderNumber != "" {
		body = strings.ReplaceAll(body, "{sender}", provider.MainSenderNumber)
	}
	fmt.Println("SMS Body:", body)

	req, err := http.NewRequest(http.MethodPost, provider.InvokeUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	return string(resBody), nil
}

func GsmSendSMSByTerminal(provider *GsmProviderEntity, message string, recp []string) (string, *fireback.IError) {

	fmt.Println("Sending sms using terminal by", provider.UniqueId)

	fmt.Println("Sms Message:", message)
	fmt.Println("Sms Recepients:", recp)

	return "", nil

}

func GsmSendSMSByMediana(provider *GsmProviderEntity, message string, recp []string) (string, *fireback.IError) {

	fmt.Println("Using mediana")
	sms := medianasms.New(provider.ApiKey)

	bulkID, err := sms.Send(provider.MainSenderNumber,
		recp, message)
	if err != nil {
		return "", fireback.GormErrorToIError(err)
	}

	return fmt.Sprintf("%v", bulkID), nil

}
