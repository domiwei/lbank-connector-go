package sve

const (
	//LbankApiHost   = "https://api.lbkex.com"
	LbankApiHost   = "https://www.lbkex.net"
	LbankWsApiHost = ""
)

const (
	PathSupplementCreatOrder          = "/v2/supplement/create_order.do"
	PathSupplementCancelOrder         = "/v2/supplement/cancel_order.do"
	PathSupplementCancelOrderBySymbol = "/v2/supplement/cancel_order_by_symbol.do"
	PathSupplementOrdersInfo          = "/v2/supplement/orders_info.do"
	PathSupplementOrdersInfoNoDeal    = "/v2/supplement/orders_info_no_deal.do"
	PathSupplementOrdersInfoHistory   = "/v2/supplement/orders_info_history.do"
	PathSupplementUserInfoAccount     = "/v2/supplement/user_info_account.do"
	PathSupplementTransactionHistory  = "/v2/supplement/transaction_history.do"
)

const (
	PathDepth                      = "/v2/depth.do"
	PathSupplementTrades           = "/v2/supplement/trades.do"
	PathSupplementTickerPrice      = "/v2/supplement/ticker/price.do"
	PathSupplementTickerBookTicker = "/v2/supplement/ticker/bookTicker.do"
)

const (
	PathSupplementSystemStatus      = "/v2/supplement/system_status.do"
	PathSupplementUserInfo          = "/v2/supplement/user_info.do"
	PathSupplementWithdraw          = "/v2/supplement/withdraw.do"
	PathSupplementDepositHistory    = "/v2/supplement/deposit_history.do"
	PathSupplementWithdraws         = "/v2/supplement/withdraws.do"
	PathSupplementGetDepositAddress = "/v2/supplement/get_deposit_address.do"
	PathSupplementAssetDetail       = "/v2/supplement/asset_detail.do"
	PathSupplementCustomerTradeFee  = "/v2/supplement/customer_trade_fee.do"
	PathSupplementApiRestrictions   = "/v2/supplement/api_Restrictions.do"
)

const (
	PathCurrencyPairs   = "/v2/currencyPairs.do"
	PathAccuracy        = "/v2/accuracy.do"
	PathUsdToCny        = "/v2/usdToCny.do"
	PathWithdrawConfigs = "/v2/withdrawConfigs.do"
	PathTimestamp       = "/v2/timestamp.do"
	PathTicker24hr      = "/v2/ticker/24hr.do"
	PathEtfTicker24hr   = "/v2/etfTicker/24hr.do"
	PathTicker          = "/v2/ticker.do"
	PathIncrDepth       = "/v2/incrDepth.do"
	PathTrades          = "/v2/trades.do"
	PathKline           = "/v2/kline.do"
)

const (
	PathUserInfo            = "/v2/user_info.do"
	PathSubscribeGetKey     = "/v2/subscribe/get_key.do"
	PathSubscribeRefreshKey = "/v2/subscribe/refresh_key.do"
	PathSubscribeDestroyKey = "/v2/subscribe/destroy_key.do"
	PathGetDepositAddress   = "/v2/get_deposit_address.do"
	PathDepositHistory      = "/v2/deposit_history.do"
)

const (
	PathCreateOrder            = "/v2/create_order.do"
	PathBatchCreateOrder       = "/v2/batch_create_order.do"
	PathCancelOrder            = "/v2/cancel_order.do"
	PathCancelClientOrders     = "/v2/cancel_clientOrders.do"
	PathOrdersInfo             = "/v2/orders_info.do"
	PathOrdersInfoHistory      = "/v2/orders_info_history.do"
	PathOrderTransactionDetail = "/v2/order_transaction_detail.do"
	PathTransactionHistory     = "/v2/transaction_history.do"
	PathOrdersInfoNoDeal       = "/v2/orders_info_no_deal.do"
)

const (
	PathWithdraw       = "/v2/withdraw.do"
	PathWithdrawCancel = "/v2/withdrawCancel.do"
	PathWithdraws      = "/v2/withdraws.do"
)

const (
	PathWs = "wss://www.lbkex.net/ws/V2/"
)
