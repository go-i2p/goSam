package main

import (
  "flag"

	"github.com/eyedeekay/goSam"
	"github.com/getlantern/go-socks5"
	"log"
)

var (
  samaddr = flag.String("sam", "127.0.0.1:7656", "SAM API address to use")
  socksaddr = flag.String("socks", "127.0.0.1:7675", "SOCKS address to use")
)

func main() {
	sam, err := goSam.NewClient(*samaddr)
	if err != nil {
		panic(err)
	}
	log.Println("Client Created")

	// create a transport that uses SAM to dial TCP Connections
	conf := &socks5.Config{
		Dial:     sam.DialContext,
		Resolver: sam,
	}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", *socksaddr); err != nil {
		panic(err)
	}
}
