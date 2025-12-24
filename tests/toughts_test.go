package tests

import (
	Thoughts "Nautilus/app/crud/thoughts"
	"fmt"
	"testing"
)

var insertedToughtId int

func TestAddTought(t *testing.T) {
	TestAddError(t)
	err, insertedToughtId = Thoughts.AddTought(*Thoughts.NewToughtToSave(errorInsertedId, 1, "teste"))
	if err != nil {
		t.Error(err)
	}
}
func TestGetUniqueThought(t *testing.T) {
	TestAddTought(t)
	err, thoughts := Thoughts.GetThought(insertedToughtId)
	if err != nil {
		t.Error(err)
	}
	if len(thoughts) != 1 {
		t.Error(fmt.Errorf("recuperou mais/menos de um app (%d)\n", len(thoughts)))
	}
}
func TestGetThoughts(t *testing.T) {
	err, _ := Thoughts.GetThought(0)
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateThoughts(t *testing.T) {
	TestAddTought(t)
	err := Thoughts.UpdateThought(*Thoughts.NewToughtToSave(errorInsertedId, 1, "teste atualizado", Thoughts.WithId(insertedToughtId)))
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteThoughts(t *testing.T) {
	TestAddTought(t)
	err := Thoughts.DeleteThought(insertedToughtId)
	if err != nil {
		t.Error(err)
	}
}
