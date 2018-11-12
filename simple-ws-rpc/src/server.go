package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"simple-ws-rpc/src/network"
)

func main() {

	var addr = flag.String("addr", fmt.Sprintf(":%d", 8889), "http service address")

	http.HandleFunc("/ws", network.ServeWs)

	log.Println("ListenAndServe at: ", *addr)

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Printf("ListenAndServe: err:%v", err)
	}

}
