package gmail

import (
	"context"
	"fmt"
	"log"

	"dominguezdev.com/morning-glory/config"
	"google.golang.org/api/gmail/v1" // Gmail API
	"google.golang.org/api/option"
)

func loginGmail() (*gmail.Service, error) {
	appConfig, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config values: %v", err)
	}

	credBytes := []byte(appConfig.GoogleConfigFile)
	ctx := context.Background()
	srv, err := gmail.NewService(ctx, option.WithCredentialsJSON(credBytes))
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	return srv, nil
}
