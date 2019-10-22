package lib

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/YaroslavRozum/go-boilerplate/settings"
)

var templateByType = map[string]*template.Template{}

// EmailSender => (to []string)
type EmailSender interface {
	Send([]string, string, interface{})
}

type emailSender struct {
	host           string
	port           string
	senderEmail    string
	senderPassword string
	auth           smtp.Auth
	fullAddres     string
}

func (e *emailSender) address() string {
	return fmt.Sprintf("%s:%s", e.host, e.port)
}

func (e *emailSender) getTemplate(templateName string) (*template.Template, error) {
	if t, ok := templateByType[templateName]; ok {
		return t, nil
	}
	templatePath := fmt.Sprintf("./templates/%s.html", templateName)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return nil, err
	}
	templateByType[templateName] = t
	return t, nil
}

func (e *emailSender) Send(to []string, templateName string, data interface{}) {
	t, err := e.getTemplate(templateName)
	if err != nil {
		log.Printf("Parsing template error %s", err.Error())
		return
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		log.Printf("Executing template error %s", err.Error())
		return
	}
	body := buf.String()
	e.send(to, body)
}

func (e *emailSender) send(to []string, body string) {
	err := smtp.SendMail(
		e.fullAddres,
		e.auth,
		e.senderEmail,
		to,
		[]byte(body),
	)

	if err != nil {
		log.Printf("smtp error: %s", err.Error())
		return
	}
}

func NewEmailSender(settings settings.Settings) EmailSender {
	newEmailSender := emailSender{
		host:           "smtp.gmail.com",
		port:           "587",
		senderEmail:    settings.SenderEmail,
		senderPassword: settings.SenderPassword,
	}
	newEmailSender.fullAddres = newEmailSender.address()
	auth := smtp.PlainAuth(
		"",
		newEmailSender.senderEmail,
		newEmailSender.senderPassword,
		newEmailSender.host,
	)
	newEmailSender.auth = auth
	return &newEmailSender
}
