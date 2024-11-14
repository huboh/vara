package event

import "time"

type Config struct {
	// BufferSize is the size of the async event buffer
	BufferSize int

	// AsyncTimeout is the timeout for async event processing
	AsyncTimeout time.Duration
}

func NewConfig() Config {
	return Config{
		BufferSize:   100,
		AsyncTimeout: 5 * time.Second,
	}
}
