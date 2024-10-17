package sve

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/gorilla/websocket"
)

type WsPrivateService struct {
	ws  *WsService
	acc *AccountService
}

type WsOrderUpdateData struct {

	// {"SERVER":"V2","orderUpdate":{"accAmt":"0","amount":"0","avgPrice":"0","customerID":"","orderAmt":"0.8328","orderPrice":"0","orderStatus":0,"price":"0","remainAmt":"0.8328","role":"taker","symbol":"lbk_usdt","type":"buy_market","updateTime":1729191582412,"uuid":"bd724682-0bb8-4f05-ab11-f578c0c0b729","volumePrice":"0"},"type":"orderUpdate","pair":"lbk_usdt","TS":"2024-10-18T02:59:42.413"}
	OrderUpdate struct {
		AccAmt      string `json:"accAmt"`
		Amount      string `json:"amount"`
		AvgPrice    string `json:"avgPrice"`
		CustomerID  string `json:"customerID"`
		OrderAmt    string `json:"orderAmt"`
		OrderPrice  string `json:"orderPrice"`
		OrderStatus int    `json:"orderStatus"`
		Price       string `json:"price"`
		RemainAmt   string `json:"remainAmt"`
		Role        string `json:"role"`
		Symbol      string `json:"symbol"`
		Type        string `json:"type"`
		UpdateTime  int64  `json:"updateTime"`
		UUID        string `json:"uuid"`
		VolumePrice string `json:"volumePrice"`
	} `json:"orderUpdate"`
	Server string `json:"SERVER"`
	Type   string `json:"type"`
	Pair   string `json:"pair"`
	Ts     string `json:"TS"`
}

func (w *WsPrivateService) SubscribeOrderUpdate(pair string, errHandler ErrHandle) (<-chan *WsOrderUpdateData, error) {
	conn, err := w.ws.CreateWsConn()
	if err != nil {
		return nil, err
	}
	// subscribe key
	resp, err := w.acc.SubscribeGetKey()
	if err != nil {
		return nil, err
	}
	subkey := resp.Data
	keyRefreshTime := time.Now()
	/*
			{
		      "action": "subscribe",
		      "subscribe": "orderUpdate",
		      "subscribeKey": "24d87a4xxxxxd04b78713f42643xxxxf4b6f6378xxxxx35836260",
		      "pair": "all",
		    }
	*/
	msg := []byte(`{"action":"subscribe","subscribe":"orderUpdate","subscribeKey":"` + subkey + `","pair":"` + pair + `"}`)
	if err := w.ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return nil, err
	}
	ch := make(chan *WsOrderUpdateData, 1024)
	go func() {
		defer close(ch)
		defer w.acc.SubscribeDestroyKey(subkey)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				errHandler(err)
				return
			}
			if isPingPong(message) {
				w.ws.RespondServerPing(message)
				continue
			}
			var data WsOrderUpdateData
			if err := json.Unmarshal(message, &data); err != nil {
				errHandler(err)
				return
			}
			if data.Type == "orderUpdate" {
				ch <- &data
			}
			if err := w.checkRefreshKey(subkey, &keyRefreshTime); err != nil {
				errHandler(err)
				return
			}
		}
	}()
	return ch, nil
}

func (w *WsPrivateService) checkRefreshKey(subkey string, keyRefreshTime *time.Time) error {
	if time.Since(*keyRefreshTime) > 15*time.Minute {
		success, err := w.acc.SubscribeRefreshKey(subkey)
		if err != nil {
			return err
		}
		if success {
			*keyRefreshTime = time.Now()
			return nil
		} else {
			return errors.New("refresh key failed")
		}
	}
	return nil
}
