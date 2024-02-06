package main

import (
	"log"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

func sendMailSimple() {
	// Set up your email details
	from := "senderemail"
	password := "password"
	to := []string{"receiveremail"}
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	subject := "Will you be my Valentine?"
	body := "I am so glad you accepted my invitation! Looking forward to our date."
	message := []byte("Subject: " + subject + "\r\n\r\n" + body)

	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		log.Fatalf("sendMailSimple failed: %s", err)
	}
}

func main() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Setup route
	r.POST("/send-valentine-email", func(c *gin.Context) {
		sendMailSimple()
		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	})

	// Run the server
	r.Run(":8080")
}
