package postgres

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
)

func GetCities() (cities []models.City){
	db := connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT * FROM %s WHERE id_hh != 0", TableCity)
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		var name string
		var hh_id, edwica_id int
		err = rows.Scan(&hh_id, &edwica_id, &name)
		checkErr(err)
		cities = append(cities, models.City{
			HH_ID: hh_id,
			EDWICA_ID: edwica_id,
			Name: name,
		})
	}
	return
}