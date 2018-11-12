package network

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"simple-ws-rpc/src/common"
	"sync"
	"time"
)

const (
	writeWait      = 10 * time.Second
	maxMessageSize = 512
	pongWait       = 10 * time.Second
	pingInterval   = (pongWait * 8) / 10
)

type Agent struct {
	*sync.Mutex
	*common.Dispatcher
	conn *websocket.Conn
	send chan []byte
}

func NewClient(conn *websocket.Conn) *Agent {
	client := new(Agent)
	client.conn = conn
	client.Mutex = new(sync.Mutex)
	client.Dispatcher = common.NewDispatcher()
	client.send = make(chan []byte)

	client.InitFunc()
	return client
}

func (c *Agent) write(mt int, payload []byte) error {
	c.Lock()
	defer c.Unlock()
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}

func (c *Agent) writeTextMessage(message []byte) error {
	return c.write(websocket.TextMessage, message)
}

func (c *Agent) readPump() {

	defer func() {
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Agent.readPump error: %v", err)
			}
			break
		}
		log.Println("Agent.readPump receive message: ", string(message))
		var arr = make([]interface{}, 0)
		json.Unmarshal(message, &arr)
		if len(arr) == 0 {
			log.Printf("Agent.readPump [json.Unmarshal] err %v", err)
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

func (c *Agent) writePump() {
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			err := c.writeTextMessage(message)
			if err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Agent) Call(args ...interface{}) {
	defer common.ErrRecover(args)
	data, err := json.Marshal(args)
	if err != nil {
		log.Printf("Agent.Call [json.Marshal] err: %v, args: %v", err, args)
		return
	}
	log.Println("Agent.Call", string(data))
	c.send <- data
}

func (c *Agent) InitFunc() {
	c.AddFunc("onConnect", c, "Connected")
	c.AddFunc("LoginReq", c, "LoginReq")
}

func (c *Agent) Connected() {

	log.Println("onConnect")

}
func (c *Agent) LoginReq(name, password string) {
	log.Println(name, password)
	c.Call("LoginRes", true)

}
