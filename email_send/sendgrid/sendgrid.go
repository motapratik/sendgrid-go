package sendgrid

import (
	"fmt"
	"log"

	"github.com/motapratik/sendgrid-go/email_send"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailDetail struct {
	toEmailAddress   string
	toName           string
	fromEmailAddress string
	fromName         string
	sendgridAPIKey   string
	subjectText      string
}

// New returns a new IEmail
// NewEmail create object with all details for sending email
func NewEmail(toEmail, toEmailUserName, fromEmail, fromEmailUserName, apiKey, subjectLine string) email_send.IEmail {
	return &EmailDetail{
		toEmailAddress:   toEmail,
		toName:           toEmailUserName,
		fromEmailAddress: fromEmail,
		fromName:         fromEmailUserName,
		sendgridAPIKey:   apiKey,
		subjectText:      subjectLine,
	}
}

func (eptr *EmailDetail) SendEmail() (errEmail error) {

	from := mail.NewEmail(eptr.fromName, eptr.fromEmailAddress)
	subject := eptr.subjectText
	to := mail.NewEmail(eptr.toName, eptr.toEmailAddress)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(eptr.sendgridAPIKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		return nil
	}

}
