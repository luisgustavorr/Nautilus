package Thoughts

import (
	General "Nautilus/general"
	"fmt"
)

func AddTought(toughtToSave ThoughtSaved) (error, int) {
	var insertedId int
	err := General.DB.QueryRow("INSERT INTO thoughts (id_errors,creator_id,thought) VALUES($1,$2,$3) RETURNING id", toughtToSave.Id_errors, toughtToSave.Creator_id, toughtToSave.Thought).Scan(&insertedId)
	return err, insertedId
}

func GetThought(thoughtId int) (error, []ThoughtSaved) {
	var thoughtsSaved = []ThoughtSaved{}
	filter := ""
	if thoughtId != 0 {
		filter = fmt.Sprintf("WHERE id = %d", thoughtId)
	}
	query := fmt.Sprintf("SELECT id,id_errors,creator_id,thought FROM thoughts %s", filter)
	rows, err := General.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var thoughtSaved = ThoughtSaved{}
			err = rows.Scan(&thoughtSaved.Id, &thoughtSaved.Id_errors, &thoughtSaved.Creator_id, &thoughtSaved.Thought)
			if err == nil {
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
