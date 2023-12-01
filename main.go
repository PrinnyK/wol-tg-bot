package main

import (
	"github.com/spf13/viper"
	"log"
	"wol-tg-bot/bot"
)

func main() {
	// read config
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("read config.json failed cannot start")
		return
	}

	// required
	token := viper.GetString("botToken")
	validUserNameList := viper.GetStringSlice("validUserNameList")
	macAddr := viper.GetString("targetMacAddr")
	if token == "" || len(validUserNameList) == 0 || macAddr == "" {
		log.Println("please set botToken, validUserNameList, and targetMacAddr in config.json")
		return
	}

	// optional
	ipAddr := viper.GetString("targetIpAddr")
	port := viper.GetInt("targetPort")
	if ipAddr == "" {
		ipAddr = "255.255.255.255"
	}
	if port == 0 {
		port = 9
	}

	// start bot
	config := bot.WolBotConfig{
		Token:             token,
		ValidUserNameList: validUserNameList,
		MacAddr:           macAddr,
		IpAddr:            ipAddr,
		Port:              port,
	}
	_ = bot.Start(&config)
}
