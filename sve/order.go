package sve

type OrderService struct {
	c  *Client
	hs *HttpService
}

func (o *OrderService) HttpService() *HttpService {
	return o.hs
}

func (o *OrderService) CreateOrder(data map[string]string) {
	url := o.c.Host + PathCreateOrder
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) BatchCreateOrder(data map[string]string) {
	url := o.c.Host + PathBatchCreateOrder
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) CancelOrder(data map[string]string) {
	url := o.c.Host + PathCancelOrder
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) CancelClientOrders(data map[string]string) {
	url := o.c.Host + PathCancelClientOrders
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) OrdersInfo(data map[string]string) {
	url := o.c.Host + PathOrdersInfo
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) OrdersInfoHistory(data map[string]string) {
	url := o.c.Host + PathOrdersInfoHistory
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) OrderTransactionDetail(data map[string]string) {
	url := o.c.Host + PathOrderTransactionDetail
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) TransactionHistory(data map[string]string) {
	url := o.c.Host + PathTransactionHistory
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}

func (o *OrderService) OrdersInfoNoDeal(data map[string]string) {
	url := o.c.Host + PathOrdersInfoNoDeal
	params := o.hs.BuildSignBody(data)
	o.hs.Post(url, params)
}
