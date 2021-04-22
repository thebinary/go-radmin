package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/thebinary/go-radmin/libradmin"
)

var sockAddr string
var command string

func main() {
	flag.StringVar(
		&sockAddr, "f",
		"/usr/local/var/run/radiusd/radiusd.sock",
		"Radiusd control socket file",
	)
	flag.StringVar(
		&command, "e",
		"show uptime",
		"Command to execute",
	)

	flag.Parse()

	r, err := libradmin.NewRadminClient(sockAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	p := make([]byte, 256)

	r.Write([]byte(command))
	for true {
		n, _ := r.Read(p)
		if int(r.LastReadChannel()) == libradmin.FR_CHANNEL_CMD_STATUS {
			break
		}
		fmt.Println(string(p[:n-1]))
	}
	os.Exit(int(r.LastReadStatus()))
}
