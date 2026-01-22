package Thoughts

import (
	General "Nautilus/general"
	"encoding/json"
	"fmt"
)

func AddTought(toughtToSave ThoughtSaved) (error, int) {
	var insertedId int
	err := General.DB.QueryRow("INSERT INTO thoughts (id_errors,creator_id,thought,created_in,vinculated_to,files) VALUES($1,$2,$3,$4,$5,$6) RETURNING id", toughtToSave.Id_errors, toughtToSave.Creator_id, toughtToSave.Thought, toughtToSave.Created_in, toughtToSave.Vinculated_to, General.JsonViewInterface(toughtToSave.Files)).Scan(&insertedId)
	return err, insertedId
}

func GetThought(thoughtId int) (error, []ThoughtSaved) {
	var thoughtsSaved = []ThoughtSaved{}
	filter := ""
	if thoughtId != 0 {
		filter = fmt.Sprintf("WHERE id = %d", thoughtId)
	}
	query := fmt.Sprintf("SELECT id,id_errors,creator_id,thought,created_in,vinculated_to,files FROM thoughts %s", filter)
	rows, err := General.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var thoughtSaved = ThoughtSaved{}
			var jsonFiles []byte
			err = rows.Scan(&thoughtSaved.Id, &thoughtSaved.Id_errors, &thoughtSaved.Creator_id, &thoughtSaved.Thought, &thoughtSaved.Created_in, &thoughtSaved.Vinculated_to, &jsonFiles)
			if err == nil {
				if len(jsonFiles) > 0 {
					var files []string
					if err := json.Unmarshal(jsonFiles, &files); err != nil {
						return err, nil
					}
					thoughtSaved.Files = &files
				}
				thoughtsSaved = append(thoughtsSaved, thoughtSaved)
			}
		}
	}
	return err, thoughtsSaved
}
func GetThoughtByErrorId(errorId int) (error, []ThoughtSaved) {
	var thoughtsSaved = []ThoughtSaved{}
	if errorId == 0 {
		return fmt.Errorf("errorid = 0"), thoughtsSaved
	}
	query := `SELECT t.id,t.id_errors,t.creator_id,t.thought,t.created_in,t.vinculated_to,t.files,u."name" as creator_name,u.profile_picture as creator_profile_picture  FROM thoughts t left join users u on u.id  = t.creator_id where  t.id_errors = $1 ORDER BY t.id asc`
	rows, err := General.DB.Query(query, errorId)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var thoughtSaved = ThoughtSaved{}
			var jsonFiles []byte
			err = rows.Scan(&thoughtSaved.Id, &thoughtSaved.Id_errors, &thoughtSaved.Creator_id, &thoughtSaved.Thought, &thoughtSaved.Created_in, &thoughtSaved.Vinculated_to, &jsonFiles, &thoughtSaved.Creator_name, &thoughtSaved.Creator_profile_picture)
			if err == nil {
				if len(jsonFiles) > 0 {
					var files []string
					if err := json.Unmarshal(jsonFiles, &files); err != nil {
						return err, nil
					}
					thoughtSaved.Files = &files
				}
				thoughtsSaved = append(thoughtsSaved, thoughtSaved)
			}
		}
	}
	return err, thoughtsSaved
}
func UpdateThought(toughtToSave ThoughtSaved) error {
	_, err := General.DB.Exec("UPDATE thoughts SET id_errors = $1,creator_id = $2,thought=$3 WHERE id = $4", toughtToSave.Id_errors, toughtToSave.Creator_id, toughtToSave.Thought, toughtToSave.Id)
	return err
}
func DeleteThought(thoughtId int) error {
	_, err := General.DB.Exec("DELETE FROM thoughts WHERE id = $1", thoughtId)
	return err
}
