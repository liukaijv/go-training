package simple_websoket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"sync"
	"time"
)

const (
	writeWait      = 10 * time.Second
	maxMessageSize = 512
	pongWait       = 10 * time.Second
	pingInterval   = (pongWait * 8) / 10
)

type Conn struct {
	*sync.Mutex
	*Dispatcher
	conn *websocket.Conn
	send chan []byte
}

func NewConn(conn *websocket.Conn) *Conn {
	c := new(Conn)
	c.conn = conn
	c.Mutex = new(sync.Mutex)
	c.Dispatcher = NewDispatcher()
	c.send = make(chan []byte)

	go c.readPump()
	go c.writePump()
	return c
}

func (c *Conn) write(mt int, payload []byte) error {
	c.Lock()
	defer c.Unlock()
	c.conn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.conn.WriteMessage(mt, payload)
}

func (c *Conn) writeTextMessage(message []byte) error {
	return c.write(websocket.TextMessage, message)
}

func (c *Conn) readPump() {

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
				log.Printf("Conn.readPump error: %v", err)
			}
			break
		}
		log.Println("Conn.readPump receive message: ", string(message))
		var arr = make([]interface{}, 0)
		json.Unmarshal(message, &arr)
		if len(arr) == 0 {
			log.Printf("Conn.readPump [json.Unmarshal] err %v", err)
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

func (c *Conn) writePump() {
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

func (c *Conn) Call(args ...interface{}) {
	//defer ErrRecover(args)
	data, err := json.Marshal(args)
	if err != nil {
		log.Printf("Conn.Call [json.Marshal] err: %v, args: %v", err, args)
		return
	}
	log.Println("Conn.Call", string(data))
	c.send <- data
}
