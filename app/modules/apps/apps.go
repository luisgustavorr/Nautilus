package Apps

import (
	General "Nautilus/general"
	"fmt"
)

func AddApp(appToSave AppSaved) (error, int) {
	insertedId := 0

	err := General.DB.QueryRow("INSERT INTO apps (name,perfil_image,description) VALUES ($1,$2,$3) RETURNING id", appToSave.Name, appToSave.Perfil_image, appToSave.Description).Scan(&insertedId)
	return err, insertedId
}
func GetApp(appId int) (error, []AppSaved) {
	apps := []AppSaved{}
	filter := ""
	if appId != 0 {
		filter = fmt.Sprintf("WHERE id = %d", appId)
	}
	query := fmt.Sprintf("SELECT id,name,perfil_image,description FROM apps %s", filter)
	rows, err := General.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var appSaved AppSaved
			err = rows.Scan(&appSaved.Id, &appSaved.Name, &appSaved.Perfil_image, &appSaved.Description)
			if err == nil {
				apps = append(apps, appSaved)
			}
		}
	}
	return err, apps
}
func DeleteApp(appId int) error {
	if appId == 0 {
		return fmt.Errorf("errorId = 0, falha ao apagar")
	}
	_, err := General.DB.Exec("DELETE FROM apps WHERE id = $1", appId)
	return err
}
func UpdateApp(appToSave AppSaved) error {
	_, err := General.DB.Exec("UPDATE apps SET name = $2,perfil_image =$3,description=$4 WHERE id = $1", appToSave.Id, appToSave.Name, appToSave.Perfil_image, appToSave.Description)
	return err
}
