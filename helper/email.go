package helper

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"greebel.core.be/core"
)

type (
	SendGrid struct {
		SenderMail       string
		SenderName       string
		RecipientMail    string
		RecipientName    string
		Subject          string
		HTMLContent      string
		PlainTextContent string
	}
	SendGridwithAttch struct {
		SenderMail       string
		SenderName       string
		RecipientMail    string
		RecipientName    string
		Subject          string
		HTMLContent      string
		PlainTextContent string
		Attchment        string
	}
)

func SendMail(data SendGrid) error {
	from := mail.NewEmail(data.SenderName, data.SenderMail)
	subject := data.Subject
	to := mail.NewEmail(data.RecipientName, data.RecipientMail)
	plainTextContent := data.PlainTextContent
	htmlContent := data.HTMLContent
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(core.App.Config.SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	return nil
}

func SendMailwithAttachment(data SendGridwithAttch) error {
	m := mail.NewV3Mail()

	from := mail.NewEmail(data.SenderName, data.SenderMail)
	htmlContent := data.HTMLContent
	content := mail.NewContent("text/html", htmlContent)
	to := mail.NewEmail(data.RecipientName, data.RecipientMail)

	m.SetFrom(from)
	m.AddContent(content)

	// create new *Personalization
	personalization := mail.NewPersonalization()
	personalization.AddTos(to)
	personalization.Subject = data.Subject

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)
	// read/attach .pdf file
	a_pdf := mail.NewAttachment()
	dat, err := ioutil.ReadFile("helper/storage/" + data.Attchment + ".pdf")
	if err != nil {
		fmt.Println(err)
	}
	encoded := base64.StdEncoding.EncodeToString([]byte(dat))
	a_pdf.SetContent(encoded)
	a_pdf.SetType("application/pdf")
	a_pdf.SetFilename(data.Attchment + ".pdf")
	a_pdf.SetDisposition("attachment")
	m.AddAttachment(a_pdf)

	request := sendgrid.GetRequest(core.App.Config.SENDGRID_API_KEY, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)

	if err != nil {
		log.Println(err)
		return err
	}

	fmt.Println(response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	return nil
}
