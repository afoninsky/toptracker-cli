package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const appName = "toptracker"
const configPath = "$HOME/.toptracker"

const secondsInHour = 60 * 60

func hours(seconds int) string {
	secondsLeft := seconds % secondsInHour
	return fmt.Sprintf("%d:%d", seconds/secondsInHour, secondsLeft*60/secondsInHour)
}

func main() {

	viper.SetConfigName(appName)
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Config read failure: %s", err))
	}

	info, err := getTopTrackerInfo(viper.GetString("email"), viper.GetString("password"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Worked total:\n\tper month - %s\n\tper day - %s\n", hours(info.SecondsPerMonth), hours(info.SecondsPerDay))
}
