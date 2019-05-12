package chat

import (
	"../../singleton/hlogger"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	// "errors" // IN PROD CODE HANDLE THESE ERRORS PLEASE!
)

var logger = hlogger.GetInstance()

func Run(connection string) error {
	listenStream, err := net.Listen("tcp", connection )
	if err != nil {
		logger.Println("Error connecting to tcp stream")
		fmt.Println("Error connecting to tcp stream")
		return err
	}

	var chatroom = CreateRoom("HydraChat")

	// this routine is a handy little beasty - should be in all net code
	go func() {
		// open observer channel on os signaling system
		var osSigChannel = make(chan os.Signal)
		signal.Notify(osSigChannel, syscall.SIGINT, syscall.SIGTERM)
		var exitCode = <-osSigChannel

		// clean up resources on microservice shutdown
		_ = listenStream.Close()
		logger.Printf("Closing tcp stream: %v", exitCode)
		fmt.Println("Closing tcp stream")
		close(chatroom.Quit)
		if chatroom.ClCount() > 0 {
			<-chatroom.Msgch
		}
		os.Exit(0)
	}()

	for {
		conn, err := listenStream.Accept()
		if err != nil {
			logger.Println("Chat client accept error", err)
			break
		}
		go handleConnection(chatroom, conn)
	}
	return err
}

func handleConnection(chatroom *room, conn net.Conn)  {
	logger.Println("Client connection request received")
	chatroom.AddClient(conn)
}
