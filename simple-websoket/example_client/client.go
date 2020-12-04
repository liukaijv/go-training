package main

import (
	"log"
	"os"
	"os/signal"
	"simple_websoket"
)

func main() {

	client, err := simple_websoket.NewClient("ws://localhost:8889/ws")

	if err != nil {
		log.Fatalln(err)
	}

	client.AddFunc("LoginRes", func(flag bool) {
		log.Println(flag)
	})

	client.Call("LoginReq", "demo", "111111")

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
