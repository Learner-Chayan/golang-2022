//mail sending using go-simple-mail
package main

import (
	"log"

	mail "github.com/xhit/go-simple-mail/v2"
)

var htmlBody = `
<html>
<head>
   <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
   <title>Hello, World</title>
</head>
<body>
   <p>This is an email using Go</p>
</body>
`

func main() {
	server := mail.NewSMTPClient()
	//server.Host = "smtp.host.com"
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = "roy.manik.ru@gmail.com"
	server.Password = "manik12.ru.psychology#873$@"
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Create email
	email := mail.NewMSG()
	email.SetFrom("roy.manik.ru@gmail.com")
	email.AddTo("learnerchayan@gmail.com")
	email.AddCc("ashikurcmt125@gmail.com")
	email.SetSubject("New Go Email")

	email.SetBody(mail.TextHTML, htmlBody)
	email.AddAttachment("testPDF.pdf")

	// Send email
	err = email.Send(smtpClient)
	if err != nil {
		log.Fatal(err)
	}
}

/// mail sending using net/smtp package

// package main

// import (
// 	"fmt"
// 	"net/smtp"
// )

// func main() {

// 	// Sender data.
// 	from := "<your gmail >"
// 	password := "<gmail password>"

// 	// Receiver email address.
// 	to := []string{
// 		"learnerchayan@gmail.com",
// 	}

// 	// smtp server configuration.
// 	smtpHost := "smtp.gmail.com"
// 	smtpPort := "587"

// 	// Message.
// 	message := []byte("This is a test email message.")

// 	// Authentication.
// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent Successfully!")
// }
