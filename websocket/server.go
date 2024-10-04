package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Websocket_main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", myws)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("Server started on port 8080")
}
