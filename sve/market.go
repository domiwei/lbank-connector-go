package sve

type MarketService struct {
	c  *Client
	hs *HttpService
}

func (m *MarketService) HttpService() *HttpService {
	return m.hs
}

func (m *MarketService) Depth(data map[string]string) {
	url := m.c.Host + PathDepth
	params := m.hs.BuildSignBody(data)
	m.hs.Get(url, params)
}
func (m *MarketService) SupplementTrades(data map[string]string) {
	url := m.c.Host + PathSupplementTrades
	params := m.hs.BuildSignBody(data)
	m.hs.Get(url, params)
}
func (m *MarketService) SupplementTickerPrice(data map[string]string) {
	url := m.c.Host + PathSupplementTickerPrice
	params := m.hs.BuildSignBody(data)
	m.hs.Get(url, params)
}
func (m *MarketService) SupplementTickerBookTicker(data map[string]string) {
	url := m.c.Host + PathSupplementTickerBookTicker
	params := m.hs.BuildSignBody(data)
	m.hs.Get(url, params)
}
