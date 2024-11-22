package calendar

import (
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	// Embed a Date using an anonymous field
	// Embedding a Date in the Event type will not cause the Date fields to be promoted to the Event
	// because the Date fields are unexported. Only exported fields can be promoted from embedded types
	Date
}

func (e *Event) Title() string {
	return e.title
}

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("invalid title")
	}
	e.title = title
	return nil
}
