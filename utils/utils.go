package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.Sender)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}

func SendMail(to []string, body string, subject string) error {
	from := os.Getenv("EMAIL")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	message := []byte(
		BuildMessage(Mail{
			Sender:  from,
			To:      to,
			Subject: subject,
			Body:    body,
		}),
	)

	authType := os.Getenv("SMTP_AUTH_TYPE")
	var auth smtp.Auth

	if authType == "PLAIN" {
		username := os.Getenv("SMTP_USERNAME")
		password := os.Getenv("SMTP_PASSWORD")

		auth = smtp.PlainAuth("", username, password, smtpHost)
	}

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		return err
	}

	log.Println("Email sent!")
	return nil
}

func HashPassword(password string) string {
	saltRounds := os.Getenv("SALT_ROUNDS")
	saltRoundsInt, err := strconv.Atoi(saltRounds)
	if err != nil {
		saltRoundsInt = 10
	}

	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRoundsInt)

	if err != nil {
		log.Fatal(err)
	}
	return string(bcryptPassword)
}

func ComparePassword(hashedPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

var RUNES = []rune("0123456789")

func GenerateVerificationCode() string {

	b := make([]rune, 6)
	for i := range b {
		b[i] = RUNES[rand.Intn(len(RUNES))]
	}
	return string(b)
}
