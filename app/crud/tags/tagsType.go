package Tags

import "fmt"

type TagSaved struct {
	Id          *int   `json:"id"`
	Id_apps     int    `json:"id_apps"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

func NewTagToSave(id_app int, name string, description string, opts ...func(*TagSaved)) *TagSaved {
	e := &TagSaved{
		Id_apps:     id_app,
		Name:        name,
		Description: description,
		Color:       "red",
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithId(v int) func(*TagSaved) {
	fmt.Println("WithId", v)
	return func(e *TagSaved) {
		e.Id = &v
	}
}
func WithColor(c string) func(*TagSaved) {
	return func(e *TagSaved) {
		e.Color = c
	}
}
