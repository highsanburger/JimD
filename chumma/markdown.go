package main

import (
	// "bufio"
	// "fmt"
	"os"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	time := time.Now().Format("02/01/2006 15:04")
	e := os.WriteFile("test.md", []byte(time), 0644)
	check(e)

	// d1 := []byte("hello\ngo\n")
	// err := os.WriteFile("/tmp/dat1", d1, 0644)
	// check(err)
	//
	// f, err := os.Create("/tmp/dat2")
	// check(err)
	//
	// defer f.Close()
	//
	// d2 := []byte{115, 111, 109, 101, 10}
	// n2, err := f.Write(d2)
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n2)
	//
	// n3, err := f.WriteString("writes\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n3)
	//
	// f.Sync()

	// w := bufio.NewWriter(f)
	// n4, err := w.WriteString("buffered\n")
	// check(err)
	// fmt.Printf("wrote %d bytes\n", n4)
	//
	// w.Flush()
	//
}
