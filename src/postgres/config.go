package postgres

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "admin"
	dbname = "vacancies"

	TableCity = "city"
	TableProfessions = "position"
	TableCurrencies = "currencies"
	TableVacancy = "vacancy"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func connect() (db *sql.DB){
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	return
}