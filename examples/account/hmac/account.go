package main

import (
	"lbank_connector_go/sve"
)

// rsa
const (
	apiKey    = "44b9cdf2-6c66-4f57-a551-a80dbc42542d"
	secretKey = "71CE6CF6E03A51C61AA6F94A453443E8"
)

var client = sve.NewClient(apiKey, secretKey)

func TestUserInfo() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{}
	client.NewAccountService().UserInfo(data)
}

func TestSubscribeGetKey() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{}
	client.NewAccountService().SubscribeGetKey(data)
}

func TestSubscribeRefreshKey() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"subscribeKey": "32a6ece19c591f1791cc07ba570db4dd00a4ab7bb32f0f3adc7969b4c9e2e2f0",
	}
	client.NewAccountService().SubscribeRefreshKey(data)
}

func TestSubscribeDestroyKey() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"subscribeKey": "32a6ece19c591f1791cc07ba570db4dd00a4ab7bb32f0f3adc7969b4c9e2e2f0",
	}
	client.NewAccountService().SubscribeDestroyKey(data)
}

func TestGetDepositAddress() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"assetCode": "BTC",
		"netWork":   "",
	}
	client.NewAccountService().GetDepositAddress(data)
}

func TestDepositHistory() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"assetCode": "",
		"startTime": "",
		"endTime":   "",
		"pageNo":    "1",
		"pageSize":  "10",
	}
	client.NewAccountService().DepositHistory(data)
}

func main() {
	TestUserInfo()
	//TestSubscribeGetKey()
	//TestSubscribeRefreshKey()
	//TestSubscribeDestroyKey()
	//TestGetDepositAddress()
	//TestDepositHistory()
}
