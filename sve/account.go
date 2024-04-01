package sve

type AccountService struct {
	c  *Client
	hs *HttpService
}

func (a *AccountService) UserInfo(data map[string]string) {
	url := a.c.Host + PathUserInfo
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) SubscribeGetKey(data map[string]string) {
	url := a.c.Host + PathSubscribeGetKey
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) SubscribeRefreshKey(data map[string]string) {
	url := a.c.Host + PathSubscribeRefreshKey
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) SubscribeDestroyKey(data map[string]string) {
	url := a.c.Host + PathSubscribeDestroyKey
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) GetDepositAddress(data map[string]string) {
	url := a.c.Host + PathGetDepositAddress
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}

func (a *AccountService) DepositHistory(data map[string]string) {
	url := a.c.Host + PathDepositHistory
	params := a.hs.BuildSignBody(data)
	a.hs.Post(url, params)
}
