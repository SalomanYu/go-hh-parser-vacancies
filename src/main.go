package main

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/api"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/postgres"
	"sync"
	"time"

)

const GroupSize = 10

func main(){
	start := time.Now().Unix()
	logger.Log.Printf("Старт программы")
	UpdateCurrency()
	professions := postgres.GetProfessions()
	for _, prof := range professions {
			parseProfession(prof)
	}
	logger.Log.Println("Время выполнения программы в секундах:" ,time.Now().Unix()-start)
	fmt.Println("Время выполнения программы в секундах:" ,time.Now().Unix()-start)

}

func UpdateCurrency() {
	var confirmation string
	fmt.Printf("Обновить текущие значения валюты в БД? [Y/n] ")
	fmt.Scan(&confirmation)
	if confirmation != "Y" {
		logger.Log.Printf("Отмена обновления валюты в БД")
		return
	}

	currency := api.GetCurrencies()
	postgres.UpdateCurrencyRate(currency)
	logger.Log.Print("Валюта в БД была обновлена")
}

func parseProfession(profession models.Profession) {
	logger.Log.Printf("Ищем профессию %s", profession.Name)
	groups := groupCities()
	for _, group := range groups {
		var wg sync.WaitGroup
		wg.Add(len(group))
		for _, city := range group {
			go parseProfessionInCity(city, profession, &wg)
		}
		wg.Wait()
		break
	}
	postgres.SetParsedStatusToProfession(profession.Id)
	logger.Log.Printf("Профессия %s спарсена", profession.Name)

}

func groupCities() (groups [][]models.City) {
	cities := postgres.GetCities()
	citiesCount := len(cities)
	var limit int
	for i:=0; i<citiesCount; i+=GroupSize {
		limit += GroupSize
		if limit > citiesCount {limit = citiesCount}
		group := cities[i:limit]
		groups = append(groups, group)
	}
	logger.Log.Printf("Ведем поиск профессии в  %d городах одновременно", GroupSize)
	return
}

func parseProfessionInCity(city models.City, profession models.Profession, wg *sync.WaitGroup) {
	logger.Log.Printf("Ищем профессию %s в городе %s", profession.Name, city.Name)
	api.GetVacanciesByQuery(city, profession.Name, profession.Id)
	for _, other_name := range profession.OtherNames {
		logger.Log.Printf("Ищем профессию %s в городе %s", other_name, city.Name)
		api.GetVacanciesByQuery(city, other_name, profession.Id)
	}
	wg.Done()
}
