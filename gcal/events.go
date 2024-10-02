package gcal

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

// retrieveEvents() takes in a pointer to a calendar.Service and returns the list of events or an error.
func retrieveEvents(srv *calendar.Service) ([]*calendar.Event, error) {
	// Retrieve the user's events from the calendar
	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(50).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next fifty of the user's events: %v", err)
	}

	return events.Items, nil
}

// reduceTodaysEvents() takes in a slice of events and returns only those that are for today
func reduceTodaysEvents(events []*calendar.Event) (reducedEvents []*calendar.Event, err error) {
	// Get the current date, truncating the time portion to midnight
	today := time.Now().Truncate(24 * time.Hour)

	dateLayout := "2006-01-02T15:04:05-07:00"

	for _, event := range events {
		// If the event has a DateTime
		if event.Start.DateTime != "" {
			parsedDate, err := time.Parse(dateLayout, event.Start.DateTime)
			if err != nil {
				log.Fatalf("error parsing date for event: %v", err)
			}

			if parsedDate.Truncate(24 * time.Hour).Equal(today) {
				reducedEvents = append(reducedEvents, event)
			}
		}
	}

	return reducedEvents, nil
}

// reducedConfirmedEvents() ensures we return only the events we've confirmed
func reduceConfirmedEvents(events []*calendar.Event) (reducedEvents []*calendar.Event, err error) {
	for _, event := range events {
		// If the event has a DateTime
		if event.Status == "confirmed" {
			reducedEvents = append(reducedEvents, event)
		}
	}

	return reducedEvents, nil
}

// printEvents() will log the events using unmarshal
func PrintEvents(events []*calendar.Event) {
	jsonData, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling event to JSON: %v", err)
	}
	fmt.Printf("Event as JSON:\n%s\n", string(jsonData))
}
