package main

import (
	"fmt"
	"math/rand"
	"time"
)

const BreakCode int = 0

type WebSocket struct {
	Signal chan struct{}
	Close  chan struct{}
}

func NewWebSocket() *WebSocket {
	return &WebSocket{
		Signal: make(chan struct{}),
		Close:  make(chan struct{}),
	}
}

func main() {
	ws := NewWebSocket()

	go ws.Sender()
	go ws.Listener()

	select {}
}

// Can write in 2 func, break loop => return
func (w *WebSocket) Sender() {
	counter := 0
	for {
		<-w.Signal
	loop:
		for {
			select {
			case <-w.Close:
				fmt.Println("disconnect")
				break loop
			default:
				time.Sleep(500 * time.Millisecond)
				counter++
				fmt.Println(counter)
			}
		}
	}
}

func (w *WebSocket) Listener() {
	data := 1
	for {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

		if data == BreakCode {
			w.Close <- struct{}{}
		} else {
			fmt.Println("connect")
			w.Signal <- struct{}{}
		}

		data++
		data = data % 2
	}
}
