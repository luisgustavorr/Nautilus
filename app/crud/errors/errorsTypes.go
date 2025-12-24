package Errors

import (
	"time"
)

type ErrorSaved struct {
	Id                *int      `json:"id"`
	Id_apps           int       `json:"id_apps"`
	Message           string    `json:"message"`
	Title             string    `json:"title"`
	Verified          bool      `json:"verified"`
	Error_level       int       `json:"error_level"`
	Creator_id        int       `json:"creator_id"`
	Created_in        time.Time `json:"created_in"`
	Last_edited_in    time.Time `json:"last_edited_in"`
	How_to_reproduce  string    `json:"how_to_reproduce"`
	Error_occurred_in time.Time `json:"error_occurred_in"`
}

func NewErrorToSave(appID int, creatorID int, title string, message string, opts ...func(*ErrorSaved)) *ErrorSaved {
	e := &ErrorSaved{
		Id_apps:           appID,
		Creator_id:        creatorID,
		Title:             title,
		Message:           message,
		Verified:          false,
		Error_level:       3,
		Created_in:        time.Now(),
		Last_edited_in:    time.Now(),
		Error_occurred_in: time.Now(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithVerified(v bool) func(*ErrorSaved) {
	return func(e *ErrorSaved) {
		e.Verified = v
	}
}
func WithId(v int) func(*ErrorSaved) {
	return func(e *ErrorSaved) {
		e.Id = &v
	}
}

func WithErrorLevel(l int) func(*ErrorSaved) {
	return func(e *ErrorSaved) {
		e.Error_level = l
	}
}

func WithOccurredAt(t time.Time) func(*ErrorSaved) {
	return func(e *ErrorSaved) {
		e.Error_occurred_in = t
	}
}

// ## ERROR TAGS ##

// ## ERROR FILES ##
