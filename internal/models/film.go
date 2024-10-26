package models

import "time"

type Film struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Details     string    `json:"details"`
	Genre       string    `json:"genre"`
	Year        time.Time `json:"year"`
}
