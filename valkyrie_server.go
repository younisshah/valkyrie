package main

import (
	"fmt"

	"github.com/younisshah/valkyrie/vadapter"
	"github.com/younisshah/valkyrie/vrabbit"
	"github.com/younisshah/valkyrie/vserver"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie Thrift server - Example implementation
 */

const _LISTEN_ADDRESS = "localhost:9090"
const _RABBIT_MQ_URL = "amqp://guest:guest@localhost:5672/"

type valkyrieHandler struct {
	messageQueue vadapter.Queuer
}

func (v valkyrieHandler) Send(message, queueName string) (bool, error) {

	fmt.Println("[+] Rcvd >", message)

	if err := v.messageQueue.Connect(_RABBIT_MQ_URL); err != nil {
		fmt.Println(err)
		return false, err
	} else {
		if err = v.messageQueue.Produce(message, queueName); err != nil {
			fmt.Println("[*] Failed to produce", err)
			return false, err
		}
		fmt.Println("[+] Produced to queue", queueName)
		return true, nil
	}
}

func main() {
	messageQueue := &vrabbit.RabbitMQ{}
	handler := valkyrieHandler{messageQueue: messageQueue}
	valkyrieServer := vserver.NewValkyrieServer("localhost:9090")
	valkyrieServer.InjectValkyrieMessageQueue(messageQueue)
	valkyrieServer.StartServer(handler)
}
