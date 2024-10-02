package config

import (
	"fmt"
	"os"

	"dominguezdev.com/morning-glory/types"
	"github.com/joho/godotenv"
)

// LoadConfig ensures env vars are loaded and are then made available to various
// callers in the program. It takes no arguments and returns a struct of Config.
func LoadConfig() (*types.Config, error) {
	// Attempt to load the .env file, but ignore errors if it doesn't exist
	_ = godotenv.Load(".env")

	// Check for required environment variables
	apiToken := os.Getenv("TODOIST_API_KEY")
	if apiToken == "" {
		return nil, fmt.Errorf("error TODOIST_API_KEY not found in env")
	}

	projectId := os.Getenv("PROJECT_ID")
	if projectId == "" {
		return nil, fmt.Errorf("error PROJECT_ID not found in env")
	}

	googleConfigFile := os.Getenv("GOOGLE_CREDENTIALS")
	if googleConfigFile == "" {
		return nil, fmt.Errorf("error GOOGLE_CREDENTIALS not found in env")
	}

	openAiApiKey := os.Getenv("OPENAI_API_KEY")
	if openAiApiKey == "" {
		return nil, fmt.Errorf("error OPENAI_API_KEY not found in env")
	}

	prompt := os.Getenv("PROMPT")
	if prompt == "" {
		return nil, fmt.Errorf("error PROMPT not found in env")
	}

	// Return a pointer to the types.Config struct with the right variables
	return &types.Config{
		TodoistApiToken:  apiToken,
		TodoistProjectId: projectId,
		GoogleConfigFile: googleConfigFile,
		OpenAiApiKey:     openAiApiKey,
		Prompt:           prompt,
	}, nil
}
