package main

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

const appName = "toptracker"
const configPath = "$HOME/.toptracker"

const secondsInHour = 60 * 60
const dailyWorkOutSeconds = 6.5 * secondsInHour

func hours(seconds int) string {
	secondsLeft := seconds % secondsInHour
	return fmt.Sprintf("%d:%02d", seconds/secondsInHour, secondsLeft*60/secondsInHour)
}

func main() {
	// get dates
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	prevOfMonth := now.AddDate(0, 0, -1)

	// check if current year is supported
	currentYear := now.Format("2006")
	_, exists := supportedYears[currentYear]
	if !exists {
		log.Fatal("current year is not supported, please update working calendar")
	}

	// generate config on fist run
	// 2DO

	viper.SetConfigName(appName)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Config read failure: %s", err)
	}

	// fetch information from the trecker
	topTrackerInfo, err := getTopTrackerInfo(viper.GetString("email"), viper.GetString("password"))
	if err != nil {
		log.Fatal(err.Error())
	}
	daysInMonth := workingDaysInRange(firstOfMonth, lastOfMonth)
	workSecondsInMonth := daysInMonth * dailyWorkOutSeconds

	dailyEffiency := effiency(dailyWorkOutSeconds, topTrackerInfo.SecondsPerDay)

	// montly effiency = how much I did in previous days
	secondsDonePrev := topTrackerInfo.SecondsPerMonth - topTrackerInfo.SecondsPerDay
	workingDaysPrev := workingDaysInRange(firstOfMonth, prevOfMonth)
	secondsExpectPrev := workingDaysPrev * dailyWorkOutSeconds
	monthlyEffiency := effiency(secondsExpectPrev, secondsDonePrev)

	fmt.Printf("\tNow %d of %d working days\n", workingDaysInRange(firstOfMonth, now), daysInMonth)
	fmt.Printf("\tMonthly stats: %sh / %sh (%d%%)", hours(topTrackerInfo.SecondsPerMonth), hours(workSecondsInMonth), monthlyEffiency)
	if secondsDonePrev > secondsExpectPrev {
		fmt.Printf("\t%sh+", hours(secondsDonePrev-secondsExpectPrev))
	} else {
		fmt.Printf("\t%sh-", hours(secondsDonePrev-secondsExpectPrev))
	}
	fmt.Printf("\n")
	fmt.Printf("\tDaily stats: %sh / %sh (%d%%)\n", hours(topTrackerInfo.SecondsPerDay), hours(dailyWorkOutSeconds), dailyEffiency)
}
