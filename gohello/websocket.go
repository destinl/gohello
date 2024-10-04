package gohello

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WsServer struct {
	*websocket.Conn
}

func StartWebSocketServer() {
	upgrader := &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // 允许所有源
			// 或者根据你的需求进行特定源的检查
			// origin := r.Header.Get("Origin")
			// return origin == "http://your-allowed-origin.com"
		},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("upgrade error: %v", err)
			w.Write([]byte("upgrade error"))
			return
		}
		conn := &WsServer{Conn: c}

		// 读取消息
		go func(conn *WsServer) {
			for {
				typ, msg, err := conn.ReadMessage()
				if err != nil {
					log.Printf("read message error: %v", err)
					conn.Close()
					return
				}
				switch typ {
				case websocket.CloseMessage:
					conn.Close()
					return
				default:
					log.Printf("msg: %s", msg)
				}
			}
		}(conn)

		// 发送消息
		go func(conn *WsServer) {
			ticker := time.NewTicker(time.Second * 3)
			for now := range ticker.C {
				err := conn.WriteMessage(websocket.TextMessage, []byte("Hello "+now.String()))
				if err != nil {
					log.Printf("write message error: %v", err)
					conn.Close()
					return
				}
			}
		}(conn)

		defer conn.Close()
	})

	// ws://localhost:8081/ws
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("listen and serve error: %v", err)
	}
}
