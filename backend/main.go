package main

import (
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

// EmailRequest struct defines the JSON request body structure
type EmailRequest struct {
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
	To      []string `json:"to"`
}

// sendMailSimple sends an email using SMTP.
func sendMailSimple(subject string, body string, to []string) error {
	auth := smtp.PlainAuth(
		"",
		"joyride.projects@gmail.com", // Use your email
		"lqbg uneu rmvr mvmu",        // Use your password
		"smtp.gmail.com",
	)

	msg := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"joyride.projects@gmail.com", // Use your email
		to,
		msg,
	)

	return err
}

func main() {
	router := gin.Default()

	// Endpoint to send an email
	router.POST("/send-email", func(c *gin.Context) {
		var emailReq EmailRequest

		// Bind the JSON to the struct
		if err := c.ShouldBindJSON(&emailReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the sendMailSimple function
		err := sendMailSimple(emailReq.Subject, emailReq.Body, emailReq.To)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully"})
	})

	// Run the server
	router.Run(":8080")
}
