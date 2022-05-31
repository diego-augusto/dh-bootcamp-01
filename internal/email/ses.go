package email

import "log"

type emailSES struct {
}

func NewSES() ServiceEmail {
	return &emailSES{}
}

func (emailSES) SendEmail(body string) {
	log.Printf("enviado pela aws: %s", body)
}
