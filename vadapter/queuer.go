package vadapter

/**
*  Created by Galileo on 13/5/17.
*
*  Adapter pattern for Valkyrie message queue
 */

type Queuer interface {
	// Adaptee must provide a *connection* field specific to the Message queue used
	Connect(url string) error
	// Arguments are message and queue name
	Produce(message interface{}, queueName string) error
	// consume from the given queue using the callback function
	Consume(queueName string, callback func(interface{}) error)
	// Closes the connection to message queue using the connection object provided
	Close()
}
