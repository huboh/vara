package event

// Listener represents a single event listener for a specific event type.
type Listener struct {
	// Once specifies if the listener function (Func) should be removed
	// after handling the event once. If true, the listener will be
	// unregistered automatically after the first execution.
	Once bool

	// Func is the function that will be called when the specified event occurs.
	Func ListenerFunc

	// Event is the name of the event this listener is subscribed to.
	Event string

	// Async indicates if the listener function (Func) should be called
	// asynchronously, allowing other listeners for the same event to be
	// triggered concurrently.
	Async bool

	// Priority defines the order in which listeners are executed for a given event.
	// Listeners with a higher priority are called before those with lower priority.
	// The default execution order is based on a descending priority value.
	Priority int
}

// ListenerFunc is a function that handles an event
type ListenerFunc func(Event) error
