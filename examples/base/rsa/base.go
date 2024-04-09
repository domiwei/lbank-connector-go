package main

const (
	apiKey    = ""
	secretKey = ""
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