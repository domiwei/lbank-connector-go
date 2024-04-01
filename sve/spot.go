package sve

type SpotService struct {
	c  *Client
	hs *HttpService
}

func (s *SpotService) CreateOrder(data map[string]string) {
	url := s.c.Host + PathSupplementCreatOrder
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) CancelOrder(data map[string]string) {
	url := s.c.Host + PathSupplementCancelOrder
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) CancelOrderBySymbol(data map[string]string) {
	url := s.c.Host + PathSupplementCancelOrderBySymbol
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) OrdersInfo(data map[string]string) {
	url := s.c.Host + PathSupplementOrdersInfo
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) OrdersInfoNoDeal(data map[string]string) {
	url := s.c.Host + PathSupplementOrdersInfoNoDeal
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) OrdersInfoHistory(data map[string]string) {
	url := s.c.Host + PathSupplementOrdersInfoHistory
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}

func (s *SpotService) UserInfoAccount(data map[string]string) {
	url := s.c.Host + PathSupplementUserInfoAccount
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}
func (s *SpotService) TransactionHistory(data map[string]string) {
	url := s.c.Host + PathSupplementTransactionHistory
	params := s.hs.BuildSignBody(data)
	s.hs.Post(url, params)
}
