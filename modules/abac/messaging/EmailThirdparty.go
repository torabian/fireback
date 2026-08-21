package messaging

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
	messagingdefs "github.com/torabian/fireback/modules/abac/messaging/defs"
	"github.com/torabian/fireback/modules/fireback/complexes"
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

type sendGridInfo struct {
	APIKey string `json:"apiKey"`
}

type mailGunInfo struct {
	ApiKey string `json:"apiKey"`
	Domain string `json:"domain"`
}

type postmarkInfo struct {
	ApiKey string `json:"apiKey"`
}

type resendInfo struct {
	ApiKey string `json:"apiKey"`
}

type smtpInfo struct {
	Host string `json:"host"`
	Port string `json:"port"`
	User string `json:"user"`
	Pass string `json:"pass"`
}

func SendMail(msg EmailMessageContent, p *messagingdefs.EmailProviderEntity) error {
	if p == nil {
		return errors.New("provider required")
	}

	switch p.Type {

	case "curl":
		m, _ := complexes.JSONTo[curlInfo](p.Config)
		return SendWithCurl(msg, m.Curl)

	case "sendgrid":
		m, _ := complexes.JSONTo[sendGridInfo](p.Config)
		return sendViaSendGrid(msg, m.APIKey)

	case "mailgun":
		m, _ := complexes.JSONTo[mailGunInfo](p.Config)
		return sendViaMailgun(msg, m.ApiKey, m.Domain)

	case "postmark":
		m, _ := complexes.JSONTo[postmarkInfo](p.Config)
		return sendViaPostmark(msg, m.ApiKey)

	case "resend":
		m, _ := complexes.JSONTo[resendInfo](p.Config)
		return sendViaResend(msg, m.ApiKey)

	case "smtp":
		m, _ := complexes.JSONTo[smtpInfo](p.Config)
		return sendViaSMTP(
			msg,
			m.Host,
			m.Port,
			m.User,
			m.Pass,
		)

	case "terminal":
		fmt.Println("EMAIL:", msg.Json())
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
