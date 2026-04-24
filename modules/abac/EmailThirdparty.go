package abac

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/smtp"
	"time"

	"github.com/474420502/gcurl"
	"github.com/torabian/fireback/modules/fireback"
)

func sendViaSendGrid(msg EmailMessageContent, apiKey string) error {
	url := "https://api.sendgrid.com/v3/mail/send"

	payload := map[string]interface{}{
		"personalizations": []map[string]interface{}{
			{
				"to": []map[string]string{
					{"email": msg.ToEmail, "name": msg.ToName},
				},
			},
		},
		"from": map[string]string{
			"email": msg.FromEmail,
			"name":  msg.FromName,
		},
		"subject": msg.Subject,
		"content": []map[string]string{
			{"type": "text/plain", "value": msg.Content},
		},
	}

	return doJSONRequest("POST", url, payload, map[string]string{
		"Authorization": "Bearer " + apiKey,
	})
}

// =====================
// Mailgun (HTTP)
// =====================

func sendViaMailgun(msg EmailMessageContent, apiKey, domain string) error {
	url := fmt.Sprintf("https://api.mailgun.net/v3/%s/messages", domain)

	data := fmt.Sprintf(
		"from=%s <%s>&to=%s <%s>&subject=%s&text=%s",
		msg.FromName, msg.FromEmail,
		msg.ToName, msg.ToEmail,
		msg.Subject, msg.Content,
	)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer([]byte(data)))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	auth := "Basic " + base64.StdEncoding.EncodeToString([]byte("api:"+apiKey))
	req.Header.Set("Authorization", auth)

	return doRequest(req)
}

// =====================
// Postmark (HTTP)
// =====================

func sendViaPostmark(msg EmailMessageContent, token string) error {
	url := "https://api.postmarkapp.com/email"

	payload := map[string]interface{}{
		"From":     fmt.Sprintf("%s <%s>", msg.FromName, msg.FromEmail),
		"To":       fmt.Sprintf("%s <%s>", msg.ToName, msg.ToEmail),
		"Subject":  msg.Subject,
		"TextBody": msg.Content,
	}

	return doJSONRequest("POST", url, payload, map[string]string{
		"X-Postmark-Server-Token": token,
	})
}

// =====================
// Resend (HTTP)
// =====================

func sendViaResend(msg EmailMessageContent, apiKey string) error {
	url := "https://api.resend.com/emails"

	payload := map[string]interface{}{
		"from":    fmt.Sprintf("%s <%s>", msg.FromName, msg.FromEmail),
		"to":      []string{msg.ToEmail},
		"subject": msg.Subject,
		"text":    msg.Content,
	}

	return doJSONRequest("POST", url, payload, map[string]string{
		"Authorization": "Bearer " + apiKey,
	})
}

// =====================
// SMTP (Fallback)
// =====================

func sendViaSMTP(msg EmailMessageContent, host, port, user, pass string) error {
	auth := smtp.PlainAuth("", user, pass, host)

	body := []byte(fmt.Sprintf(
		"Subject: %s\r\nFrom: %s <%s>\r\nTo: %s <%s>\r\n\r\n%s",
		msg.Subject,
		msg.FromName, msg.FromEmail,
		msg.ToName, msg.ToEmail,
		msg.Content,
	))

	return smtp.SendMail(
		host+":"+port,
		auth,
		msg.FromEmail,
		[]string{msg.ToEmail},
		body,
	)
}

// func SendMail(message EmailMessageContent, provider *EmailProviderEntity) error {

// 	if provider == nil {
// 		return errors.New("GENERAL_EMAIL_PROVIDER_IS_NEEDED")
// 	}

// 	if provider.Type == EmailProviderType.Sendgrid {
// 		res, err := SendMailViaSendGrid(message, provider.ApiKey)

// 		if res != nil {
// 			fmt.Println(res.Body)
// 		}

// 		return err

// 	} else if provider.Type == EmailProviderType.Terminal {

// 		log.Default().Println(message.Json())

// 		return nil

// 	} else {
// 		return errors.New("EMAIL_PROVIDER_IS_NOT_AVAILABLE")
// 	}

// }

func SendWithCurl(msg EmailMessageContent, curlScript string) error {
	curl, err := gcurl.Parse(curlScript)
	if err != nil {
		return err
	}

	resp, err := curl.Request().Execute()
	if err != nil {
		return err
	}

	fmt.Printf("Status: %d\n", resp.GetStatusCode())
	fmt.Printf("Response: %s\n", resp.ContentString())

	return nil
}

type curlInfo struct {
	Curl string `json:"curl"`
}

func SendMail(msg EmailMessageContent, p *EmailProviderEntity) error {
	if p == nil {
		return errors.New("provider required")
	}

	switch p.Type {

	case "curl":
		m, _ := fireback.JSONTo[curlInfo](p.Config)

		return SendWithCurl(msg, m.Curl)

	// case "sendgrid":
	// 	return sendViaSendGrid(msg, p.Config["apiKey"])

	// case "mailgun":
	// 	return sendViaMailgun(msg, p.Config["apiKey"], p.Config["domain"])

	// case "postmark":
	// 	return sendViaPostmark(msg, p.Config["apiKey"])

	// case "resend":
	// 	return sendViaResend(msg, p.Config["apiKey"])

	// case "smtp":
	// 	return sendViaSMTP(
	// 		msg,
	// 		p.Config["host"],
	// 		p.Config["port"],
	// 		p.Config["user"],
	// 		p.Config["pass"],
	// 	)

	case "terminal":
		fmt.Println("EMAIL:", msg)
		return nil

	default:
		return errors.New("unsupported provider")
	}
}

// =====================
// Helpers
// =====================

func doJSONRequest(method, url string, payload interface{}, headers map[string]string) error {
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(method, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return doRequest(req)
}

func doRequest(req *http.Request) error {
	client := &http.Client{Timeout: 10 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode >= 300 {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("email error: %s", string(b))
	}

	return nil
}
