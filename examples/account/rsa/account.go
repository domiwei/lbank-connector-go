package main

import (
	"lbank_connector_go/sve"
)

// rsa
const (
	apiKey    = "1cc62cfb-2f36-4ac9-b5fb-2c40138db8ab"
	secretKey = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAILOdckXVqiMcLEdpockDZ1bLQc/XDlFxBSXMenZaok1PGljO6j2f45+PRH4X1tls78FHLBjCE4uizJb9kcd6uQTDCV9mMU4Tqyy9hflI1xt4K+u2Oxi+2z+NgrIOQXBazaZ+SJ1t1NMK5DhR7QUMPqMUg+JX7e2Xv89xOTiSfU1AgMBAAECgYB+OtcXs9oA1WZ6xW5Kw9QPokkV0WMiMd1DMZUNYq6YsjMWUJjmONpnnBM7IECFZuPK1xgUb704FVpmwrAreQeOpkS8i8PegP0yB/uaQAw1RYmnhOVpeJJPpHaLBwgSNP+EBBzi8/2ZYJPNbXaQm19QC0Y2grYMz35Z8Ro8zdF4gQJBAMvV5LRS1mSDpn6GG6APv/DjEblgaOAV4RWE1OJNxM3o8FeP41XPJMal6mX6YqwNVIvUFEv9ukiOiWmxrwrwRkkCQQCkSCINnWOQFRpsRkSIZ/ZgAz9PhdRcyIfqKCMstfaYu9SMubBD/rsJZSV27i+bmGwTR/Gmm4T51vadq/NwzUeNAkEAjnIYlKe7KZ0S8iJ4FcBL62RT0497WvYPSQF93/RnD1q08wwb27CZy7TQ/Jkg8YmTRvBbistyrhfmEZXZdLR6+QJBAIlKIvM/0cHKcQ+FVaatQy+P5yvdCtETYMpmCqdF1jRj3EhSsiTQz5wVZE7U1QJySfd/C0sR8vocFHNGDSb61s0CQQDLhsj7WLHTxZKiBvhxuXIwebQOVoqFomeAPloAMs2JUgIKqGXgVcByII2WiROmpWWx/W1ZfRxtHdD62v2BYBEI"
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
