package main

import (
	"calendarProj/internal/events"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"time"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/calendar/")
	viper.AddConfigPath("$HOME/.calendar")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

func main() {
	jsonFile, err := os.Open("configZap.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var cfg zap.Config
	if err := json.Unmarshal(byteValue, &cfg); err != nil {
		panic(err)
	}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("logger construction succeeded")
	usr := events.AddUser("yo", "123")
	events.AddEvent(&events.Event{
		Id:                   0,
		Title:                "Event 1",
		Date:                 time.Time{},
		Deadline:             time.Time{}.AddDate(0, 0, 2),
		Description:          "wow this is event 1",
		User:                 usr,
		NotificationDeadline: time.Time{},
	})
	logger.Info(fmt.Sprintf("%#v", events.GetEventById(0)))
	logger.Info(fmt.Sprintf("%#v", events.GetEventsByDate(time.Time{})))
	logger.Info(fmt.Sprintf("%#v", events.GetEventsWeek(time.Time{})))
	logger.Info(fmt.Sprintf("%#v", events.GetEventsMonth(time.Time{})))
}
