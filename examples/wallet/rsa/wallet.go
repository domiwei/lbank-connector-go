package main

const (
	apiKey    = ""
	secretKey = ""
)

var client = sve.NewClient(apiKey, secretKey)

func TestSupplementApiRestrictions() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"symbol": "lbk_usdt",
		"size":   "1",
	}
	client.NewWalletService().SupplementApiRestrictions(data)
}

func main() {
	TestSupplementApiRestrictions()
}