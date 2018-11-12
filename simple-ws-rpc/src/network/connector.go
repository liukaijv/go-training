package network

import (
	"common"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Connector struct {
	*common.Dispatcher
	conn *websocket.Conn
}

func NewConnector(url string) (connector *Connector, err error) {

	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	connector = &Connector{
		common.NewDispatcher(),
		c,
	}
	connector.readPump()
	return
}

func (c *Connector) Call(args ...interface{}) (err error) {

	defer common.ErrRecover(args)
	data, err := json.Marshal(args)
	if err != nil {
		log.Printf("Connector.Call [json.Marshal] err: %v, args: %v", err, args)
		return
	}
	log.Println("Agent.Call", string(data))

	return c.conn.WriteMessage(websocket.TextMessage, data)

}

func (c *Connector) readPump() {

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
		log.Println("Connector receive message: ", string(message))
		var arr = make([]interface{}, 0)
		json.Unmarshal(message, &arr)
		if len(arr) == 0 {
			log.Printf("Connector [json.Unmarshal] err %v", err)
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
