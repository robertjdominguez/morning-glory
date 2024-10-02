package todoist

import (
	"testing"
	"time"

	"dominguezdev.com/morning-glory/types"
)

// Test the reduceTasks function to ensure we're only returning today's tasks
func TestReduceTasks(t *testing.T) {
	// Define the layout for formatting dates just like we do in our function
	const dateLayout = "2006-01-02"

	// Get today's date, as we'll need it later
	today := time.Now().UTC().Truncate(24 * time.Hour).Format(dateLayout)

	// Mock that data with our types
	tasks := []types.Task{
		{ID: "1", Due: &types.Due{Date: today}},
		{ID: "2", Due: &types.Due{Date: todayMinus(1)}},
		{ID: "3", Due: nil},
	}

	// Call the function we're testing
	reducedTasks, err := reduceTasks(tasks)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	t.Logf("Number of tasks due today: %d", len(reducedTasks))
	// Assertions
	if len(reducedTasks) != 1 {
		t.Errorf("Expected 1 task due today, got %d", len(reducedTasks))
	}
	if reducedTasks[0].Due.Date != today {
		t.Errorf("Expected task due today, got due date: %v", reducedTasks[0].Due.Date)
	}
}

// Test the fetcher by mocking the Todoist API to ensure we're pulling in all the tasks for the project
func TestFetchTasks(t *testing.T) {
}

// Helper function to get the date `n` days ago in the correct format
func todayMinus(days int) string {
	return time.Now().AddDate(0, 0, -days).Truncate(24 * time.Hour).Format("2006-01-02")
}
