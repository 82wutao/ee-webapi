package main

import (
	"log"
	"os"
	"strconv"

	"github.com/82wutao/ee-webapi/webapi"
)

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("webapi listen_addr listen_port\n")
		return
	}
	addr := os.Args[1]
	port, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("get port argument error %+v\n", err)
		return
	}
	webapi.Launch(addr, int16(port))
}
