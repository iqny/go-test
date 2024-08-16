package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	//  w.Write([]byte("hello"))
	var (
		wsConn      *websocket.Conn
		err         error
		message     []byte
		messageType int
	)
	// 完成ws协议的握手操作
	// Upgrade:websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	// 启动线程，不断发消息
	go func() {
		for {
			messageType, message, err = wsConn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("received %s\n", message)

			if err = wsConn.WriteMessage(messageType, []byte("heartbeat")); err != nil {
				return
			}
			fmt.Printf("send %s\n", message)
			time.Sleep(1 * time.Second)
		}
	}()

}

func main() {

	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
