package tests

import (
	Tags "Nautilus/app/crud/tags"
	"testing"
)

var tagsInsertedId int

func TestAddTag(t *testing.T) {
	err, tagsInsertedId = Tags.AddTag(*Tags.NewTagToSave(1, "ERROR", "QUando erro"))
	if err != nil {
		t.Error(err)
	}
}
func TestGetUniqueTag(t *testing.T) {
	err, tags := Tags.GetTags(2)
	if err != nil {
		t.Error(err)
	}
	if len(tags) != 1 {
		t.Errorf("erro, more than one tag returned")
	}
}
func TestGetTags(t *testing.T) {
	err, _ := Tags.GetTags(0)
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateTag(t *testing.T) {
	err := Tags.UpdateTag(*Tags.NewTagToSave(1, "ERROR tu", "QUando erro TU", Tags.WithId(2), Tags.WithColor("blue")))
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteTag(t *testing.T) {
	err := Tags.DeleteTag(2)
	if err != nil {
		t.Error(err)
	}
}
