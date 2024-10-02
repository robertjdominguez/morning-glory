package gcal

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"google.golang.org/api/calendar/v3"
)

func ShapeEvents() (events []*calendar.Event, err error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()

	s.Suffix = " Logging into Google services..."
	srv, err := login()
	if err != nil {
		s.Stop()
		return nil, fmt.Errorf("error logging into Google services: %v", err)
	}

	s.Suffix = " Retrieving events..."
	returnedEvents, err := retrieveEvents(srv)
	if err != nil {
		s.Stop()
		return nil, fmt.Errorf("error retrieving events: %v", err)
	}

	s.Suffix = " Filtering today's events..."
	todaysEvents, err := reduceTodaysEvents(returnedEvents)
	if err != nil {
		s.Stop()
		return nil, fmt.Errorf("error reducing events: %v", err)
	}

	s.Suffix = " Filtering confirmed events..."
	reducedEvents, err := reduceConfirmedEvents(todaysEvents)
	if err != nil {
		s.Stop()
		return nil, fmt.Errorf("error reducing confirmed events: %v", err)
	}

	s.Stop()
	fmt.Println("✅ Events shaping completed!")
	return reducedEvents, nil
}
