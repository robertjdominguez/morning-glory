package todoist

import (
	"fmt"
	"net/http"
	"time"

	"dominguezdev.com/morning-glory/types"
	"github.com/briandowns/spinner"
)

func ShapeTasks() (tasks []types.Task, err error) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	s.Suffix = " Fetching Todoist tasks..."

	rawTasks, _ := fetchTasks(http.DefaultClient)

	reducedTasks, _ := reduceTasks(rawTasks)
	s.Suffix = " Isolating today's events..."

	s.Stop()
	fmt.Println("✅ Todoist items shaping completed!")

	return reducedTasks, nil
}
