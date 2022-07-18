package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	url := "https://www.googleapis.com/calendar/v3/calendars/en.indonesian%23holiday%40group.v.calendar.google.com/events?key=" + apiKey
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = saveResponseAsFile(body)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveResponseAsFile(data []byte) error {
	f, err := os.Create("holidays.txt")
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(data)

	return err
}
