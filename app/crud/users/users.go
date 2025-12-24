package Users

import (
	General "Nautilus/general"
	"fmt"
)

func AddUser(userToSave UserSaved) (error, int) {
	var insertedId int
	var err error

	err = General.DB.QueryRow("INSERT INTO users (id_apps,name,description,role,permission_level) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		userToSave.Id_apps, userToSave.Name, userToSave.Description, userToSave.Role, userToSave.Permission_level).Scan(&insertedId)
	return err, insertedId
}
func GetUsers(userId int) (error, []UserSaved) {
	usersSaved := []UserSaved{}
	var err error
	filter := ""
	if userId != 0 {
		filter = fmt.Sprintf(" WHERE id = %d", userId)
	}
	query := fmt.Sprintf("SELECT  id,id_apps,name,description,role,permission_level FROM users %s", filter)
	rows, err := General.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			userSaved := UserSaved{}
			err = rows.Scan(&userSaved.Id, &userSaved.Id_apps, &userSaved.Name, &userSaved.Description, &userSaved.Role, &userSaved.Permission_level)
			if err == nil {
				usersSaved = append(usersSaved, userSaved)
			}
		}
	}
	return err, usersSaved
}

func UpdateUser(userToSave UserSaved) error {
	fmt.Println(General.JsonViewInterface(userToSave))
	_, err := General.DB.Exec("UPDATE users SET id_apps = $1,name =$2,description=$3,role=$4,permission_level=$5 WHERE id = $6",
		userToSave.Id_apps, userToSave.Name, userToSave.Description, userToSave.Role, userToSave.Permission_level, userToSave.Id)
	return err
}

func DeleteUser(userId int) error {
	if userId == 0 {
		return fmt.Errorf("userId = 0")
	}
	_, err := General.DB.Exec("DELETE from users WHERE id = $1", userId)
	return err
}
