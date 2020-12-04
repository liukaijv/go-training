package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"simple_websoket"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	defer simple_websoket.ErrRecover()

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("upgrader err: %v", err))
		return
	}

	NewPlayer(conn)

}
