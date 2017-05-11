package main

import (
	"fmt"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/younisshah/valkyrie/vhandler"
	"github.com/younisshah/valkyrie/vservice"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie Thrift server
 */

const _LISTEN_ADDRESS = "localhost:9090"

func main() {

	handler := vhandler.NewValkyrieHandler()
	processor := vservice.NewValkyrieServiceProcessor(handler)
	transport, err := thrift.NewTServerSocket(_LISTEN_ADDRESS)

	failOnError(err)

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)
	fmt.Println("[+] Serving on:", _LISTEN_ADDRESS)
	err = server.Serve()
	failOnError(err)
}

func failOnError(err error) {
	if err != nil {
		log.Fatal("[*] Failed on error", err)
	}
}
