// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewTask struct {
	Text string   `json:"Text"`
	Tags []string `json:"Tags,omitempty"`
}

type Query struct {
}

type Task struct {
	ID   string   `json:"Id"`
	Text string   `json:"Text"`
	Tags []string `json:"Tags,omitempty"`
}
