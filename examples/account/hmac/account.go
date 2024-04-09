package main

import "github.com/LBank-exchange/lbank-connector-go/sve"

// rsa
const (
	apiKey    = ""
	secretKey = ""
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