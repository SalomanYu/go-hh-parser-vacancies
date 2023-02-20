package api

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/postgres"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func scrapeVacancy(url string, city_edwica int, id_profession int)  { 
	var vacancy models.Vacancy
	json, err := GetJson(url)
	if err != nil {
		logger.Log.Printf("Ошибка при подключении к странице %s.\nТекст ошибки: %s", err, url)
	}
	if checkManyRequestsError(json) {
		time.Sleep(60)
		fmt.Println("Слишком много запросов в секунду. Делаем паузу на 60 секунд")
		json, err = GetJson(url) 
		if err != nil {
			logger.Log.Printf("Ошибка при подключении к странице %s.\nТекст ошибки: %s", err, url)
		}
		if checkManyRequestsError(json) {
			logger.Log.Printf("Не смогли спарсить вакансию: %s\nТекст ошибки: %s", url, json)
			return
		} 
	}

	salary := getSalary(json)
	vacancy.CityId = city_edwica
	vacancy.SalaryFrom = salary.From
	vacancy.ProfessionId = id_profession
	vacancy.SalaryTo = salary.To
	vacancy.Id = int(gjson.Get(json, "id").Int())
	vacancy.Title = gjson.Get(json, "name").String()
	vacancy.Url = gjson.Get(json, "alternate_url").String()
	vacancy.Experience = gjson.Get(json, "experience.name").String()
	vacancy.DateUpdate = gjson.Get(json, "created_at").String()
	vacancy.Skills = getSkills(json)
	vacancy.Specializations = getSpecializations(json)
	vacancy.ProfAreas = getProfAreas(json)
	postgres.SaveOneVacancy(vacancy)
}

func checkManyRequestsError(json string) bool {
	err := gjson.Get(json ,"errors.0.type").String()
	if err == "" {
		return false
	} else {
		return true
	}
}

func getSalary(vacancyJson string) (salary models.Salary) {
	salary.Currency = gjson.Get(vacancyJson, "salary.currency").String()
	salary.From = gjson.Get(vacancyJson, "salary.from").Float()
	salary.To = gjson.Get(vacancyJson, "salary.to").Float()

	switch salary.Currency {
		case "RUR": return salary
		case "": return models.Salary{}
		default: return convertSalaryToRUB(salary)
	}

}

func getSpecializations(vacancyJson string) string {
	var specializations []string
	for _, item := range gjson.Get(vacancyJson, "specializations").Array() {
		specializations = append(specializations, item.Get("name").String())
	}
	return strings.Join(removeDuplicateStr(specializations), "|")
}

func getProfAreas(vacancyJson string) string {
	var profAreas []string
	for _, item := range gjson.Get(vacancyJson, "specializations").Array() {
		profAreas = append(profAreas, item.Get("profarea_name").String())
	}
	return strings.Join(removeDuplicateStr((profAreas)), "|")
}

func getSkills(vacancyJson string) string {
	var skills []string
	for _, item := range gjson.Get(vacancyJson, "key_skills").Array() {
		skills = append(skills, item.Get("name").String())
	}
	languages := getLanguages(vacancyJson)
	skills = append(skills, languages...)
	return strings.Join(skills, "|")
}

func getLanguages(vacancyJson string) (languages []string) {
	for _, item := range gjson.Get(vacancyJson, "languages").Array() {
		lang := item.Get("name").String()
		level := item.Get("level.name").String()
		languages = append(languages, fmt.Sprintf("%s (%s)", lang, level))
	}
	return
}

func removeDuplicateStr(strSlice []string) []string {
    allKeys := make(map[string]bool)
    list := []string{}
    for _, item := range strSlice {
        if _, value := allKeys[item]; !value {
            allKeys[item] = true
            list = append(list, item)
        }
    }
    return list
}