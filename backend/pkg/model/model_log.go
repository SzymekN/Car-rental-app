package model

import "time"

type Log struct {
	ID        int       `json:"id,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Key       string    `json:"key,omitempty"`
	Value     string    `json:"value,omitempty"`
}
