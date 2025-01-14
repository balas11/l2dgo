package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//default Logger
	log.Println("Message from default logger!")

	log.SetFlags(log.Lshortfile | log.LstdFlags | log.Lmicroseconds)
	log.Println("Message using flags")

	log.SetPrefix("[INFO] ")
	log.Println("Message with prefix")

	log.Printf("This message is printed with prefix %s, and flags %d", log.Prefix(), log.Flags())

	//custom Logger
	customLogger := log.New(os.Stdout, "CUSTOM: ", log.Lshortfile|log.LstdFlags)
	customLogger.Println("Message from custom logger!")

	//custom Logger to byte buffer
	var buf bytes.Buffer

	customLogger = log.New(&buf, "CUSTOM: ", log.Lshortfile|log.LstdFlags)
	customLogger.Println("Message from custom logger to buffer!")
	fmt.Println("Printing from buffer : ", buf.String())

	//custom Logger to a file
	file, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fileLogger := log.New(file, "FILE: ", log.Lshortfile|log.LstdFlags)
	fileLogger.Println("Message from file logger!")

	file.Close()

	//To file and stderr
	file1, err := os.OpenFile("custom.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writer := io.MultiWriter(os.Stderr, file1)
	mwLogger := log.New(writer, "FILE: ", log.Lshortfile|log.LstdFlags)

	mwLogger.Println("Message from mw file logger!")

	file1.Close()

}
