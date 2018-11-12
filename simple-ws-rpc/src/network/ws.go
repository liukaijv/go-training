package network

import (
	"common"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ServeWs(w http.ResponseWriter, r *http.Request) {
	defer common.ErrRecover()

	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(fmt.Sprintf("upgrader err: %v", err))
		return
	}

	agent := NewClient(conn)

	go agent.readPump()
	go agent.writePump()

	agent.Run("onConnect")

}
