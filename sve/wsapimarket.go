package sve

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type ErrHandle func(err error)

type WsMarketService struct {
	Ws *WsService
}

/*
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
*/
type WsTradeData struct {
	// eg: {"trade":{"volume":0.1167,"amount":7858.687698,"price":67340.94,"direction":"buy","TS":"2024-10-18T00:44:51.998"},"SERVER":"V2","type":"trade","pair":"btc_usdt","TS":"2024-10-18T00:44:52.005"}
	Trade struct {
		Volume    float64 `json:"volume"`
		Amount    float64 `json:"amount"`
		Price     float64 `json:"price"`
		Direction string  `json:"direction"`
		Ts        string  `json:"TS"`
	} `json:"trade"`
	Server string `json:"SERVER"`
	Type   string `json:"type"`
	Pair   string `json:"pair"`
	Ts     string `json:"TS"`
}

func (w *WsMarketService) SubscribeTrade(pair string, errHandler ErrHandle) (<-chan *WsTradeData, error) {
	conn, err := w.Ws.CreateWsConn()
	if err != nil {
		return nil, err
	}
	msg := []byte(`{"action":"subscribe","subscribe":"trade","pair":"` + pair + `"}`)
	if err := w.Ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return nil, err
	}
	dataCh := make(chan *WsTradeData, 1024)
	go func() {
		defer close(dataCh)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				errHandler(err)
				return
			}
			if isPingPong(message) {
				w.Ws.RespondServerPing(message)
				continue
			}
			var tradeData WsTradeData
			if err := json.Unmarshal(message, &tradeData); err != nil {
				errHandler(err)
				return
			}
			if tradeData.Type == "trade" {
				dataCh <- &tradeData
			}
		}
	}()
	return dataCh, nil
}

type WsPingPongData struct {
	// eg: { "action":"ping", "ping":"0ca8f854-7ba7-4341-9d86-d3327e52804e" }
	Action string `json:"action"`
	Ping   string `json:"ping"`
}

/*
func (w *WsMarketService) keepAlive(message []byte) {
	var pingPongData WsPingPongData
	if err := json.Unmarshal(message, &pingPongData); err != nil {
		log.Printf("unmarshal ping pong data failed: %v. Data %v", err, string(message))
		return
	}
	pong := map[string]string{
		"pong":   pingPongData.Ping,
		"action": "pong",
	}
	msg, _ := json.Marshal(pong)
	if err := w.Ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("send pong failed: %v", err)
	}
}
*/
