package Apps

import (
	General "Nautilus/general"
	"fmt"
)

func AddApp(appToSave AppSaved) (error, int) {
	err := General.DB.Create(&appToSave).Error
	if err != nil {
		return err, 0
	}

	if appToSave.Id != nil {
		return nil, *appToSave.Id
	}
	return nil, 0
}

func GetApp(appId int) (error, []AppSaved) {
	apps := []AppSaved{}

	db := General.DB.Model(&AppSaved{})
	if appId != 0 {
		db = db.Where("id = ?", appId)
	}

	err := db.Find(&apps).Error
	return err, apps
}

func DeleteApp(appId int) error {
	if appId == 0 {
		return fmt.Errorf("errorId = 0, falha ao apagar")
	}
	return General.DB.Delete(&AppSaved{}, appId).Error
}

func UpdateApp(appToSave AppSaved) error {
	if appToSave.Id == nil {
		return fmt.Errorf("app id is nil")
	}

	return General.DB.
		Model(&AppSaved{}).
		Where("id = ?", *appToSave.Id).
		Updates(map[string]interface{}{
			"name":         appToSave.Name,
			"perfil_image": appToSave.Perfil_image,
			"description":  appToSave.Description,
		}).Error
}
