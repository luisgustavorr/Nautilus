package tests

import (
	General "Nautilus/general"
	Store "Nautilus/store"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var err error

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("failed to load .env:", err)
	}
	General.DB, err = Store.ConnectPsql()
	if err != nil {
		log.Fatal(err)
	}
	code := m.Run()
	os.Exit(code)
}
