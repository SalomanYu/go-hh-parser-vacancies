package postgres

import (
	"fmt"
	"strings"

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

func InsertNew(currencies []models.Currency) {
	db := connect()
	defer db.Close()

	valueStrings := []string{}
	valueArgs := []interface{}{}
	valueInsertCount := 1

	for _, cur := range currencies{
		valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", valueInsertCount, valueInsertCount+1, valueInsertCount+2, valueInsertCount+3))
		valueInsertCount += 4
		valueArgs = append(valueArgs, cur.Code)
		valueArgs = append(valueArgs, cur.Abbr)
		valueArgs = append(valueArgs, cur.Name)
		valueArgs = append(valueArgs, cur.Rate)
	}
	smt := fmt.Sprintf("INSERT INTO %s (code, abbr, name, rate) VALUES", TableCurrencies)
	smt = fmt.Sprintf(smt + "%s", strings.Join(valueStrings, ","))
	tx, _ := db.Begin()
	_, err := db.Exec(smt, valueArgs...)
	if err != nil {
		logger.Log.Printf("Ошибка: Не удалось добавить валюту в базу - %s", err)
		return
	}
	tx.Commit()
	fmt.Println("Успешно добавили валюту в базу!")
}