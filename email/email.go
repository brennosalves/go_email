package email

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

// STRUCT CONTAINING THE E-MAIL CREDENTIALS
type EmailCredentials struct {
	SMTPServer   string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

// STRUCT CONTAINING THE E-MAIL DATA
type EmailData struct {
	To      string
	Subject string
	Body    string
}

// FUNCTION RESPONSIBLE TO SEND THE E-MAIL
func SendEmail(emailCredentials EmailCredentials, emailData EmailData) error {
	// CHECK IF THE CREDENTIALS WERE INFORMED CORRECTLY
	{
		if emailCredentials.SMTPServer == "" {
			return fmt.Errorf("the smtp server address was not informed")
		}
		if emailCredentials.SMTPPort <= 0 {
			return fmt.Errorf("the smtp port was not informed")
		}
		if emailCredentials.SMTPUser == "" {
			return fmt.Errorf("the smtp user was not informed")
		}
		if emailCredentials.SMTPPassword == "" {
			return fmt.Errorf("the smtp password was not informed")
		}
	}

	// CHECK IF THE E-MAIL DATA WAS INFORMED CORRECTLY
	{
		if emailData.To == "" {
			return fmt.Errorf("the e-mail recipient list was empty")
		}
		if emailData.Subject == "" {
			return fmt.Errorf("the e-mail subject was not informed")
		}
		if emailData.Body == "" {
			return fmt.Errorf("the e-mail body is empty")
		}
	}

	// MOUNT MESSAGE
	msg := gomail.NewMessage()
	msg.SetHeader("From", emailCredentials.SMTPUser)
	msg.SetHeader("To", emailData.To)
	msg.SetHeader("Subject", emailData.Subject)
	msg.SetBody("text/html", emailData.Body)

	// CONNECT TO THE SERVER
	dialer := gomail.NewDialer(emailCredentials.SMTPServer, emailCredentials.SMTPPort, emailCredentials.SMTPUser, emailCredentials.SMTPPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := dialer.DialAndSend(msg); err != nil {
		return err
	}

	return nil
}
