package smtp

import (
	"bytes"
	"fmt"
	"net/smtp"
)

type Config struct {
	From     string
	Host     string
	Port     string
	Password string
	UserName string
}

type EmailService interface {
	GetMailData(subject string, body string) []byte
	SendMail(to []string, data []byte) error
}

type SmtpService struct {
	cfg *Config
}

func NewSmtpService(cfg *Config) *SmtpService {
	return &SmtpService{cfg: cfg}
}

func (s *SmtpService) GetMailData(subject string, body string) []byte {
	var content bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	content.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", subject, mimeHeaders)))
	content.Write([]byte(body))
	return content.Bytes()
}

func (s *SmtpService) SendMail(to []string, data []byte) error {
	auth := smtp.PlainAuth("", s.cfg.UserName, s.cfg.Password, s.cfg.Host)
	if err := smtp.SendMail(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port), auth, s.cfg.From, to, data); err != nil {
		return err
	}

	return nil
}
