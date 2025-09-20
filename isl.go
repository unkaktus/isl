package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	tls "github.com/driftnet-io/insecure-tls"
	"github.com/unkaktus/tlspin"
)

const tlspinKey = "InsecureSocketLayer="

func run() error {
	var targetAddr = flag.String("t", "", "Target address")
	var listenAddr = flag.String("l", ":443", "Listen address")
	flag.Parse()

	if *targetAddr == "" {
		log.Fatal("target address is not specified")
	}

	listener, err := tlspin.Listen("tcp", *listenAddr, tlspinKey)
	if err != nil {
		return fmt.Errorf("tlspin listen: %w", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept: %v", err)
			break
		}
		defer conn.Close()

		config := &tls.Config{
			InsecureSkipVerify: true,
		}

		go func() {
			targetConn, err := tls.Dial("tcp", *targetAddr, config)
			if err != nil {
				log.Printf("dial TLS: %v", err)
				return
			}

			go io.Copy(conn, targetConn)
			go io.Copy(targetConn, conn)
		}()
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
