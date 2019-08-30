package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
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
}

func main() {

}
