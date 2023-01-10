package emailsender

type IEmailSender interface {
	Send(subject, body, to, from string) error
}
