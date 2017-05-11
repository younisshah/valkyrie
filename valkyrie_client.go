package main

import (
	"fmt"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/younisshah/valkyrie/vservice"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie Thrift client
 */

const _URL = "localhost:9090"

func main() {

	transport, err := thrift.NewTSocket(_URL)
	failOnErr(err)

	protocolFactory := thrift.NewTCompactProtocolFactory()
	transport.Open()
	defer transport.Close()

	vclient := vservice.NewValkyrieServiceClientFactory(transport, protocolFactory)
	status, err := vclient.Send("Hello World", "hola")
	failOnErr(err)
	fmt.Println("[+] Send Status:", status)
}

func failOnErr(err error) {
	if err != nil {
		log.Fatal("[*] Failed on error", err)
	}
}
