package postgres

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
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

	// TableCity = "city"
	// TableProfessions = "position"
	// TableCurrencies = "currencies"
	// TableVacancy = "vacancy"

	TableCity = "h_city"
	TableProfessions = "h_position"
	TableCurrencies = "h_currency"
	TableVacancy = "h_vacancy"

)

func checkErr(err error) {
	if err != nil {
		logger.Log.Fatal(err)
		panic(err)
	}
}

func connect() (db *sql.DB){
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", conn)
	checkErr(err)
	return
}

func Connect2() (db *sql.DB){
	db, err := sql.Open("mysql", "edwica_root:b00m5gQ40WB1@tcp(83.220.175.75:3306)/edwica")
	checkErr(err)
	db.Ping()
	return db
}