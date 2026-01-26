package Tags

import (
	General "Nautilus/general"
	"fmt"
)

func AddTag(tagToSave TagSaved) (error, int) {
	err := General.DB.Create(&tagToSave).Error
	if err != nil {
		return err, 0
	}

	if tagToSave.Id != nil {
		return nil, *tagToSave.Id
	}
	return nil, 0
}
func GetTags(tagId int) (error, []TagSaved) {
	tagsSaveds := []TagSaved{}

	db := General.DB.Model(&TagSaved{})
	if tagId != 0 {
		db = db.Where("id = ?", tagId)
	}

	err := db.Find(&tagsSaveds).Error
	return err, tagsSaveds
}

func GetErrorTagsByAppId(app_id int) (error, []TagSaved) {
	tagsSaveds := []TagSaved{}

	err := General.DB.
		Where("id_apps = ? AND type = 0", app_id).
		Find(&tagsSaveds).
		Error

	return err, tagsSaveds
}

func UpdateTag(tagToSave TagSaved) error {
	if tagToSave.Id == nil {
		return fmt.Errorf("tag id is nil")
	}

	return General.DB.
		Model(&TagSaved{}).
		Where("id = ?", *tagToSave.Id).
		Updates(map[string]interface{}{
			"id_apps":     tagToSave.Id_apps,
			"name":        tagToSave.Name,
			"description": tagToSave.Description,
			"color":       tagToSave.Color,
			"background":  tagToSave.Background,
			"type":        tagToSave.Type,
		}).Error
}

func DeleteTag(tagId int) error {
	if tagId == 0 {
		return fmt.Errorf("tagId = 0")
	}
	return General.DB.Delete(&TagSaved{}, tagId).Error
}
