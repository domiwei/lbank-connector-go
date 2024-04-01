package sve

type WithDrawService struct {
	c  *Client
	hs *HttpService
}

func (w *WithDrawService) Withdraw(data map[string]string) {
	url := w.c.Host + PathWithdraw
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WithDrawService) WithdrawCancel(data map[string]string) {
	url := w.c.Host + PathWithdrawCancel
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WithDrawService) Withdraws(data map[string]string) {
	url := w.c.Host + PathWithdraws
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}
