package tests

import (
	Users "Nautilus/app/crud/users"
	"testing"
)

var userInsertedId int

func TestAddUser(t *testing.T) {
	err, userInsertedId = Users.AddUser(*Users.NewUserSaved(1, "Teste", "teste 1", Users.WithRole("Tester"), Users.WithPermissionLevel(2)))
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateUser(t *testing.T) {
	TestAddUser(t)
	err = Users.UpdateUser(*Users.NewUserSaved(1, "Teste atualizado", "teste 2", Users.WithRole("Tester"), Users.WithPermissionLevel(2), Users.WithId(userInsertedId)))
}

func TestGetUniqueUser(t *testing.T) {
	TestAddUser(t)
	err, users := Users.GetUsers(userInsertedId)
	if err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("len != 1")
	}
}
func TestGetUsers(t *testing.T) {
	err, _ := Users.GetUsers(0)
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteUser(t *testing.T) {
	TestAddUser(t)
	err := Users.DeleteUser(userInsertedId)
	if err != nil {
		t.Error(err)
	}
}
