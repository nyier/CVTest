package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var DB *sql.DB

func ConnectDatabase() {
	var err error
	dsn := "postgres://postgres:123@localhost:5432/CVTest" // Cambia esto según tu configuración
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Database connected successfully")
}
