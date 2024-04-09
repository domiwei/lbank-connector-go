package main

var client = sve.NewWsClient("", "")

func TestKbar() {
	client.Debug = true
	//client.SetHost(sve.LbankApiHost)
	client.NewWsMarketService().Kbar("kbar", "usdt_btc")
}

func main() {
	TestKbar()
}