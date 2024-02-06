package main

import (
	"fmt"
	"net/smtp"
)

func sendMailSimple(subject string, body string, to []string) {
	auth := smtp.PlainAuth(
		"",
		"joyride.projects@gmail.com",
		"lqbg uneu rmvr mvmu",
		"smtp.gmail.com",
	)

	msg := "Subject: " + subject + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"joyride.projects@gmail.com",
		to,
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sendMailSimple("Valentine Date", "17th Febuary 2024", []string{"sumfeis3@gmail.com"})
}
