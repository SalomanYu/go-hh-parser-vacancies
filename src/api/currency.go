package api

import (
	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"

	"github.com/tidwall/gjson"
)

const dictionariesUrl = "https://api.hh.ru/dictionaries"

func GetCurrencies() (currencies []models.Currency) {
	json, err := GetJson(dictionariesUrl)
	if err != nil {
		logger.Log.Printf("Не удалось обновить валюту. Текст сообщения: %s", err)
	}
	for _, item := range gjson.Get(json, "currency").Array() {
		currencies = append(currencies, models.Currency{
			Code: item.Get("code").String(),
			Abbr: item.Get("abbr").String(),	
			Name: item.Get("name").String(),
			Rate: item.Get("rate").Float(),
		})		
	}
	return
}