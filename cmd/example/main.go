package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/brennosalves/go_email/email"
)

// .ENV VARIABLES
var (
	SMTP_SERVER   string
	SMTP_PORT     int
	SMTP_USER     string
	SMTP_PASSWORD string
)

func main() {
	// LOAD PROGRAM CONFIGURATION
	if resp := loadConfiguration(); resp != nil {
		log.Fatal(resp)
	}

	// CONFIGURES THE E-MAIL CREDENTIALS BASED ON THE ENVIRONMENT
	emailCredentials := email.EmailCredentials{
		SMTPServer:   SMTP_SERVER,
		SMTPPort:     SMTP_PORT,
		SMTPUser:     SMTP_USER,
		SMTPPassword: SMTP_PASSWORD,
	}

	// CONFIGURES THE E-MAIL DATA TO BE SEND
	emailData := email.EmailData{
		To:      "recipient@example.com.br",
		Subject: "This is an e-mail example",
		Body:    "Add your e-mail body here",
	}

	err := email.SendEmail(emailCredentials, emailData)
	if err != nil {
		log.Fatal(("The e-mail couldn't be sent. Details: " + err.Error()))
	}

	log.Println("Program finished successfully")
}

func loadConfiguration() []map[string]string {
	var errors []map[string]string
	// CHECK IF THERES A PROBLEM WHILE LOADING THE .ENV FILE
	err := godotenv.Load()
	if err != nil {
		errors = append(errors, map[string]string{"Parameter": "General", "Message": err.Error()})
		return errors
	}

	// GET THE VALUES FROM THE ENVIRONMENT
	SMTP_SERVER = os.Getenv("SMTP_SERVER")
	SMTP_PORT_STR := os.Getenv("SMTP_PORT")
	SMTP_USER = os.Getenv("SMTP_USER")
	SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

	// CHECK IF THE PARAMETERS WERE FILLED CORRECTLY
	if SMTP_SERVER == "" {
		errors = append(errors, map[string]string{"Parameter": "SMTP_SERVER", "Message": "Not filled."})
	}
	if SMTP_PORT_STR == "" { // SMTP PORT WILL BE CONVERTED TO INT
		errors = append(errors, map[string]string{"Parameter": "SMTP_PORT", "Message": "Not filled."})
	} else {
		SMTP_PORT, err = strconv.Atoi(SMTP_PORT_STR)
		if err != nil {
			errors = append(errors, map[string]string{"Parameter": "SMTP_PORT", "Message": "Filled incorrectly, value must be integer."})
		}
	}
	if SMTP_USER == "" {
		errors = append(errors, map[string]string{"Parameter": "SMTP_USER", "Message": "Not filled."})
	}
	if SMTP_PASSWORD == "" {
		errors = append(errors, map[string]string{"Parameter": "SMTP_PASSWORD", "Message": "Not filled."})
	}

	// IF ANY ERRORS WERE FOUND, RETURN THEM
	if len(errors) > 0 {
		return errors
	}

	return nil
}
