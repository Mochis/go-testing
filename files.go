package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Create file
	f, _ := os.Create("gofile.go")
	f2, _ := os.Create("otherfile.log")

	// Open a file
	os.Open("gofile.go")

	// Write in file
	f.WriteString("package main\nimport \"fmt\"\n" +
		"func main(){ fmt.Println(\"The go metafile\") }")
	barr := []byte{'a', 'b', 'c', '\n'}
	f2.Write(barr)

	ioutil.WriteFile("anotherfile.log", []byte{'a', 'b'}, 0777)

	// Read file
	ioutil.ReadFile("anotherfile.log") // Read all file
	someBytes := make([]byte, 100)
	f.Sync() // sync changes
	f.Seek(0, 0) // Set the head on first byte of file
	f1Readed, _ := f.Read(someBytes)
	fmt.Println(f1Readed, "bytes readed from gofile:", string(someBytes))

	// Close files
	f.Close()
	f2.Close()
}
