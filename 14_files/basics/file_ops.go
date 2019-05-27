package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

// File struct implements ReadWriteCloser interface
// FileStats is a handy inspection struct

func main() {
	// read only (this should return us a non-existence error
	// ---------
	// file1, err := os.Open("example1.txt")
	// handleError(err)
	// defer func() { _ = file1.Close() }()

	// create and return handle for
	// ----------------------------
	// file2, err := os.Create("example2.txt")
	// handleError(err)
	// defer func() { _ = file2.Close() }()

	// os package has a bunch of functions to move&rename (Rename), delete (Remove).
	// general file handle with filename, capabilities, file permission bits set
	// Capabilities:
	//   os.O_RDONLY
	//   os.O_WRONLY
	//   os.O_RDWR
	//   os.O_APPEND
	//   os.O_CREATE
	//   os.O_TRUNC
	//   caps are combined with | as they are bitmapped
	// usage:
	// ------
	// file3, err := os.OpenFile("example3.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// handleError(err)
	// defer func() { _ = file3.Close() }()



	// reading a file to a byte slice:
	orwell1984, err := os.Open("../assets/1984.html") // from here
	handleError(err)                                        //
	// defer func() { _ = orwell1984.Close() }()            //
	b, err := ioutil.ReadAll(orwell1984)                    // to here is captured in `b, err = ioutil.ReadFile(<filename>)`
	handleError(err)
	fmt.Printf("%s", b)
	closeFile(orwell1984)

	// reading line by line using the bufio.Scanner
	orwell1984, err = os.Open("../assets/1984.html")
	handleError(err)
	var fileReader = bufio.NewScanner(orwell1984)
	var bbCount = 0
	var lines = []string{}
	for fileReader.Scan() {
		var line = fileReader.Text()
		regex, _ := regexp.Compile("Big Brother|big brother")
		var bigBrother = regex.MatchString(line)
		if bigBrother{
			bbCount++
			lines = append(lines, line)
		}
	}
	fmt.Println("'Big Brother' appears", bbCount, "times in", orwell1984.Name())
	out, err := os.OpenFile("BBLines.txt", os.O_CREATE|os.O_RDWR, 0666)
	handleError(err)
	// writing to file using bufio.Writer
	var scribe = bufio.NewWriter(out)
	for _, line := range lines {
		_ = scribe.WriteString(line)
	}
	_ = scribe.Flush()

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

