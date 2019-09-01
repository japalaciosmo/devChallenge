

package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
	"github.com/japalaciosmo/devChallenge/bot/parser"

	"github.com/gorilla/websocket"
	"github.com/subosito/gotenv"
)



func main() {

	if err := gotenv.Load("env"); err != nil {
		log.Printf("error: %s",err.Error())
	}

	var address=os.Getenv("ADDRESS")
	if address==""{
		address="localhost:8080"
	}

	var addr = flag.String("addr", address, "http service address")

	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

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
			command,resend, err:=parser.ParseMessage(message)
			fmt.Println(command,resend)
			if err != nil {
				log.Println("read:", err)
				return
			}
			if resend{
				err := c.WriteMessage(websocket.TextMessage, []byte(command))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}
		}
	}()


	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
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