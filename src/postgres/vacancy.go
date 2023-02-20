package postgres

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
	"strings"
	"time"
)

var ids []int

func init() {
	ids = GetVacanciesId()
}


func SaveVacancies(vacancies []models.Vacancy) {
	db := connect()
	defer db.Close()

	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueInsertCount := 1
	for _, vac := range vacancies {
		if vac.Title == ""{continue}
		if vac.DateUpdate == "" {vac.DateUpdate = time.Now().String()}
		if checkVacancyInList(vac.Id, valueArgs) || checkVacancyInDB(vac.Id){
			continue
		}
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)", valueInsertCount, valueInsertCount+1, valueInsertCount+2, valueInsertCount+3, valueInsertCount+4, valueInsertCount+5, valueInsertCount+6, valueInsertCount+7, valueInsertCount+8, valueInsertCount+9, valueInsertCount+10, valueInsertCount+11))
		valueInsertCount += 12
		

		valueArgs = append(valueArgs, vac.Id)
		valueArgs = append(valueArgs, vac.Url)
		valueArgs = append(valueArgs, vac.Title)
		valueArgs = append(valueArgs, vac.CityId)
		valueArgs = append(valueArgs, vac.ProfessionId)
		valueArgs = append(valueArgs, vac.ProfAreas)
		valueArgs = append(valueArgs, vac.Specializations)
		valueArgs = append(valueArgs, vac.Experience)
		valueArgs = append(valueArgs, vac.SalaryFrom)
		valueArgs = append(valueArgs, vac.SalaryTo)
		valueArgs = append(valueArgs, vac.Skills)
		valueArgs = append(valueArgs, vac.DateUpdate)
	}
	if len(valueArgs) == 0 {return}
	smt := fmt.Sprintf("INSERT INTO %s (hh_id, hh_url, name, city_id, position_id, hh_prof_areas, hh_specs, experience, salary_from, salary_to, key_skills, vacancy_date) VALUES ", TableVacancy)
	smt = fmt.Sprintf(smt + "%s", strings.Join(valueStrings, ","))
	tx, _ := db.Begin()
	_, err := db.Exec(smt, valueArgs...)
	if err != nil {
		logger.Log.Printf("ERROR SAVING VACANCIES: %s", err)
		return
	}
	tx.Commit()
	fmt.Println("Успешно добавили вакансии", len(vacancies))
}

func checkVacancyInList(vacancyId int, list []interface{}) bool {
	for _,item := range list {
		if item == vacancyId {
			return true
		}
	}
	return false
}

func GetVacanciesId() (ids []int)	{
	db := connect()
	defer db.Close()

	query := fmt.Sprintf("SELECT id FROM %s", TableVacancy)
	rows, err := db.Query(query)
	checkErr(err)
	for rows.Next(){
		var id int
		rows.Scan(&id)
		ids = append(ids, id)
	}
	return 
}

func checkVacancyInDB(vacancyId int) bool{
	for _, id := range ids {
		if id == vacancyId {
			logger.Log.Printf("Vacancy %d уже существует", id)
			return true
		}
	}
	return false
}