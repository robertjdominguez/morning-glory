package gmail

import (
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/api/gmail/v1"
)

// sendEmail sends an email using the Gmail API service
func sendEmail(srv *gmail.Service, body string) error {
	var message strings.Builder

	message.WriteString(fmt.Sprintf("From: %s\r\n", "rob.dominguez@hasura.io"))
	message.WriteString(fmt.Sprintf("To: %s\r\n", "rob.dominguez@hasura.io"))
	message.WriteString(fmt.Sprintf("Subject: %s\r\n", "Morning Brief"))
	message.WriteString("Content-Type: text/plain; charset=\"UTF-8\"\r\n")
	// Blank line between headers and body
	message.WriteString("\r\n")
	message.WriteString(body)

	// Base64 encode the message (with URL encoding)
	rawMessage := base64.URLEncoding.EncodeToString([]byte(message.String()))

	gmailMessage := &gmail.Message{
		Raw: rawMessage,
	}

	_, err := srv.Users.Messages.Send("me", gmailMessage).Do()
	if err != nil {
		return fmt.Errorf("unable to send email: %v", err)
	}

	return nil
}
