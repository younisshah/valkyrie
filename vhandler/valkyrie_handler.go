package vhandler

import (
	"fmt"

	"github.com/younisshah/valkyrie/vrabbit"
)

/**
*  Created by Galileo on 12/5/17.
*
*  Valkyrie service handler
 */

type valkyrieHandler struct{}

func NewValkyrieHandler() valkyrieHandler {
	return valkyrieHandler{}
}

func (v valkyrieHandler) Send(message, queueName string) (bool, error) {

	fmt.Println("[+] Rcvd >", message)

	rabbit := vrabbit.NewRabbitMQConfig(queueName)
	if err := rabbit.Produce(message); err != nil {
		fmt.Println("[*] Failed to produce", err)
	} else {
		fmt.Println("[+] Produced to queue", queueName)
	}

	return true, nil
}
