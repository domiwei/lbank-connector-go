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
	/*
			{
		      "orderUpdate":{
		          "accAmt": "0.003",
		          "amount":"0.003",
		          "avgPrice": "0.02455211",
		          "symbol":"eth_btc",
		          "type":"buy",
		          "orderAmt": "0.003",
		          "orderStatus":2,
		          "orderPrice": "0.02455211",
		          "price":"0.02455211",
		          "role":"maker",
		          "remainAmt":"0",
		          "updateTime":1561704577786,
		          "uuid":"d0db191d-xxxxx-4418-xxxxx-fbb1xxxx2ea9",
		          "txUuid":"da88f354d5xxxxxxa12128aa5bdcb3",
		          "volumePrice":"0.00007365633"
		      },
		      "pair":"eth_btc",
		      "type":"orderUpdate",
		      "SERVER":"V2",
		      "TS":"2019-06-28T14:49:37.816"
		  }
	*/
	OrderUpdate struct {
		AccAmt      float64 `json:"accAmt"`
		Amount      float64 `json:"amount"`
		AvgPrice    float64 `json:"avgPrice"`
		Symbol      string  `json:"symbol"`
		Type        string  `json:"type"`
		OrderAmt    float64 `json:"orderAmt"`
		OrderStatus int     `json:"orderStatus"`
		OrderPrice  float64 `json:"orderPrice"`
		Price       float64 `json:"price"`
		Role        string  `json:"role"`
		RemainAmt   float64 `json:"remainAmt"`
		UpdateTime  int64   `json:"updateTime"`
		Uuid        string  `json:"uuid"`
		TxUuid      string  `json:"txUuid"`
		VolumePrice float64 `json:"volumePrice"`
	} `json:"orderUpdate"`
	Pair   string `json:"pair"`
	Type   string `json:"type"`
	Server string `json:"SERVER"`
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
	subkey := resp.Key
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
			if err := w.checkRefresh(subkey, &keyRefreshTime); err != nil {
				errHandler(err)
				return
			}
		}
	}()
	return ch, nil
}

func (w *WsPrivateService) checkRefresh(subkey string, keyRefreshTime *time.Time) error {
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
