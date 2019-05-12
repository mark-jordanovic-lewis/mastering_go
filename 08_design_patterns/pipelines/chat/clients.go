package chat

import (
	"bufio"
	"io"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	messageChannel chan string
}

func StartClient(msgCh chan<- string, connection io.ReadWriteCloser, quit chan struct{}) (chan<- string, <-chan struct{}) {
	var client = client{
		Reader: bufio.NewReader(connection),
		Writer: bufio.NewWriter(connection),
		messageChannel: make(chan string),
	}
	var done = make(chan struct{})

	go func() {
		var scanner = bufio.NewScanner(client.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			msgCh<- scanner.Text()
		}
		done<-struct{}{}
	}()

	client.writeMonitor()

	go func() {
		select {
		case <-quit:
			_ = connection.Close()
		case <-done:
			// NoOp - exits the go routine (select is not in `for`)
		}
	}()

	return client.messageChannel, done
}

func (self *client) writeMonitor()  {
	go func() {
		for s := range self.messageChannel {
			_, _ = self.WriteString(s+"\n")
			_ = self.Flush()
		}
	}()
}
