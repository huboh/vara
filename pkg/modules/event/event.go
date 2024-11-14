package event

import (
	"context"
	"time"
)

// Event represents a single event instance
type Event struct {
	// Payload for the event
	Payload any

	// Metadata about the event
	Metadata EventMetadata
}

// EventMetadata contains information about the event
type EventMetadata struct {
	// Name of the event
	Name string

	// Additional context data
	Context context.Context

	// Time the event was created
	CreatedAt time.Time
}
