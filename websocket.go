package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

func connect_to_websocket() {
	var addr = os.Getenv("WS_URL_ROOT")

	ws_channel := Coinbase_Websocket_Channel_Subscribe{
		Name:        "ticker",
		Product_ids: []string{"BTC-USD"},
	}

	ws_request := Coinbase_Websocket_Subscribe{
		Type:        "subscribe",
		Product_ids: []string{"BTC-USD"},
		Channels:    []Coinbase_Websocket_Channel_Subscribe{ws_channel},
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	log.Printf("INFO: Websocket connecting to %s\n", addr)

	c, _, err := websocket.DefaultDialer.Dial(addr, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	msg, err := json.Marshal(ws_request)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Println("ERROR: Could not connect to Coinbase websocket: ", err)
	} else {
		log.Println("INFO: Connected to Coinbase websocket")
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ERROR: Message read error: ", err)
				return
			}
			log.Printf("DATA: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			log.Printf("INFO: Ticked %s\n", t.String())
			// err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			// if err != nil {
			// 	log.Println("write:", err)
			// 	return
			// }
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("ERROR: Websocket close error: ", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
