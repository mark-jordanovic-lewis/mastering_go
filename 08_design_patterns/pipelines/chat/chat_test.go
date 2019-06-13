package chat

import (
	"bufio"
	"math/rand"
	"net"
	"strings"
	"sync"
	"testing"
	"time"
	"fmt"
)

var once sync.Once

func chatServerFunc(t *testing.T) func() {
	return func() {
		t.Log("Starting chat server")
		if err := Run(":2300"); err != nil {
			t.Error("Could not stater the chat server", err)
			return
		} else {
			t.Log("Chat server up")
		}
	}
}

func readMsgFunc(conn net.Conn, inChan chan string, t *testing.T) {
	var scanner = bufio.NewScanner(conn)
	for scanner.Scan() {
		var msg = scanner.Text()
		var incoming = <-inChan
		if strings.Compare(msg, incoming) != 0 {
			t.Errorf("Chat message [%s] does not match incoming message [%s]", msg, incoming)
		}
	}
}

func TestServerConnection(t *testing.T) {
	t.Log("Testing chat connection")
	go once.Do(chatServerFunc(t))
	time.Sleep(1*time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	defer func() { _ = conn.Close() }()
	if err != nil {
		t.Fatal(err)
	}
}


func TestRun(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test in short mode")
	}
	t.Log("Testing chat server send and receive")
	f := chatServerFunc(t)
	go once.Do(f)

	time.Sleep(1*time.Second)
	rand.Seed(time.Now().UnixNano())
	var name = fmt.Sprintf("Anon%d", rand.Intn(400))
	t.Logf("Hello %s, connecting to chat.\n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		t.Fatal("Could not connect to hydra chat system", err)
	}
	t.Log("Connected to hydra chat system")
	name += ":"


	defer func() { _ = conn.Close() }()
	var msgChan = make(chan string)

	go readMsgFunc(conn, msgChan, t)

	for i := 0 ; i <= 10; i++ {
		var msgBody = fmt.Sprintf("Random Msg: %d", rand.Intn(400))
		var msg = name+msgBody
		_, err = fmt.Fprintf(conn, msg+"\n")
		if err != nil {
			t.Fatal(err)
		}
		msgChan<-msg
	}
}
