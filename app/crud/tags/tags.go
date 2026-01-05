package Tags

import (
	General "Nautilus/general"
	"fmt"
)

func AddTag(tagToSave TagSaved) (error, int) {
	var insertedId int
	var err error
	err = General.DB.QueryRow("INSERT INTO tags (id_apps,name,description,color) VALUES ($1,$2,$3,$4)", tagToSave.Id_apps, tagToSave.Name, tagToSave.Description, tagToSave.Color).Scan(&insertedId)
	return err, insertedId
}
func GetTags(tagId int) (error, []TagSaved) {
	tagsSaveds := []TagSaved{}
	var err error
	filter := ""
	if tagId != 0 {
		filter = fmt.Sprintf(" WHERE id = %d", tagId)
	}
	query := fmt.Sprintf("SELECT id,id_apps,name,description,color FROM tags %s", filter)
	rows, err := General.DB.Query(query)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			tagSaved := TagSaved{}
			err = rows.Scan(&tagSaved.Id, &tagSaved.Id_apps, &tagSaved.Name, &tagSaved.Description, &tagSaved.Color)
			if err == nil {
				tagsSaveds = append(tagsSaveds, tagSaved)
			}
		}
	}
	return err, tagsSaveds
}
func UpdateTag(tagToSave TagSaved) error {
	var err error
	fmt.Println(General.JsonViewInterface(tagToSave))
	_, err = General.DB.Exec("UPDATE tags SET id_apps = $1, name =$2,description =$3,color=$4 where id = $5", tagToSave.Id_apps, tagToSave.Name, tagToSave.Description, tagToSave.Color, tagToSave.Id)
	return err
}
func DeleteTag(tagId int) error {
	if tagId == 0 {
		return fmt.Errorf("tagId = 0")
	}
	var err error
	_, err = General.DB.Exec("DELETE FROM tags WHERE id = $1", tagId)
	return err
}
