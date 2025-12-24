package Users

type UserSaved struct {
	Id               *int   `json:"id"`
	Id_apps          int    `json:"id_apps"`
	Name             string `json:"name"`
	Description      string `json:"description"`
	Role             string `json:"role"`
	Permission_level int    `json:"permission_level"`
}

func NewUserSaved(id_apps int, name string, description string, opts ...func(*UserSaved)) *UserSaved {
	e := &UserSaved{
		Id_apps:          id_apps,
		Name:             name,
		Description:      description,
		Role:             "Usuário",
		Permission_level: 0,
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}
func WithId(id int) func(*UserSaved) {
	return func(e *UserSaved) {
		e.Id = &id
	}
}
func WithRole(r string) func(*UserSaved) {
	return func(e *UserSaved) {
		e.Role = r
	}
}
func WithPermissionLevel(p int) func(*UserSaved) {
	return func(e *UserSaved) {
		e.Permission_level = p
	}
}
