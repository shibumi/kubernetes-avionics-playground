package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

var addr = "localhost:5000"

type FGCommand struct {
	Command string `json:"command"`
	Node string `json:"node"`
}

func main() {
	u := url.URL{Scheme: "ws", Host: addr, Path: "/PropertyListener"}
	log.Printf("connecting to %s", u.String())
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			log.Println(err)
		}
	}(c)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	err = c.WriteJSON(FGCommand{
		Command: "addListener",
		Node:    "sim/time/gmt",
	})
	if err != nil {
		log.Println(err)
	}

	for {
		select {
		case <-done:
			return
		}
	}
}