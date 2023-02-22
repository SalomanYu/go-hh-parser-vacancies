package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"io"

	"github.com/SalomanYu/go-hh-parser-vacancies/src/logger"
	"github.com/tidwall/gjson"
)

const token = "6105028983:AAG5qYpp0KKkhBiyOHri6MmB5e_UVzfC9pU"
var chats = []string{
	"544490770",
}

func Mailing(text string) {
	for _, chat := range chats {
		SendMessage(text, chat)
	}
}

func SendMessage(text string, chatId string) (bool, error) {
	url := fmt.Sprintf("%s/sendMessage", getUrl())
	body, _ := json.Marshal(map[string]string{
		"chat_id": chatId,
		"text": text,
	})
	response, err := http.Post(
		url,
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		return false, err
	}
	defer response.Body.Close()
	return true, nil
}

func getUrl() string {
	return fmt.Sprintf("https://api.telegram.org/bot%s", token)
}

func GetUpdates() (countMessages int) {
	json, _ := getJson(fmt.Sprintf("%s/getUpdates", getUrl()))
	return len(gjson.Get(json, "result").Array())
}

func getJson(url string) (json string, err error) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
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

func checkErr(err error) {
	if err != nil {
		logger.Log.Fatal(err)
		panic(err)
	}
}