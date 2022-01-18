package main

import (
	"log"
	"net/smtp"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

const ()

var toAddress = []string{"guto.ps34@gmail.com", "gustavocioccari@gmail.com"}

func SendEmail(message string, toAddress []string) (response bool, err error) {
	loadEnv()

	fromAddress := os.Getenv("EMAIL")
	fromEmailPassword := os.Getenv("PASSWORD")
	smtpServer := os.Getenv("SMTP_SERVER")
	smptPort := os.Getenv("SMTP_PORT")

	var auth = smtp.PlainAuth("", fromAddress, fromEmailPassword, smtpServer)
	err = smtp.SendMail(smtpServer+":"+smptPort, auth, fromAddress, toAddress, []byte(message))
	log.Println(err)
	if err == nil {
		return true, nil
	}
	return false, err

}

func main() {
	subject := "Subject: Subject\n\n"
	body := "Body"
	message := strings.Join([]string{subject, body}, " ")
	SendEmail(message, toAddress)
}
