package Errors

import (
	Tags "Nautilus/app/modules/tags"
	General "Nautilus/general"
	"encoding/json"
	"fmt"
)

func AddError(errorToSave ErrorSaved) (error, int) {
	err := General.DB.Create(&errorToSave).Error
	if err != nil {
		return err, 0
	}

	if errorToSave.Id != nil {
		return nil, *errorToSave.Id
	}
	return nil, 0
}

func GetErrors(errorId int) (error, []ErrorSaved) {
	filter := ""
	args := []interface{}{}

	if errorId != 0 {
		filter = "WHERE e.id = ?"
		args = append(args, errorId)
	}

	query := fmt.Sprintf(`
	SELECT
	  e.id,
	  e.id_apps,
	  e.message,
	  e.title,
	  e.verified,
	  e.error_level,
	  e.creator_id,
	  e.created_in,
	  e.last_edited_in,
	  e.how_to_reproduce,
	  e.error_occurred_in,
	  t.tags,
	  u.name as creator_name,
	  f.files
	FROM errors e
	LEFT JOIN (
	  SELECT
	    et.id_errors,
	    jsonb_agg(
	      json_build_object(
	        'id', t.id,
	        'id_apps', t.id_apps,
	        'name', t.name,
	        'description', t.description,
	        'color', t.color,
	        'background', t.background
	      )
	    ) AS tags
	  FROM errors_tags et
	  LEFT JOIN tags t ON t.id = et.id_tags
	  GROUP BY et.id_errors
	) t ON t.id_errors = e.id
	LEFT JOIN users u ON u.id = e.creator_id
	LEFT JOIN (
	  SELECT id_errors, jsonb_agg(file_name) AS files
	  FROM errors_files
	  GROUP BY id_errors
	) f ON f.id_errors = e.id
	%s
	ORDER BY e.id DESC
	`, filter)

	var rows []errorRow
	err := General.DB.Raw(query, args...).Scan(&rows).Error
	if err != nil {
		return err, nil
	}

	var result []ErrorSaved
	for _, r := range rows {
		e := r.ErrorSaved

		if len(r.TagsJSON) > 0 {
			var tags []Tags.TagSaved
			_ = json.Unmarshal(r.TagsJSON, &tags)
			e.Tags = &tags
		}

		if len(r.FilesJSON) > 0 {
			var files []string
			_ = json.Unmarshal(r.FilesJSON, &files)
			e.Files = &files
			e.Files_count = len(files)
		}

		result = append(result, e)
	}

	return nil, result
}
func GetErrorsByAppId(appId int) (error, []ErrorSaved) {
	if appId == 0 {
		return fmt.Errorf("appId = 0, falha ao selecionar errors"), nil
	}

	query := `
	SELECT
	  e.id,
	  e.id_apps,
	  e.message,
	  e.title,
	  e.verified,
	  e.error_level,
	  e.creator_id,
	  e.created_in,
	  e.last_edited_in,
	  e.how_to_reproduce,
	  e.error_occurred_in,
	  t.tags,
	  u.name as creator_name,
	  f.files
	FROM errors e
	LEFT JOIN (
	  SELECT
	    et.id_errors,
	    jsonb_agg(
	      json_build_object(
	        'id', t.id,
	        'id_apps', t.id_apps,
	        'name', t.name,
	        'description', t.description,
	        'color', t.color,
	        'background', t.background
	      )
	    ) AS tags
	  FROM errors_tags et
	  LEFT JOIN tags t ON t.id = et.id_tags
	  GROUP BY et.id_errors
	) t ON t.id_errors = e.id
	LEFT JOIN users u ON u.id = e.creator_id
	LEFT JOIN (
	  SELECT id_errors, jsonb_agg(file_name) AS files
	  FROM errors_files
	  GROUP BY id_errors
	) f ON f.id_errors = e.id
	 WHERE e.id_apps = ? ORDER BY e.id DESC`

	var rows []errorRow
	err := General.DB.Raw(query, appId).Scan(&rows).Error
	if err != nil {
		return err, nil
	}

	var result []ErrorSaved
	for _, r := range rows {
		e := r.ErrorSaved

		if len(r.TagsJSON) > 0 {
			var tags []Tags.TagSaved
			_ = json.Unmarshal(r.TagsJSON, &tags)
			e.Tags = &tags
		}

		if len(r.FilesJSON) > 0 {
			var files []string
			_ = json.Unmarshal(r.FilesJSON, &files)
			e.Files = &files
			e.Files_count = len(files)
		}

		result = append(result, e)
	}

	return nil, result
}

func UpdateErrors(errorToSave ErrorSaved) error {
	if errorToSave.Id == nil {
		return fmt.Errorf("error id is nil")
	}

	return General.DB.
		Model(&ErrorSaved{}).
		Where("id = ?", *errorToSave.Id).
		Updates(map[string]interface{}{
			"id_apps":           errorToSave.Id_apps,
			"message":           errorToSave.Message,
			"title":             errorToSave.Title,
			"verified":          errorToSave.Verified,
			"error_level":       errorToSave.Error_level,
			"creator_id":        errorToSave.Creator_id,
			"created_in":        errorToSave.Created_in,
			"last_edited_in":    errorToSave.Last_edited_in,
			"how_to_reproduce":  errorToSave.How_to_reproduce,
			"error_occurred_in": errorToSave.Error_occurred_in,
		}).Error
}

func DeleteError(errorId int) error {
	if errorId == 0 {
		return fmt.Errorf("errorId = 0, falha ao apagar")
	}
	return General.DB.Delete(&ErrorSaved{}, errorId).Error
}
