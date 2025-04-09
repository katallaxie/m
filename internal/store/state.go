package store

const (
	Initial = iota
	Loading
	Error
	Success
)

// State ...
type State struct {
	Status   int
	Error    error
	Messages []string
}

// NewState returns a new state.
func NewState() State {
	return State{
		Status:   Initial,
		Messages: make([]string, 0),
	}
}
