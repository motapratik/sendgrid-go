package email_send

// IEmail is interface to send email
type IEmail interface {
	// Send Email to user
	SendEmail() (err error)
}
