package main

import (
	"flag"
	go_wol_server "github.com/mint-rabi/go-wol-server"
)

func main() {
	address := flag.String("address", "localhost", "listen address")
	port := flag.Int("port", 8080, "listen TCP port")

	flag.Parse()
	go_wol_server.NewServer(*address, *port)
}
