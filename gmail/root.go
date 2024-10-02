package gmail

import "log"

func SendMessage(body string) {
	srv, err := loginGmail()
	if err != nil {
		log.Fatalf("Failed to create Gmail service: %v", err)
	}

	err = sendEmail(srv, body)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	log.Println("Email sent successfully!")
}
