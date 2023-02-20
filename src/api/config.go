package api

import (
	"fmt"
	"github.com/SalomanYu/go-hh-parser-vacancies/src/models"
	"io"
	"net/http"
	"time"
)

var RelevantCurrencies  = []models.Currency{}

const TOKEN = "QQAVSIBVU4B0JCR296THKB22JP05A92H329U49TDD9CRIS8DT9BRPPT7M9OLQ6HD"

var headers = map[string]string{
	"User-Agent": "Mozilla/5.0 (iPad; CPU OS 12_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148",
	"Authorization": fmt.Sprintf("Bearer %s", TOKEN),
}

func init() {
	RelevantCurrencies = GetCurrencies()
}

func GetJson(url string) (json string, err error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Set("User-Agent", headers["User-Agent"])
	response, err := client.Do(req)
	if err != nil {
		return 
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	return string(data), nil
}

func convertSalaryToRUB(salary models.Salary) models.Salary {
	for _, cur := range RelevantCurrencies {
		if cur.Code == salary.Currency {
			salary.To = salary.To / cur.Rate
			salary.From = salary.From / cur.Rate
			salary.Currency = "RUR"
			return salary
		}
	}
	return salary
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}