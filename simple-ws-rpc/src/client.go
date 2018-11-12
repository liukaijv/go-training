package main

import (
	"log"
	"network"
	"os"
	"os/signal"
)

func main() {

	connector, err := network.NewConnector("ws://localhost:8889/ws")

	if err != nil {
		log.Fatalln(err)
	}

	connector.AddFunc("LoginRes", nil, func(flag bool) {
		log.Println(flag)
	})

	connector.Call("LoginReq", "demo", "111111")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			return
		}
	}

}
