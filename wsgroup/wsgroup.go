// package wsgroup implements a convienient way to multiplex messages to a group
// of websocket clients
package wsgroup

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"net/http"
)

type wsgroup struct {
	members    map[*websocket.Conn]bool
	messages   chan []byte
	newMembers chan *websocket.Conn
}

func NewGroup() *wsgroup {
	return &wsgroup{
		members:    make(map[*websocket.Conn]bool),
		messages:   make(chan []byte),
		newMembers: make(chan *websocket.Conn),
	}
}

func (wg *wsgroup) Start() {
	go func() {
		for {
			select {
			case member := <-wg.newMembers:
				wg.members[member] = true
			case message := <-wg.messages:
				for member, _ := range wg.members {
					err := websocket.Message.Send(member, string(message))
					if err != nil {
						delete(wg.members, member)
					}
				}
			}
		}
	}()
}

func (wg *wsgroup) SendJSON(value interface{}) error {
	if bytes, err := json.MarshalIndent(value, "", "  "); err != nil {
		return err
	} else {
		wg.messages <- bytes
	}
	return nil
}

//func (wg *wsgroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
func (wg *wsgroup) Handler() http.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		wg.newMembers <- ws

		//var message string
		websocket.Message.Receive(ws, nil)
	})
}
