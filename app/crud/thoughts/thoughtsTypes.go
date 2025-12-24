package Thoughts

type ThoughtSaved struct {
	Id         *int   `json:"id"`
	Id_errors  int    `json:"id_errors"`
	Creator_id int    `json:"creator_id"`
	Thought    string `json:"thought"`
}

func NewToughtToSave(errorId int, creatorID int, thought string, opts ...func(*ThoughtSaved)) *ThoughtSaved {
	e := &ThoughtSaved{
		Id_errors:  errorId,
		Creator_id: creatorID,
		Thought:    thought,
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
