package Thoughts

import "time"

type ThoughtSaved struct {
	Id            *int      `json:"id"`
	Id_errors     int       `json:"id_errors"`
	Creator_id    int       `json:"creator_id"`
	Thought       string    `json:"thought"`
	Created_in    time.Time `json:"created_in"`
	Vinculated_to *int      `json:"vinculated_to"`
	Files         *[]string `json:"files"`
	//not native, recovered via QUERIES
	Creator_name            string `json:"creator_name"`
	Creator_profile_picture string `json:"creator_profile_picture"`
}

func NewToughtToSave(errorId int, creatorID int, thought string, opts ...func(*ThoughtSaved)) *ThoughtSaved {
	e := &ThoughtSaved{
		Id_errors:  errorId,
		Creator_id: creatorID,
		Thought:    thought,
		Created_in: time.Now(),
		Files:      &[]string{},
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithId(id int) func(*ThoughtSaved) {
	return func(e *ThoughtSaved) {
		e.Id = &id
	}
}
