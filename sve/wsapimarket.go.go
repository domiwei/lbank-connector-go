package sve

import (
	"time"
)

type WsMarketService struct {
	Ws *WsService
}

func (w *WsMarketService) Kbar(kbar, pair string) {
	conn, _ := w.Ws.CreateWsConn()
	defer conn.Close()
	//payloay := map[string]string{
	//	"action":    "subscribe",
	//	"subscribe": "kbar",
	//	"kbar":      kbar,
	//	"pair":      pair,
	//}
	w.Ws.KeepAlive(10 * time.Second)
	//msg, _ := json.Marshal(payloay)
	msg := []byte(`{"action":"subscribe","subscribe":"kbar","kbar":"5min","pair":"btc_usdt"}`)
	go w.Ws.SendMsg(msg)
	go w.Ws.ReceiveMsg()
	select {}
}
