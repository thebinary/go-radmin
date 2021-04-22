package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thebinary/go-radmin/libradmin"
)

var sockAddr string
var sockType string
var command string

func main() {
	flag.StringVar(
		&command, "e",
		"show uptime",
		"Command to execute",
	)
	flag.StringVar(
		&sockAddr, "a",
		"/usr/local/var/run/radiusd/radiusd.sock",
		"Radiusd control socket",
	)
	flag.StringVar(
		&sockType, "p",
		"unix",
		"socket type",
	)

	flag.Parse()

	r, err := libradmin.NewRadminClientWithConn(sockType, sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	result, _, err := r.Execute([]byte(command))
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range result {
		fmt.Println(string(line))
	}
	os.Exit(int(r.LastReadStatus()))
}
