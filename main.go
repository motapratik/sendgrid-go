package main

import (
	"fmt"

	emailsendgrid "github.com/motapratik/sendgrid-go/email_send/sendgrid"
)

func main() {
	//sendgrid return interface object
	email := emailsendgrid.NewEmail("to@gmail.com", "toName", "from@gmail.com", "formName", "API_KEY", "Subject line")
	// using interface object call send email function
	err := email.SendEmail()
	if err != nil {
		fmt.Println("Error in sending Email !!")
	}
}
