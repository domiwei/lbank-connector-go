package sve

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"

	"github.com/gorilla/websocket"
)

type WsService struct {
	conn *websocket.Conn
	Wc   *WsClient
}

/*
	func (w *WsService) KeepAlive(timeout time.Duration) {
		ticker := time.NewTicker(timeout)
		go func() {
			defer ticker.Stop()
			for {
				pingMap := map[string]string{
					"ping":   pkg.RandomUUID(),
					"action": "ping",
				}
				msg, _ := json.Marshal(pingMap)
				w.SendMsg(msg)
				<-ticker.C
			}
		}()
	}
*/
func (w *WsService) RespondServerPing(message []byte) {
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
	if err := w.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		log.Printf("send pong failed: %v", err)
	}
}

func (w *WsService) CreateWsConn() (*websocket.Conn, error) {
	u := url.URL{Scheme: "wss", Host: "www.lbkex.net", Path: "/ws/V2/"}
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	w.conn = conn
	return conn, err
}

/*
func (w *WsService) SendMsg(content []byte) {
	err := w.conn.WriteMessage(websocket.TextMessage, content)
	if err != nil {
		w.Wc.Logger.Fatal("SendMsg:", err)
		return
	}
	if w.Wc.Debug {
		w.Wc.debug("Sent: %s", content)

	}
	log.Printf("Sent: %s", content)
}

func (w *WsService) ReceiveMsg() {
	for {
		_, message, err := w.conn.ReadMessage()
		if err != nil {
			w.Wc.Logger.Fatal("read:", err)
			return
		}
		w.Wc.debug("Received: %s", message)
		fmt.Printf("Received: %s", message)
	}
}
*/

func isPingPong(message []byte) bool {
	return strings.Contains(string(message), "ping")
}
