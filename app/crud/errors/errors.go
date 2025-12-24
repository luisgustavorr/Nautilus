package Errors

import (
	General "Nautilus/general"
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
