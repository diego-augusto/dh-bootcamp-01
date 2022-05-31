package email

import "log"

type sendgrid struct {
}

func NewSendgrid() ServiceEmail {
	return &sendgrid{}
}

func (sendgrid) SendEmail(body string) {
	log.Printf("enviado pelo sendgrid: %s", body)
}
