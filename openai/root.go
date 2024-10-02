package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"google.golang.org/api/calendar/v3"

	"dominguezdev.com/morning-glory/config"
	"dominguezdev.com/morning-glory/types"
	"github.com/briandowns/spinner"
)

// GenerateMessage() tases in tasks and events as arguments and sends them to the LLM.
func GenerateMessage(events []*calendar.Event, tasks []types.Task) (string, error) {
	appConfig, err := config.LoadConfig()
	if err != nil {
		return "", fmt.Errorf("error loading config values: %v", err)
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Suffix = " Preparing events and tasks..."

	// Shape the events and tasks by turning them into strings and dropping them in
	// TODO: Move this to a separate function and change the arguments for GenerateMessage()
	eventsSlice := []string{}
	for _, event := range events {
		json, err := json.Marshal(event)
		if err != nil {
			return "", fmt.Errorf("failed to parse events into JSON: %v", err)
		}
		eventsSlice = append(eventsSlice, string(json))
	}

	tasksSlice := []string{}
	for _, task := range tasks {
		json, err := json.Marshal(task)
		if err != nil {
			return "", fmt.Errorf("failed to parse tasks into JSON: %v", err)
		}
		tasksSlice = append(tasksSlice, string(json))
	}

	// TODO: Put this into a function:
	combinedSlice := append(eventsSlice, tasksSlice...)

	eventsAndTasks := strings.Join(combinedSlice, ", ")

	s.Suffix = " Pinging LLM..."
	// Prepare the request payload
	reqBody, _ := json.Marshal(types.OpenAIRequest{
		Model: "gpt-4",
		Messages: []types.Message{
			{Role: "system", Content: appConfig.Prompt},
			{Role: "user", Content: eventsAndTasks},
		},
	})

	// Create a new request to the OpenAI Chat API
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+appConfig.OpenAiApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response from OpenAI
	body, _ := io.ReadAll(resp.Body)
	var openAIResp types.OpenAIResponse
	json.Unmarshal(body, &openAIResp)

	s.Stop()
	fmt.Println("✅ Message generated!\n")

	return openAIResp.Choices[0].Message.Content, nil
}
