package main

import (
	"fmt"
	"log"
	"os"
)

const calendarURL = "http://xmlcalendar.ru/data/ru/%s/calendar.xml"

func downloadCalendar(year string) error {
	filename := fmt.Sprintf("%s.xml", year)
	if _, err := os.Stat(fmt.Sprintf("%s/%s", configPath, filename)); err == nil {
		// calendar exists - no need to download
		log.Println("exist")
		return nil
	} else if os.IsNotExist(err) {
		log.Println("not exist")
		return nil
	} else {
		log.Println("err")
		return err
	}
	// client := &http.Client{}
	// res, err := client.Get(fmt.Sprintf(calendarURL, "2019"))
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()
	// return json.NewDecoder(res.Body).Decode(jsonRes)
}

func currentHolidays() (map[string]bool, error) {
	result := map[string]bool{}
	// download actual calendar
	// if err := downloadCalendar(); err != nil {
	// 	return result, err
	// }
	return result, nil
}
