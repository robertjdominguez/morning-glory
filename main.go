package main

import (
	"dominguezdev.com/morning-glory/gcal"
	"dominguezdev.com/morning-glory/gmail"
	"dominguezdev.com/morning-glory/openai"
	"dominguezdev.com/morning-glory/todoist"
)

func main() {
	events, _ := gcal.ShapeEvents()
	tasks, _ := todoist.ShapeTasks()
	message, _ := openai.GenerateMessage(events, tasks)
	gmail.SendMessage(message)
}
