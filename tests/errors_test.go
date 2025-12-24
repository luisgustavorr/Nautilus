package tests

import (
	Errors "Nautilus/app/crud/errors"
	"fmt"
	"testing"
)

var errorInsertedId int

func TestAddError(t *testing.T) {
	var err error
	err, errorInsertedId = Errors.AddError(*Errors.NewErrorToSave(1, 1, "Teste", "Erro de teste"))
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateError(t *testing.T) {
	TestAddError(t)
	var err error
	err = Errors.UpdateErrors(*Errors.NewErrorToSave(1, 1, "Teste atualizado", "Erro de teste", Errors.WithId(errorInsertedId)))
	if err != nil {
		t.Error(err)
	}
}
func TestGetError(t *testing.T) {
	err, errorsRecovereds := Errors.GetErrors(1)
	if err != nil {
		t.Error(err)
	}
	if len(errorsRecovereds) > 1 {
		t.Error(fmt.Errorf("More than one error recovered when getting ID"))
		return
	}
	t.Logf("%d error recovered \n", len(errorsRecovereds))

}
func TestGetMultipleErrors(t *testing.T) {
	err, errorsRecovereds := Errors.GetErrors(0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%d errors selected \n", len(errorsRecovereds))
}
func TestDeleteError(t *testing.T) {
	TestAddError(t)
	err := Errors.DeleteError(errorInsertedId)
	if err != nil {
		t.Error(err)
	}
}
