package todoist

import (
	"log"
	"time"

	"dominguezdev.com/morning-glory/types"
)

/*
* Here, we'll reduce the tasks to be only those due today
 */
func reduceTasks(tasks []types.Task) (reducedTasks []types.Task, err error) {
	// Figure out what today is and truncate it down to just the date
	today := time.Now().Truncate(24 * time.Hour)

	// Create a string layout for parsing datetime objects
	dateLayout := "2006-01-02"

	// Loop over 'em all
	for _, task := range tasks {
		// Let's check for "not set valus first"
		if task.Due != nil {
			parsedDate, err := time.Parse(dateLayout, task.Due.Date)
			if err != nil {
				log.Fatalf("Error parsing date: %v", err)
			}

			if parsedDate.Equal(today) {
				reducedTasks = append(reducedTasks, task)
			}

		}
	}

	return reducedTasks, nil
}
