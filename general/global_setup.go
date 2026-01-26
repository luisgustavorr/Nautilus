package General

import (
	Store "Nautilus/store"
	"log"

	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	var err error
	DB, err = Store.ConnectPsql()
	if err != nil {
		log.Fatal("Erro ao conectar na base de dados :", err)
	}
}
