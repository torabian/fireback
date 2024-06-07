package workspaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	medianasms "github.com/medianasms/go-rest-sdk"
)

func GsmProviderActionCreate(
	dto *GsmProviderEntity, query QueryDSL,
) (*GsmProviderEntity, *IError) {
	return GsmProviderActionCreateFn(dto, query)
}

func GsmProviderActionUpdate(
	query QueryDSL,
	fields *GsmProviderEntity,
) (*GsmProviderEntity, *IError) {
	return GsmProviderActionUpdateFn(query, fields)
}

/**
*	Returns an specific template for an occiasion
*   for example, getting sms template for otp in europe area
**/

func GsmSendSMSUsingNotificationConfig(message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(QueryDSL{WorkspaceId: ROOT_VAR})

	if err != nil {
		return nil, err
	}

	provider := config.GeneralGsmProvider

	if provider == nil {
		return nil, Create401Error(&WorkspacesMessages.GsmConfigurationIsNotAvailable, []string{})
	} else {
		return provider.SendSms(message, recp)
	}
}

func GsmSendSMS(providerId string, message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *IError) {

	if provider, err := GsmProviderActionGetOne(QueryDSL{UniqueId: providerId}); err != nil {
		return nil, err
	} else {
		return provider.SendSms(message, recp)
	}
}

func (x *GsmProviderEntity) SendSms(message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *IError) {

	if *x.Type == GsmProviderType.Url {
		if j, err := GsmSendSMSByHttpCall(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: &j}, nil
		}
	} else if *x.Type == GsmProviderType.Terminal {
		if j, err := GsmSendSMSByTerminal(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: &j}, nil
		}
	} else if *x.Type == GsmProviderType.Mediana {
		if j, err := GsmSendSMSByMediana(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: &j}, nil
		}
	}

	fmt.Println(x.Json())
	return nil, Create401Error(&WorkspacesMessages.SmsNotSent, []string{})
}

func GsmSendSMSByHttpCall(provider *GsmProviderEntity, message string, recp []string) (string, *IError) {
	fmt.Println("Sending sms using http call", provider.UniqueId)

	if provider.InvokeUrl == nil {
		return "", Create401Error(&WorkspacesMessages.InvokeUrlMissing, []string{})
	}

	m, _ := json.MarshalIndent(recp, "", "  ")

	body := `{"apiKey":"{apiKey}","recipients":{recipients},"sender":"{sender}"}`

	if provider.InvokeBody != nil {
		body = *provider.InvokeBody
	}

	if provider.ApiKey != nil {
		body = strings.ReplaceAll(body, "{apiKey}", *provider.ApiKey)
	}
	body = strings.ReplaceAll(body, "{recipients}", string(m))
	if provider.MainSenderNumber != nil {
		body = strings.ReplaceAll(body, "{sender}", *provider.MainSenderNumber)
	}
	fmt.Println("SMS Body:", body)

	req, err := http.NewRequest(http.MethodPost, *provider.InvokeUrl, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return "", GormErrorToIError(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", GormErrorToIError(err)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", GormErrorToIError(err)
	}

	return string(resBody), nil
}

func GsmSendSMSByTerminal(provider *GsmProviderEntity, message string, recp []string) (string, *IError) {

	fmt.Println("Sending sms using terminal by", provider.UniqueId)

	fmt.Println("Sms Message:", message)
	fmt.Println("Sms Recepients:", recp)

	return "", nil

}

func GsmSendSMSByMediana(provider *GsmProviderEntity, message string, recp []string) (string, *IError) {

	fmt.Println("Using mediana")
	sms := medianasms.New(*provider.ApiKey)

	bulkID, err := sms.Send(*provider.MainSenderNumber,
		recp, message)
	if err != nil {
		return "", GormErrorToIError(err)
	}

	return fmt.Sprintf("%v", bulkID), nil

}
