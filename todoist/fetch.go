package todoist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"dominguezdev.com/morning-glory/config"
	"dominguezdev.com/morning-glory/types"
)

func fetchTasks(client *http.Client) (tasks []types.Task, retrievalError error) {
	// Load the config values
	appConfig, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("error loading config values: %v", err)
	}

	url := fmt.Sprintf("https://api.todoist.com/rest/v2/tasks?project_id=%s", appConfig.TodoistProjectId)

	// Create the request
	req, err := createRequest(url, appConfig.TodoistApiToken)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the status code is 200 OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	err = json.Unmarshal(body, &tasks)
	if err != nil {
		return nil, fmt.Errorf("error parsing JSOn: %v", err)
	}

	return tasks, nil
}

// Helper function to create the request itself
func createRequest(url, apiToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
