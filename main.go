package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/tegarap/indonesian-holiday-google-calendar/router"
)

func main() {
	r := router.NewRouter()

	r.Methods(http.MethodGet).Handler(`/holiday/indo`, holidayHandler())

	port := "9898"
	fmt.Println("Running server at port:", port)
	http.ListenAndServe(":"+port, r)
}

func holidayHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// baseurl?apiKey=xxxxxxx&startDate=yyyy-mm-dd&endDate=yyyy-mm-dd&saveAsJson=true
		apiKey := r.URL.Query().Get("apiKey")
		startDate := r.URL.Query().Get("startDate")
		endDate := r.URL.Query().Get("endDate")
		saveAsJson := r.URL.Query().Get("saveAsJson")

		result, err := getHoliday(apiKey, startDate, endDate)
		if err != nil {
			fmt.Println("| 500 |", result)
		}

		var holidays []Holiday
		switch v := result.(type) {
		case ResponseSuccess:
			// Looping on result items
			for _, item := range v.Items {
				holiday := Holiday{
					Title: item.Summary,
					Date:  item.Start.Date,
				}
				holidays = append(holidays, holiday)
			}
		}

		// Print on the browser as html
		err = PrintAsHtml(w, holidays)
		if err != nil {
			fmt.Println("| 500 |", err)
		}

		if saveAsJson == "true" {
			// Save as json file
			err = SaveAsJsonFile(holidays)
			if err != nil {
				fmt.Println("| 500 |", err)
			}
		}

		if err == nil {
			fmt.Println("| 200 | Success")
		}
	})
}

func getHoliday(apiKey, timeMin, timeMax string) (interface{}, error) {
	// timeMin & timeMax must be an RFC3339 timestamp with mandatory time zone offset,
	// for example, 2011-06-03T00:00:00-07:00, 2011-06-03T00:00:00Z.
	if timeMin != "" {
		timeMin = "&timeMin=" + timeMin + "T00:00:00Z"
	}
	if timeMax != "" {
		timeMax = "&timeMin=" + timeMax + "T00:00:00Z"
	}

	method := "GET"
	url := "https://www.googleapis.com/calendar/v3/calendars/en.indonesian%23holiday%40group.v.calendar.google.com/events?key=" + apiKey + timeMin + timeMax

	resultBody, err := httpClient(method, url)
	if err != nil {
		rErr := ResponseError{}
		json.Unmarshal(resultBody, &rErr)
		return rErr, err
	}

	// Put resultbody to struct
	rSuccess := ResponseSuccess{}
	err = json.Unmarshal(resultBody, &rSuccess)

	return rSuccess, err
}

func httpClient(method, url string) (resBody []byte, err error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

	res, err := client.Do(req)
	if err != nil {
		return
	}

	defer res.Body.Close()
	resBody, err = ioutil.ReadAll(res.Body)

	return
}

func SaveAsJsonFile(data []Holiday) error {
	jData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	f, err := os.Create("holidays.json")
	if err != nil {
		return err
	}

	defer f.Close()
	_, err = f.Write(jData)

	return err
}

func PrintAsHtml(w http.ResponseWriter, data interface{}) error {
	templ := `{{define "T"}}<html><ul>{{range .}}<li>{{.}}</li>{{end}}</ul></html>{{end}}`

	t, err := template.New("Indonesia holidays").Parse(templ)
	if err != nil {
		return err
	}

	err = t.ExecuteTemplate(w, "T", data)

	return err
}
