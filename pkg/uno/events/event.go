package events

// Event holds information about what happened and
// some iportant data that would reflect on the app
type Event struct {
	Name string
}

// NewEvent creates new object of type Event
func NewEvent(name string) Event {
	return Event{
		Name: name,
	}
}
