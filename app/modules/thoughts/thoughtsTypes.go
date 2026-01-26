package Thoughts

import (
	"html/template"
	"time"
)

type ThoughtSaved struct {
	Id            *int          `json:"id" gorm:"column:id;primaryKey"`
	Id_errors     int           `json:"id_errors" gorm:"column:id_errors"`
	Creator_id    int           `json:"creator_id" gorm:"column:creator_id"`
	Thought       template.HTML `json:"thought" gorm:"column:thought"`
	Created_in    time.Time     `json:"created_in" gorm:"column:created_in"`
	Vinculated_to *int          `json:"vinculated_to" gorm:"column:vinculated_to"`
	Files         *[]string     `json:"files" gorm:"column:files;type:jsonb"`

	// not native, recovered via QUERIES
	Creator_name            string `json:"creator_name" gorm:"creator_name"`
	Creator_profile_picture string `json:"creator_profile_picture" gorm:"creator_profile_picture"`
}

func (ThoughtSaved) TableName() string {
	return "thoughts"
}

type thoughtRow struct {
	ThoughtSaved
	FilesJSON []byte `gorm:"column:files"`
}

func NewToughtToSave(errorId int, creatorID int, thought string, opts ...func(*ThoughtSaved)) *ThoughtSaved {

	e := &ThoughtSaved{
		Id_errors:  errorId,
		Creator_id: creatorID,
		Thought:    template.HTML(thought),
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
