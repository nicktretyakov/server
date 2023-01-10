package emailsender

import (
	"gopkg.in/gomail.v2"
)

func (s *service) Send(subject, body, to, from string) error {
	m := gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
	m.SetAddressHeader("From", from, "")
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return s.dialer.DialAndSend(m)
}
