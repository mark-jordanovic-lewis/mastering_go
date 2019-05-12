package chat

import (
	"fmt"
	"io"
	"net"
	"sync"
	"../../singleton/hlogger"
)

var logger = hlogger.GetInstance()

type room struct {
	name string
	Msgch chan string
	clients map[chan<-string]struct{}
	Quit chan struct{}
	*sync.RWMutex
}

func CreateRoom(name string) *room  {
	var room *room = &room{
		name: name,
		Msgch: make(chan string),
		RWMutex: sync.RWMutex{},
		clients: make(map[chan<-string]struct{}),
		Quit: make(chan struct{}),
	}
	room.Run()
	return room
}

func (self *room) Run()  {
	logger.Println("Starting chat room ", self.name)
	go func() {
		for msg := range self.Msgch {
			self.broadcastMsg(msg)
		}
	}()
}

func (self *room) broadcastMsg(msg string)  {
	self.RLock()
	defer self.RUnlock()
	fmt.Println("Received message: ", msg)
	for client, _ := range self.clients {
		// spawn go routines for execution rather than execute the channel send inside the loop
		go func(client chan <-string) {
			client<-msg
		}(client) // pass by copy/value so that the value doesn't change with loop variable
	}
}

func (self *room) AddClient(client io.ReadWriteCloser) {
 	self.Lock()
 	clientChan, done := StartClient(self.Msgch, client, self.Quit)
 	self.clients[clientChan] = struct{}{}
 	self.Unlock() // not deferred as further code to execute

 	// remove client when client is done - setting a trigger for the client kill
 	go func() {
 		<-done
 		self.RemoveClient(clientChan)
 	}()
}

func (self *room) ClCount() int {
	return len(self.clients)
}

func (self *room) RemoveClient(clientChan chan<- string)  {
	logger.Println("Reomving client")
	self.Lock()
	close(clientChan)
	delete(self.clients, clientChan)
	self.Unlock()
	select {
	case <-self.Quit:
		if len(self.clients) == 0 {
			close(self.Msgch)
		}
	default:
		// NoOp
	}
}

