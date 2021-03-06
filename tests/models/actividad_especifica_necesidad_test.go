package models

import (
	"testing"

	"github.com/udistrital/necesidades_crud/models"
	"github.com/udistrital/necesidades_crud/tests/testshelpers"
)

func TestCrudOpeartionActividadEspecificaNecesidad(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {

			t.Error("Error", r)
			t.Fail()
		} else {
			t.Log("OK")
		}
	}()
	testInsertObj := models.ActividadEspecificaNecesidad{
		Descripcion: "testTipoNecesidad",
		Activo:      true,
	}
	testUpdateObj := models.ActividadEspecificaNecesidad{
		Descripcion: "updateTipoNecesidad",
		Activo:      false,
	}
	if err := testshelpers.TestCrudOperations(&testInsertObj, &testUpdateObj); err != nil {
		panic(err.Error())
	}
}
