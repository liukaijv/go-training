package simple_websoket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	*Dispatcher
	conn *websocket.Conn
}

func NewClient(url string) (client *Client, err error) {

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	client = &Client{
		NewDispatcher(),
		c,
	}
	go client.readPump()
	return
}

func (c *Client) Call(args ...interface{}) (err error) {

	//defer ErrRecover(args)
	data, err := json.Marshal(args)
	if err != nil {
		log.Printf("Client.Call [json.Marshal] err: %v, args: %v", err, args)
		return
	}
	log.Println("Conn.Call", string(data))

	return c.conn.WriteMessage(websocket.TextMessage, data)

}

func (c *Client) readPump() {

	defer func() {
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if err != nil {
				log.Println("read:", err)
				return
			}
		}
		log.Println("Client receive message: ", string(message))
		var arr = make([]interface{}, 0)
		json.Unmarshal(message, &arr)
		if len(arr) == 0 {
			log.Printf("Client [json.Unmarshal] err %v", err)
			return
		} else {
			if len(arr) > 1 {
				c.Run(arr[0].(string), arr[1:]...)
			} else {
				c.Run(arr[0].(string))
			}
		}
	}

}
