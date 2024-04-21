package model

type Task struct {
	ID   string   `json:"Id"`
	Text string   `json:"Text"`
	Tags []string `json:"Tags,omitempty"`
}
