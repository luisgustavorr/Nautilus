package Apps

type AppSaved struct {
	Id           *int   `json:"id"`
	Name         string `json:"name"`
	Perfil_image string `json:"perfil_image"`
	Description  string `json:"description"`
}

func NewAppSaved(name string, perfil_image string, description string, opts ...func(*AppSaved)) *AppSaved {
	e := &AppSaved{
		Name:         name,
		Perfil_image: perfil_image,
		Description:  description,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithId(id int) func(*AppSaved) {
	return func(as *AppSaved) {
		as.Id = &id
	}
}
