package General

import (
	Store "Nautilus/store"
	"database/sql"
	"log"
)

var DB *sql.DB

func Setup() {
	var err error
	DB, err = Store.ConnectPsql()
	if err != nil {
		log.Fatal("Erro ao conectar na base de dados :", err)
	}
}
