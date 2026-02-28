package models

import "time"

type Event struct {
	Timestamp time.Time       `json:"timestamp"`
	Source    string          `json:"source"`
	EventType string          `json:"event_type"`
	UserID    string          `json:"user_id,omitempty"`
	IPAddress string          `json:"ip_address,omitempty"`
	Status    string          `json:"status,omitempty"`
	Metadata  map[string]any  `json:"metadata,omitempty"`
}