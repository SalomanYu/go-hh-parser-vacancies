package postgres

import (
	"fmt"
	"strings"

	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
)


func SetParsedStatusToProfession(id int) {
	db := Connect2()
	defer db.Close()

	query := fmt.Sprintf(`UPDATE "%s" SET parsed = true WHERE id = %d`, TableProfessions, id)
	fmt.Println(query)
	tx, _ := db.Begin()
	_, err := db.Exec(query)
	checkErr(err)
	tx.Commit()
} 


func GetProfessions() (professions []models.Profession) {
	db := Connect2()
	defer db.Close()

	query := fmt.Sprintf("SELECT id, name, other_names, level, parent_id FROM %s WHERE parsed = false", TableProfessions)
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next() {
		var (
			name, other string
			id, level, parent_id int
		)
		err = rows.Scan(&id, &name, &other, &level, &parent_id, )
		checkErr(err)
		prof := models.Profession{
			Id: id,
			Name: name,
			OtherNames: strings.Split(other, "|"),
			Level: level,
			ParentId: parent_id,
		}
		professions = append(professions, prof)

	}
	return
}
