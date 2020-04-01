package treehole

import "time"

type ID uint

type Neighbour struct {
	Prev ID `json:"prev_id"`
	Next ID `json:"next_id"`
}

type Comment struct {
	ID      int
	Date    time.Time
	Author  string
	Content string
}

type Treehole struct {
	ID        ID
	Neighbour Neighbour
	Author    string
	Content   string
	Comments  *[]Comment
}
