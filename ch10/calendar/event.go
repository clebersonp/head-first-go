package calendar

type Event struct {
	Title string
	// Embed a Date using an anonymous field
	// Embedding a Date in the Event type will not cause the Date fields to be promoted to the Event
	// because the Date fields are unexported. Only exported fields can be promoted from embedded types
	Date
}
