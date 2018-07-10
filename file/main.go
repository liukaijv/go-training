package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"sync"
	"os"
	"os/signal"
)

func getDog(resChan chan string) {
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println("get error", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read body", err)
		return
	}

	//{"status":"success","message":"https:\/\/images.dog.ceo\/breeds\/schnauzer-miniature\/n02097047_2910.jpg"}
	jsonObj := make(map[string]string)
	err = json.Unmarshal(body, &jsonObj)

	if err != nil {
		fmt.Println("json parse", err)
		return
	}

	resChan <- jsonObj["message"]

}

func main() {

	var (
		mutex   = new(sync.Mutex)
		limit   = 9
		resChan = make(chan string)
		dogs    = make([]string, 0)
	)

	var interrupt = make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("get url")
	for i := 0; i < limit; i++ {
		go getDog(resChan)
	}

	for {
		select {
		case <-interrupt:
			fmt.Println("interrupt")
			data, err := json.Marshal(dogs)
			if err != nil {
				fmt.Println("json marshal error", err)
				return
			}
			var f *os.File
			f, err = os.OpenFile("data.json", os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				fmt.Println("open file error", err)
				return
			}
			_, err = f.WriteString(string(data))
			if err == nil {
				fmt.Println("write to file")
			}
			return
		case d := <-resChan:
			mutex.Lock()
			fmt.Println(d)
			if d != "" {
				dogs = append(dogs, d)
			}
			mutex.Unlock()
		}
	}

}
