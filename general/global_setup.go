package General

import (
	Store "Nautilus/store"
	"database/sql"
	"log"
	"sync"
)

var DB *sql.DB
var DBMutex sync.Mutex

func Setup() {
	var err error
	DB, err = Store.ConnectPsql()
	if err != nil {
		log.Fatal("Erro ao conectar na base de dados :", err)
	}
}
