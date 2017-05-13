package vserver

import (
	"fmt"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/younisshah/valkyrie/vadapter"
	"github.com/younisshah/valkyrie/vhandler"
	"github.com/younisshah/valkyrie/vservice"
)

/**
*  Created by Galileo on 13/5/17.
 */

type valkyrieServer struct {
	listenAddress string
	messageQueue  vadapter.Queuer
}

func NewValkyrieServer(url string) *valkyrieServer {
	return &valkyrieServer{listenAddress: url}
}

func (v *valkyrieServer) InjectValkyrieMessageQueue(queuer vadapter.Queuer) {
	v.messageQueue = queuer
}

func (v *valkyrieServer) StartServer() {
	handler := vhandler.NewValkyrieHandler(v.messageQueue)
	processor := vservice.NewValkyrieServiceProcessor(handler)
	transport, err := thrift.NewTServerSocket(v.listenAddress)

	failServerOnErr(err)

	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTCompactProtocolFactory()

	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)
	fmt.Println("[+] Serving on:", v.listenAddress)
	err = server.Serve()
	failServerOnErr(err)
}

func failServerOnErr(err error) {
	if err != nil {
		log.Fatal("[*] Failed on error", err)
	}
}
