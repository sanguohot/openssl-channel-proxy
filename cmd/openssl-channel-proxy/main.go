package main

import (
	"flag"
	"log"
	"net"

	proxy "github.com/sanguohot/openssl-channel-proxy"
)

var (
	localAddrPtr  = flag.String("l", ":8000", "local address")
	remoteAddrPtr = flag.String("r", "10.6.250.54:8822", "remote address")
	rcertFile = flag.String("r-cert", "/opt/conf/sdk.crt", "A PEM eoncoded certificate file.")
	rkeyFile  = flag.String("r-key", "/opt/conf/sdk.key", "A PEM encoded private key file.")
	rcaFile   = flag.String("r-ca", "/opt/conf/ca.crt", "A PEM eoncoded ca's certificate file.")
)

func main() {
	flag.Parse()

	laddr, err := net.ResolveTCPAddr("tcp", *localAddrPtr)
	if err != nil {
		log.Fatalf("Failed to resolve local address: %s\n", err)
	}

	raddr, err := net.ResolveTCPAddr("tcp", *remoteAddrPtr)
	if err != nil {
		log.Fatalf("Failed to resolve remote address: %s\n", err)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatalf("Failed to open local port: %s\n", err)
	}

	// Setup a proxy for each incomming connection
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Printf("Failed to accept connection: %s\n", err)
			continue
		}

		log.Printf("New connection from: %s\n", conn.RemoteAddr())
		p := proxy.New(conn, raddr, rkeyFile, rcertFile, rcaFile)
		go p.Start()
	}
}