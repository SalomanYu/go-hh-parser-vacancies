package api

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/tidwall/gjson"

	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"

)

const (
	search_field = "name"
	per_page = "100"
)
const domain = "https://api.hh.ru/vacancies?"

func CreateLink(name string, area int) (link string) {
	params := url.Values{
		"search_field": {search_field},
		"per_page": {per_page},
		"text": {name},
		"area": {strconv.Itoa(area)},
	}
	link = domain + params.Encode()
	return
}

func GetVacanciesByQuery(city models.City, professionName string, professionId int) (vacancies []models.Vacancy) {
	url :=CreateLink(professionName, city.HH_ID)
	checkCaptcha(url)
	json, err := GetJson(url)
	if err != nil {
		logger.Log.Printf("Ошибка при подключении к странице с вакансиями: %s. Error: %s", err, url)
	}
	pagesCount := gjson.Get(json, "pages").Int()
	found := gjson.Get(json, "found").Int()
	logger.Log.Printf("Найдено %d вакансий для профессии %s в городе %s", found, professionName, city.Name)
	for page:=0; page<int(pagesCount); page++ {
		ParseVacanciesFromPage(fmt.Sprintf("%s&page=%d", url, page), city.EDWICA_ID, professionId)
	}
	return
}

func ParseVacanciesFromPage(url string, city_edwica int, id_profession int)  (vacancies []models.Vacancy){
	json, err := GetJson(url)
	if err != nil {
		logger.Log.Printf("Не удалось подключиться к странице %s.\nТекст ошибки: %s", err, url)
	}

	items := gjson.Get(json, "items").Array()
	for _, item := range items {
		scrapeVacancy(item.Get("url").String(), city_edwica,  id_profession)
	}
	return
}
