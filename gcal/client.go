package gcal

import (
	"context"
	"fmt"
	"log"

	"dominguezdev.com/morning-glory/config"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// login() uses the config file to creaete a new Google calendar service for use in the application.
// It returns either the service or an error.
func login() (service *calendar.Service, loginError error) {
	// Load the .env file
	appConfig, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config values: %v", err)
	}

	// Create the Google Calendar service using the credentials file
	credBytes := []byte(appConfig.GoogleConfigFile)
	ctx := context.Background()
	srv, err := calendar.NewService(ctx, option.WithCredentialsJSON(credBytes))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	return srv, nil
}
