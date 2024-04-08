package main

import (
	"lbank-connector-go/sve"
)

const (
	apiKey    = ""
	secretKey = ""
)

var client = sve.NewClient(apiKey, secretKey)

func TestCreateOrder() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"symbol":    "lbk_usdt",
		"type":      "buy",
		"price":     "0.01",
		"amount":    "1",
		"custom_id": "test",
	}
	client.NewSpotService().CreateOrder(data)
}

func main() {
	TestCreateOrder()
}