package tests

import (
	Apps "Nautilus/app/crud/apps"
	"fmt"
	"testing"
)

var appInsertedId int

func TestAddApp(t *testing.T) {
	var err error
	err, appInsertedId = Apps.AddApp(*Apps.NewAppSaved("Teste", "teste 1", "teste 2"))
	if err != nil {
		t.Error(err)
	}
}

func TestGetUniqueApp(t *testing.T) {
	err, apps := Apps.GetApp(1)
	if err != nil {
		t.Error(err)
	}
	if len(apps) != 1 {
		t.Error(fmt.Errorf("recuperou mais/menos de um app (%d)\n", len(apps)))
	}
}
func TestGetApps(t *testing.T) {
	err, _ := Apps.GetApp(0)
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateApp(t *testing.T) {
	TestAddApp(t)
	err := Apps.UpdateApp(*Apps.NewAppSaved("Teste Atualizado", "teste 1", "teste 2", Apps.WithId(appInsertedId)))
	if err != nil {
		t.Error(err)
	}
}
func TestDeleteApp(t *testing.T) {
	err := Apps.DeleteApp(appInsertedId)
	if err != nil {
		t.Error(err)
	}
}
