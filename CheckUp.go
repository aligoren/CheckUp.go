package main

import (
	"log"
	"net"
	"time"
	"os"
	"fmt"
)


func checkStatus(d time.Duration, f func(time.Time)) {
	fmt.Println("Initializing..")
	for x := range time.Tick(d) {
		f(x)
	}
}

func checkPort(t time.Time) {
	var status string
	port := "80"
	host := "localhost:"
	if len(os.Args) == 2 {
		port = os.Args[1]
	} else if len(os.Args) == 3 {
		host = os.Args[2] + ":"
	}
	conn, err := net.Dial("tcp", host+port)
	if err != nil {
		status = "Unreachable"
	} else {
		status = "Online"
		defer conn.Close()
	}

	log.Println(status)
}

func main() {
	checkStatus(5*time.Second, checkPort)
}
