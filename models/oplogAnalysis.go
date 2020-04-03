package models

import "time"

type OplogAnalysis struct {
	Timestamp  time.Time `json:"timestamp"`
	Database   string    `json:"database"`
	Collection string    `json:"collection"`
	Operation  string    `json:"operation"`
	Query      string    `json:"query"`
}
