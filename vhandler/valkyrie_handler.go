package vhandler

import (
	"fmt"

	"github.com/younisshah/valkyrie/vadapter"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie service handler
 */

const _RABBIT_MQ_URL = "amqp://guest:guest@localhost:5672/"

type valkyrieHandler struct {
	messageQueue vadapter.Queuer
}

func NewValkyrieHandler(messageBroker vadapter.Queuer) valkyrieHandler {
	return valkyrieHandler{
		messageQueue: messageBroker,
	}
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
