package main

import (
	"lbank_connector_go/sve"
)

const (
	apiKey    = "1cc62cfb-2f36-4ac9-b5fb-2c40138db8ab"
	secretKey = "MIICeAIBADANBgkqhkiG9w0BAQEFAASCAmIwggJeAgEAAoGBAILOdckXVqiMcLEdpockDZ1bLQc/XDlFxBSXMenZaok1PGljO6j2f45+PRH4X1tls78FHLBjCE4uizJb9kcd6uQTDCV9mMU4Tqyy9hflI1xt4K+u2Oxi+2z+NgrIOQXBazaZ+SJ1t1NMK5DhR7QUMPqMUg+JX7e2Xv89xOTiSfU1AgMBAAECgYB+OtcXs9oA1WZ6xW5Kw9QPokkV0WMiMd1DMZUNYq6YsjMWUJjmONpnnBM7IECFZuPK1xgUb704FVpmwrAreQeOpkS8i8PegP0yB/uaQAw1RYmnhOVpeJJPpHaLBwgSNP+EBBzi8/2ZYJPNbXaQm19QC0Y2grYMz35Z8Ro8zdF4gQJBAMvV5LRS1mSDpn6GG6APv/DjEblgaOAV4RWE1OJNxM3o8FeP41XPJMal6mX6YqwNVIvUFEv9ukiOiWmxrwrwRkkCQQCkSCINnWOQFRpsRkSIZ/ZgAz9PhdRcyIfqKCMstfaYu9SMubBD/rsJZSV27i+bmGwTR/Gmm4T51vadq/NwzUeNAkEAjnIYlKe7KZ0S8iJ4FcBL62RT0497WvYPSQF93/RnD1q08wwb27CZy7TQ/Jkg8YmTRvBbistyrhfmEZXZdLR6+QJBAIlKIvM/0cHKcQ+FVaatQy+P5yvdCtETYMpmCqdF1jRj3EhSsiTQz5wVZE7U1QJySfd/C0sR8vocFHNGDSb61s0CQQDLhsj7WLHTxZKiBvhxuXIwebQOVoqFomeAPloAMs2JUgIKqGXgVcByII2WiROmpWWx/W1ZfRxtHdD62v2BYBEI"
)

var client = sve.NewClient(apiKey, secretKey)

func TestDepth() {
	client.Debug = true
	client.SetHost(sve.LbankApiHost)
	data := map[string]string{
		"symbol": "lbk_usdt",
		"size":   "1",
	}
	client.NewMarketService().Depth(data)
}

func main() {
	TestDepth()
}
