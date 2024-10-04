package websocket

import "encoding/json"

var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

type hub struct {
	c map[*connection]bool //一个映射，用于存储当前连接的用户，其键是连接对象指针，值是布尔值（用于表示连接是否有效）。
	b chan []byte          //一个字节通道，用于广播消息。
	r chan *connection     //一个连接通道，用于接收新的连接请求
	u chan *connection     //一个连接通道，用于接收需要断开的连接。
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.r:
			h.c[c] = true
			c.data.Ip = c.ws.RemoteAddr().String()
			c.data.Type = "handshake"
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			c.sc <- data_b
		case c := <-h.u:
			if _, ok := h.c[c]; ok {
				delete(h.c, c)
				close(c.sc)
			}
		case data := <-h.b:
			for c := range h.c {
				select {
				case c.sc <- data:
				default:
					delete(h.c, c)
					close(c.sc)
				}
			}
		}
	}
}
