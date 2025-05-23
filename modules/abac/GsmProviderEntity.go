package abac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/torabian/fireback/modules/fireback"
	"github.com/urfave/cli"

	medianasms "github.com/medianasms/go-rest-sdk"
)

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
	Action: func(c *cli.Context) error {
		message := c.String("message")
		result, err := GsmSendSMS(c.String("id"), message, []string{c.String("to")})
		fireback.HandleActionInCli(c, result, err, map[string]map[string]string{})

		return nil
	},
}

func init() {
	GsmProviderCliCommands = append(GsmProviderCliCommands, GsmProviderTestCmd)
}

func GsmProviderActionCreate(
	dto *GsmProviderEntity, query fireback.QueryDSL,
) (*GsmProviderEntity, *fireback.IError) {
	return GsmProviderActionCreateFn(dto, query)
}

func GsmProviderActionUpdate(
	query fireback.QueryDSL,
	fields *GsmProviderEntity,
) (*GsmProviderEntity, *fireback.IError) {
	return GsmProviderActionUpdateFn(query, fields)
}

/**
*	Returns an specific template for an occiasion
*   for example, getting sms template for otp in europe area
**/

func GsmSendSMSUsingNotificationConfig(message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *fireback.IError) {

	config, err := NotificationConfigActionGetOneByWorkspace(fireback.QueryDSL{WorkspaceId: ROOT_VAR})
	if err != nil {
		// If there are no configuration, skip returning error, we use some terminal stuff for development.
		if err.HttpCode != 404 {
			return nil, err
		}
	}

	if config == nil || config.GeneralGsmProvider == nil {
		log.Default().Println("There is no gsm configuration unfortunately. We are printing the sms to the terminal for the sake of development.")
		log.Default().Println(message, recp)

		terminalQueue := "print-to-terminal"
		return &GsmSendSmsWithProviderActionResDto{QueueId: terminalQueue}, nil
	}

	return config.GeneralGsmProvider.SendSms(message, recp)
}

func GsmSendSMS(providerId string, message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *fireback.IError) {

	if provider, err := GsmProviderActions.GetOne(fireback.QueryDSL{UniqueId: providerId}); err != nil {
		return nil, err
	} else {
		return provider.SendSms(message, recp)
	}
}

func (x *GsmProviderEntity) SendSms(message string, recp []string) (*GsmSendSmsWithProviderActionResDto, *fireback.IError) {

	if x.Type == GsmProviderType.Url {
		if j, err := GsmSendSMSByHttpCall(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: j}, nil
		}
	} else if x.Type == GsmProviderType.Terminal {
		if j, err := GsmSendSMSByTerminal(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: j}, nil
		}
	} else if x.Type == GsmProviderType.Mediana {
		if j, err := GsmSendSMSByMediana(x, message, recp); err != nil {
			return nil, err
		} else {
			return &GsmSendSmsWithProviderActionResDto{QueueId: j}, nil
		}
	}

	fmt.Println(x.Json())
	return nil, fireback.Create401Error(&AbacMessages.SmsNotSent, []string{})
}

func GsmSendSMSByHttpCall(provider *GsmProviderEntity, message string, recp []string) (string, *fireback.IError) {
	fmt.Println("Sending sms using http call", provider.UniqueId)

	if provider.InvokeUrl == "" {
		return "", fireback.Create401Error(&AbacMessages.InvokeUrlMissing, []string{})
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
