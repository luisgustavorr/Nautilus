package Users

type UserSaved struct {
	Id               *int   `gorm:"column:id;primaryKey"`
	Id_apps          int    `gorm:"column:id_apps"`
	Name             string `gorm:"column:name"`
	Description      string `gorm:"column:description"`
	Role             string `gorm:"column:role"`
	Permission_level int    `gorm:"column:permission_level"`
	Profile_picture  string `gorm:"column:profile_picture"`
}

func (UserSaved) TableName() string {
	return "users"
}

func NewUserSaved(id_apps int, name string, description string, opts ...func(*UserSaved)) *UserSaved {
	e := &UserSaved{
		Id_apps:          id_apps,
		Name:             name,
		Description:      description,
		Role:             "Usuário",
		Profile_picture:  "",
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
func WithProfile_picture(p string) func(*UserSaved) {
	return func(e *UserSaved) {
		e.Profile_picture = p
	}
}
