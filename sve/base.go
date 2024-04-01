package sve

type BaseService struct {
	c  *Client
	hs *HttpService
}

func (b *BaseService) CurrencyPairs(data map[string]string) {
	url := b.c.Host + PathCurrencyPairs
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Accuracy(data map[string]string) {
	url := b.c.Host + PathAccuracy
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) UsdToCny(data map[string]string) {
	url := b.c.Host + PathUsdToCny
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) WithdrawConfigs(data map[string]string) {
	url := b.c.Host + PathWithdrawConfigs
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Timestamp(data map[string]string) {
	url := b.c.Host + PathTimestamp
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Ticker24hr(data map[string]string) {
	url := b.c.Host + PathTicker24hr
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) EtfTicker24hr(data map[string]string) {
	url := b.c.Host + PathEtfTicker24hr
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Ticker(data map[string]string) {
	url := b.c.Host + PathTicker
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) IncrDepth(data map[string]string) {
	url := b.c.Host + PathIncrDepth
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Trades(data map[string]string) {
	url := b.c.Host + PathTrades
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}

func (b *BaseService) Kline(data map[string]string) {
	url := b.c.Host + PathKline
	params := b.hs.BuildSignBody(data)
	b.hs.Get(url, params)
}
