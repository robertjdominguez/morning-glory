package types

import (
	"time"
)

type Task struct {
	ID           string    `json:"id"`
	AssignerID   *string   `json:"assigner_id"`
	AssigneeID   *string   `json:"assignee_id"`
	ProjectID    string    `json:"project_id"`
	SectionID    *string   `json:"section_id"`
	ParentID     *string   `json:"parent_id"`
	Order        int       `json:"order"`
	Content      string    `json:"content"`
	Description  string    `json:"description"`
	IsCompleted  bool      `json:"is_completed"`
	Labels       []string  `json:"labels"`
	Priority     int       `json:"priority"`
	CommentCount int       `json:"comment_count"`
	CreatorID    string    `json:"creator_id"`
	CreatedAt    time.Time `json:"created_at"`
	Due          *Due      `json:"due"`
	URL          string    `json:"url"`
	Duration     *string   `json:"duration"`
	Deadline     *string   `json:"deadline"`
}

// Due represents the due date for a task
type Due struct {
	Date        string `json:"date"`
	String      string `json:"string"`
	Lang        string `json:"lang"`
	IsRecurring bool   `json:"is_recurring"`
}
