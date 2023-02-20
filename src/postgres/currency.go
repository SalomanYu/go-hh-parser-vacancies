package postgres

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
)

func UpdateCurrencyRate(currencies []models.Currency) {
	db := connect()
	defer db.Close()

	for _, cur := range currencies {
		query := fmt.Sprintf(`UPDATE "%s" SET rate=%f WHERE code='%s'; UPDATE "%s" SET date_update=current_timestamp WHERE code='%s'`, 
			TableCurrencies, cur.Rate, cur.Code, TableCurrencies, cur.Code)
		fmt.Println(query)
		tx, _ := db.Begin()
		_, err := db.Exec(query)
		if err == nil {
			tx.Commit()
		} else {
			logger.Log.Printf("Не удалось обновить валюты - %s", err)
		}
	}
}