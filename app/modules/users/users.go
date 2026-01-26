package Users

import (
	General "Nautilus/general"
	"fmt"
)

func AddUser(userToSave UserSaved) (error, int) {
	err := General.DB.Create(&userToSave).Error
	if err != nil {
		return err, 0
	}

	if userToSave.Id != nil {
		return nil, *userToSave.Id
	}
	return nil, 0
}

func GetUsers(userId int) (error, []UserSaved) {
	usersSaved := []UserSaved{}

	db := General.DB.Model(&UserSaved{})
	if userId != 0 {
		db = db.Where("id = ?", userId)
	}

	err := db.Find(&usersSaved).Error
	return err, usersSaved
}

func UpdateUser(userToSave UserSaved) error {
	if userToSave.Id == nil {
		return fmt.Errorf("user id is nil")
	}

	return General.DB.
		Model(&UserSaved{}).
		Where("id = ?", *userToSave.Id).
		Updates(map[string]interface{}{
			"id_apps":          userToSave.Id_apps,
			"name":             userToSave.Name,
			"description":      userToSave.Description,
			"role":             userToSave.Role,
			"permission_level": userToSave.Permission_level,
			"profile_picture":  userToSave.Profile_picture,
		}).Error
}

func DeleteUser(userId int) error {
	if userId == 0 {
		return fmt.Errorf("userId = 0")
	}
	return General.DB.Delete(&UserSaved{}, userId).Error
}

func GetUserIdByToken(token string) (error, int) {
	return nil, General.ToInt(token)
}
