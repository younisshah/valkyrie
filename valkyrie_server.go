package main

import (
	"github.com/younisshah/valkyrie/vrabbit"
	"github.com/younisshah/valkyrie/vserver"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie Thrift server - Example implementation
 */

const _LISTEN_ADDRESS = "localhost:9090"

func main() {
	messageQueue := &vrabbit.RabbitMQ{}
	valkyrieServer := vserver.NewValkyrieServer("localhost:9090")
	valkyrieServer.InjectValkyrieMessageQueue(messageQueue)
	valkyrieServer.StartServer()
}
