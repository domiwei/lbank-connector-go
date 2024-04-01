package sve

type WalletService struct {
	c  *Client
	hs *HttpService
}

func (w *WalletService) SupplementSystemStatus(data map[string]string) {
	url := w.c.Host + PathSupplementSystemStatus
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementUserInfo(data map[string]string) {
	url := w.c.Host + PathSupplementUserInfo
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementWithdraw(data map[string]string) {
	url := w.c.Host + PathSupplementWithdraw
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementDepositHistory(data map[string]string) {
	url := w.c.Host + PathSupplementDepositHistory
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementWithdraws(data map[string]string) {
	url := w.c.Host + PathSupplementWithdraws
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementGetDepositAddress(data map[string]string) {
	url := w.c.Host + PathSupplementGetDepositAddress
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementAssetDetail(data map[string]string) {
	url := w.c.Host + PathSupplementAssetDetail
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementCustomerTradeFee(data map[string]string) {
	url := w.c.Host + PathSupplementCustomerTradeFee
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}

func (w *WalletService) SupplementApiRestrictions(data map[string]string) {
	url := w.c.Host + PathSupplementApiRestrictions
	params := w.hs.BuildSignBody(data)
	w.hs.Post(url, params)
}
