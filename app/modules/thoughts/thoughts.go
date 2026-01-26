package Thoughts

import (
	General "Nautilus/general"
	"encoding/json"
	"fmt"
)

func AddTought(toughtToSave ThoughtSaved) (error, int) {
	err := General.DB.Create(&toughtToSave).Error
	if err != nil {
		return err, 0
	}

	if toughtToSave.Id != nil {
		return nil, *toughtToSave.Id
	}
	return nil, 0
}

func GetThought(thoughtId int) (error, []ThoughtSaved) {
	var thoughtsSaved []ThoughtSaved

	db := General.DB.Model(&ThoughtSaved{})
	if thoughtId != 0 {
		db = db.Where("id = ?", thoughtId)
	}

	err := db.Order("id asc").Find(&thoughtsSaved).Error
	return err, thoughtsSaved
}

func GetThoughtByErrorId(errorId int) (error, []ThoughtSaved) {
	if errorId == 0 {
		return fmt.Errorf("errorid = 0"), nil
	}

	query := `
	SELECT
		t.id,
		t.id_errors,
		t.creator_id,
		t.thought,
		t.created_in,
		t.vinculated_to,
		t.files,
		u.name as creator_name,
		u.profile_picture as creator_profile_picture
	FROM thoughts t
	LEFT JOIN users u ON u.id = t.creator_id
	WHERE t.id_errors = ?
	ORDER BY t.id ASC
	`
	var rows []thoughtRow
	err := General.DB.Raw(query, errorId).Scan(&rows).Error
	if err != nil {
		return err, nil
	}

	var result []ThoughtSaved
	for _, r := range rows {
		t := r.ThoughtSaved

		if len(r.FilesJSON) > 0 {
			var files []string
			_ = json.Unmarshal(r.FilesJSON, &files)
			t.Files = &files
		}

		result = append(result, t)
	}

	return nil, result
}

func UpdateThought(toughtToSave ThoughtSaved) error {
	if toughtToSave.Id == nil {
		return fmt.Errorf("thought id is nil")
	}

	return General.DB.
		Model(&ThoughtSaved{}).
		Where("id = ?", *toughtToSave.Id).
		Updates(map[string]interface{}{
			"id_errors":  toughtToSave.Id_errors,
			"creator_id": toughtToSave.Creator_id,
			"thought":    toughtToSave.Thought,
		}).Error
}

func DeleteThought(thoughtId int) error {
	return General.DB.Delete(&ThoughtSaved{}, thoughtId).Error
}
