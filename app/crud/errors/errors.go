package Errors

import (
	Tags "Nautilus/app/crud/tags"
	General "Nautilus/general"
	"encoding/json"
	"fmt"
)

func AddError(errorToSave ErrorSaved) (error, int) {
	query := `
	INSERT INTO errors (
		id_apps,
		message,
		title,
		verified,
		error_level,
		creator_id,
		created_in,
		last_edited_in,
		how_to_reproduce,
		error_occurred_in
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING id
	`
	insertedId := 0
	err := General.DB.QueryRow(
		query,
		errorToSave.Id_apps,
		errorToSave.Message,
		errorToSave.Title,
		errorToSave.Verified,
		errorToSave.Error_level,
		errorToSave.Creator_id,
		errorToSave.Created_in,
		errorToSave.Last_edited_in,
		errorToSave.How_to_reproduce,
		errorToSave.Error_occurred_in,
	).Scan(&insertedId)

	return err, insertedId
}
func GetErrors(errorId int) (error, []ErrorSaved) {
	filter := ""
	if errorId != 0 {
		filter = fmt.Sprintf("WHERE id = %d", errorId)
	}
	rows, err := General.DB.Query(fmt.Sprintf("SELECT id,id_apps,message,title,verified,error_level,creator_id,created_in,last_edited_in,how_to_reproduce,error_occurred_in FROM errors %s", filter))
	var errorsRecovereds []ErrorSaved
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var errorSelected ErrorSaved
			err = rows.Scan(&errorSelected.Id, &errorSelected.Id_apps, &errorSelected.Message, &errorSelected.Title, &errorSelected.Verified,
				&errorSelected.Error_level, &errorSelected.Creator_id, &errorSelected.Created_in, &errorSelected.Last_edited_in,
				&errorSelected.How_to_reproduce, &errorSelected.Error_occurred_in)
			if err == nil {
				errorsRecovereds = append(errorsRecovereds, errorSelected)
			}
		}
	}
	return err, errorsRecovereds
}
func GetErrorsByAppId(appId int) (error, []ErrorSaved) {
	if appId == 0 {
		return fmt.Errorf("appId = 0, falha ao selecionar errors"), []ErrorSaved{}
	}
	rows, err := General.DB.Query(`SELECT
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
  u.name as creator_name
FROM
  errors e
  LEFT JOIN (
    SELECT
      et.id_errors,
      jsonb_agg (
        json_build_object (
          'id',
          t.id,
          'id_apps',
          t.id_apps,
          'name',
          t."name",
          'description',
          t.description,
          'color',
          t.color,
		    'background',
          t.background
        )
      ) AS tags
    FROM
      errors_tags et
      LEFT JOIN tags t ON t.id = et.id_tags
    GROUP BY
      et.id_errors
  ) t ON t.id_errors = e.id
	LEFT JOIN users u ON u.id = e.creator_id
WHERE
  e.id_apps = $1 ORDER BY e.id DESC`, appId)
	var errorsRecovereds []ErrorSaved
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tagsJSON []byte
			var errorSelected ErrorSaved
			err = rows.Scan(&errorSelected.Id, &errorSelected.Id_apps, &errorSelected.Message, &errorSelected.Title, &errorSelected.Verified,
				&errorSelected.Error_level, &errorSelected.Creator_id, &errorSelected.Created_in, &errorSelected.Last_edited_in,
				&errorSelected.How_to_reproduce, &errorSelected.Error_occurred_in, &tagsJSON, &errorSelected.Creator_name)
			if err == nil {
				if len(tagsJSON) > 0 {
					var tags []Tags.TagSaved
					if err := json.Unmarshal(tagsJSON, &tags); err != nil {
						return err, nil
					}
					errorSelected.Tags = &tags
				}
				errorsRecovereds = append(errorsRecovereds, errorSelected)
			}
		}
	}
	return err, errorsRecovereds
}
func UpdateErrors(errorToSave ErrorSaved) error {
	query := `
	UPDATE errors SET 
		id_apps = $1,
		message = $2,
		title = $3,
		verified = $4,
		error_level =$5,
		creator_id = $6,
		created_in =$7,
		last_edited_in =$8,
		how_to_reproduce =$9,
		error_occurred_in =$10
		WHERE id = $11

	`
	_, err := General.DB.Exec(
		query,
		errorToSave.Id_apps,
		errorToSave.Message,
		errorToSave.Title,
		errorToSave.Verified,
		errorToSave.Error_level,
		errorToSave.Creator_id,
		errorToSave.Created_in,
		errorToSave.Last_edited_in,
		errorToSave.How_to_reproduce,
		errorToSave.Error_occurred_in,
		errorToSave.Id,
	)
	return err
}

func DeleteError(errorId int) error {
	if errorId == 0 {
		return fmt.Errorf("errorId = 0, falha ao apagar")
	}
	_, err := General.DB.Exec("DELETE FROM errors WHERE id = $1", errorId)
	return err
}
