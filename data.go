package main

import "strconv"

type Task struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Priority string `json:"priority"`
	Status   string `json:"status"`
	Assignee string `json:"assignee"`
}

var tasks = []Task{
	{ID: "1", Title: "Write Medium Article", Priority: strconv.Itoa(3), Status: "In Progress", Assignee: "Vesna M"},
	{ID: "2", Title: "Code Review", Priority: strconv.Itoa(1), Status: "In Progress", Assignee: "Jane S"},
	{ID: "3", Title: "Feature Implementation", Priority: strconv.Itoa(3), Status: "To Do", Assignee: "Bob J"},
	{ID: "4", Title: "Bug Fixing", Priority: strconv.Itoa(1), Status: "In Progress", Assignee: "John D"},
	{ID: "5", Title: "Documentation Update", Priority: strconv.Itoa(2), Status: "To Do", Assignee: "Chris B"},
}
