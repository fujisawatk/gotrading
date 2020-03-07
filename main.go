package main

import (
	"gotrading/bitflyer"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.Apikey, config.Config.ApiSecret)
	// ticker, _ := apiClient.GetTicker("BTC_USD")
	// fmt.Println(ticker)
	// fmt.Println(ticker.GetMidPrice())
	// fmt.Println(ticker.DateTime())
	// fmt.Println(ticker.TruncateDateTime(time.Hour))
	tickerChannel := make(chan bitflyer.Ticker)
	apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannel)
}
