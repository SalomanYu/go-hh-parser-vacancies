package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	err := godotenv.Load(".env")
	checkErr(err)
	host = os.Getenv("POSTGRES_HOST")
	port, _ = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	user = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname = os.Getenv("POSTGRES_DATABASE")
	fmt.Println(host)
}

var (
	host string
	port int
	user string
	password string
	dbname  string

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