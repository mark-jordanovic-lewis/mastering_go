package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
	// "errors" // HANDLE ERRORS CORRECTLY IN PROD CODE
)

// config files to be parsed for this data.
var ip = "127.0.0.1"
var port = "2300"
var address = ip+":"+port

func main()  {
	// Getting user input for chat handle
	rand.Seed(time.Now().UnixNano())
	var name = fmt.Sprintf("Anon:%d", rand.Intn(400))
	var inString string
	fmt.Println("Starting Chat Client")
	fmt.Print("Input handle: ")
	_, _ = fmt.Scanln(&inString)
	if !(inString == "") {
		name = inString
	}
	name += ": "

	// Connecting to chat and scanning for input from net.Conn object (is a ReadWriteCloser)
	fmt.Println("Hello %s, connecting to chat...")
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Cannot connect to chat")
	}
	fmt.Println("Connected to chat")
	defer func(){
		_ = conn.Close()
	}()
	go func() {
		var scanner = bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	// send messages to chat server in main process
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var msg= scanner.Text()
		_, err := fmt.Fprintf(conn, name+msg+"\n")
		if err != nil {
			break
		}
	}
}
