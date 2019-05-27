package helpful

import (
	"io"
	"log"
	"os"
	"time"
	"fmt"
)

// This is a really nice use of goto and channels to watch the changing and deletion of files.
// Could be used for:
// - security to watch for files changing (known_hosts being added to etc)
// - running environments in dev mode to reload/recompile libraries etc
// Should be extended to include kill signal for graceful shutdown
func FileChangeNotifier(name string) chan bool {
	var changeChan = make(chan bool)
	go func() {
	Begin:
		currentState, err := os.Stat(name)
		if err != nil {
			fmt.Println("Awaiting creation of:", name)
			for {
				time.Sleep(1 * time.Second)
				currentState, err = os.Stat(name)
				if err == nil {
					changeChan<-true
					break
				}
			}
		}
		fmt.Println("Watching", name)
		for {
			comparisonState, err := os.Stat(name)
			if err != nil { goto Begin }
			if comparisonState.ModTime() != currentState.ModTime() {
				fmt.Println(name,"was modified")
				currentState = comparisonState
				changeChan<-true
			}
			time.Sleep(1*time.Second)
		}
	}()
	return changeChan
}

// There is no library function for copying a file:
func CopyFile(origin, destination string)  {
	in, err := os.Open(origin)
	handleError(err)
	defer closeFile(in)
	out, err := os.Create(destination)
	handleError(err)
	defer closeFile(out)

	_, err = io.Copy(in, out)
	handleError(err)
	// flush io buffer to output file
	err = out.Sync()
	handleError(err)
}

func closeFile(f *os.File)  {
	handleError(f.Close())
}

func handleError(err error) {
	if err != nil {
		log.Fatal("There was an err", err)
	}
}