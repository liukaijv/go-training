package main

import (
	"github.com/gorilla/websocket"
	"log"
	"simple_websoket"
)

type Player struct {
	*simple_websoket.Conn
}

func NewPlayer(conn *websocket.Conn) *Player {
	p := &Player{
		simple_websoket.NewConn(conn),
	}
	p.InitFunc()
	return p
}

func (player *Player) InitFunc() {
	player.AddFunc("onConnect", "Connected", player)
	player.AddFunc("LoginReq", "LoginReq", player)
}

func (player *Player) Connected() {

	log.Println("onConnect")

}
func (player *Player) LoginReq(name, password string) {
	log.Println(name, password)
	player.Call("LoginRes", true)
}
