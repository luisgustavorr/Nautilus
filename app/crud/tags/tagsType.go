package Tags

type TagSaved struct {
	Id          *int   `json:"id"`
	Id_apps     int    `json:"id_apps"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	Background  string `json:"background"`
	Type        int    `json:"type"`
}

func NewTagToSave(id_app int, name string, description string, opts ...func(*TagSaved)) *TagSaved {
	e := &TagSaved{
		Id_apps:     id_app,
		Name:        name,
		Description: description,
		Color:       "white",
		Background:  "red",
		Type:        0,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithId(v int) func(*TagSaved) {
	return func(e *TagSaved) {
		e.Id = &v
	}
}
func WithType(v int) func(*TagSaved) {
	return func(e *TagSaved) {
		e.Type = v
	}
}
func WithColor(c string) func(*TagSaved) {
	return func(e *TagSaved) {
		e.Color = c
	}
}
func WithBackground(c string) func(*TagSaved) {
	return func(e *TagSaved) {
		e.Background = c
	}
}
