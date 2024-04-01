package main

import (
	"lbank_connector_go/sve"
)

// SHA256
const (
	apiKey    = "44b9cdf2-6c66-4f57-a551-a80dbc42542d"
	secretKey = "71CE6CF6E03A51C61AA6F94A453443E8"
)

var client = sve.NewClient(apiKey, secretKey)

func TestAccuracy() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"symbol": "lbk_usdt",
		"size":   "1",
	}
	client.NewBaseService().Accuracy(data)
}

func main() {
	TestAccuracy()
}
