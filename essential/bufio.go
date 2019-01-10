package main

import (
	"bufio"
	"fmt"
	"github.com/benhalstead/gotraining/tutorial"
	"io/ioutil"
	"os"
)

func main() {

	/*
		File and stream IO in golang is spread over a number of packages. The low-level functions for performing
		IO are in the io package, but you will more commonly use the higher-level functions in ioutil, bufio and text/scanner
	*/

	exampleWritingAFile()
	exampleReadingFiles()
}

// We have hard-coded a UNIX style file path
// The correct way to build a platforn independent path in go is to use the path/filepath package
var testFile = "/tmp/go-example"

func exampleWritingAFile() {

	tutorial.Section("Writing files")

	// If you have a byte array or string, there is a one-liner to write that data to a file
	// The last argument
	if err := ioutil.WriteFile(testFile, []byte("Hello, world!\n"), 0600); err != nil {
		fmt.Println(err)
	}

	//Write file will overwrite any existing content

	//If you want to generate the content of a file while you're writing it, you will need to open the file explicity and use a Writer
	var f *os.File
	var err error

	if f, err = os.Create(testFile); err != nil { //Use os.Open if you want to append
		fmt.Println(err.Error())
		return
	}

	defer f.Close() // Defer the closing of a file as soon as possible after is has been successfully opened

	// file.File already implements the io.Writer interface, but if you use it directly, any writes go straight to disk
	// which is often undesirable. Instead, we can wrap the file in a Buffered Writer

	b := bufio.NewWriter(f)
	defer b.Flush()

	for i := 0; i < 10; i++ {

		line := fmt.Sprintf("I am line %d\n", i)
		fmt.Print(line)

		b.WriteString(line)

	}
}

func exampleReadingFiles() {

	tutorial.Section("Reading files")

	//Go provides a function to read a whole file into a byte array
	if content, err := ioutil.ReadFile(testFile); err == nil {

		// The content is a byte array (which is appropriate, as the most common use case for this method is reading binary files)
		// For the purposes of this example, we can convert it to a string like:
		s := string(content)

		fmt.Println(s)

	} else {
		fmt.Println(err.Error())
		return
	}

	// If you a processing a text file, it is more likely that you will want to use a bufio.Scanner
	var f *os.File
	var err error

	if f, err = os.Open(testFile); err != nil { // os.Open defaults to read only
		fmt.Println(err.Error())
		return
	}

	defer f.Close()

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		fmt.Printf("Line read: %s\n", line)
	}

	// If you need to tokenize your file, look at  https://golang.org/pkg/text/scanner/

}
