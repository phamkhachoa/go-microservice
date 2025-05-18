package sendto

import (
	"bytes"
	"fmt"
	"go-ecommerce-backend-api/global"
	"go.uber.org/zap"
	"html/template"
	"net/smtp"
	"strings"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

var (
	SMTPHost = "smtp.gmail.com"
	SMTPPort = "587"
	SMTPUser = "phamkhachoabk@gmail.com"
	SMTPPass = "tdsidyoqcmqadrlg"
)

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendTextEmailOtp(to []string, from string, otp string) error {
	//contentEmail := Mail{
	//	From:    EmailAddress{Address: from, Name: "test"},
	//	To:      to,
	//	Subject: "OTP Verification",
	//	Body:    fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp),
	//}

	message := fmt.Sprintf("Your OTP is %s. Please enter it to verify your account.", otp)

	//messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, SMTPUser, to, []byte(message))

	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}

	return nil
}

func SendTemplateEmailOtp(to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {
	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}
	
	return send(to, from, htmlBody)
}

func getMailTemplate(nameTemplate string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTemplate).ParseFiles("templates-email/" + nameTemplate))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}
	return htmlTemplate.String(), nil
}

func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: from, Name: "test"},
		To:      to,
		Subject: "OTP Verification",
		Body:    htmlTemplate,
	}

	messageMail := BuildMessage(contentEmail)

	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPHost)

	err := smtp.SendMail(SMTPHost+":"+SMTPPort, auth, SMTPUser, to, []byte(messageMail))

	if err != nil {
		global.Logger.Error("Email send failed::", zap.Error(err))
		return err
	}

	return nil
}
