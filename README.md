# Valkyrie

##### An Apache Thrift server with pluggable message broker!

Valkyrie implements Adapter pattern to plug-in any message broker you want to process the messages.
The current implementation provides an example implementation of **RabbitMQ**.

You should provide an implementation of `Queuer` and `ValkyrieService` interface.
 
 ```go
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
```

```go
    type ValkyrieService interface {
        // Parameters:
        //  - Message
        //  - Queue
        Send(message string, queue string) (r bool, err error)
    }
```

After writing your custom message queue implementation, use it to construct the **Apache Thrift** `handlder`

Example __RabbitMQ__ message broker handler:

```go
handler := vhandler.NewValkyrieHandler(&vrabbit.RabbitMQ{})
```
Finally implement the `Send` method of `ValkyrieService`.

Example implementation:

```go
    if err := v.messageBroker.Connect(RABBIT_MQ_URL); err != nil {
    		fmt.Println(err)
    		return false, err
    	} else {
    		if err = v.messageBroker.Produce(message, queueName); err != nil {
    			fmt.Println("[*] Failed to produce", err)
    			return false, err
    		}
    		fmt.Println("[+] Produced to queue", queueName)
    		return true, nil
    }
```


---

Example server implementation:

```go
messageQueue := &vrabbit.RabbitMQ{}
valkyrieServer := vserver.NewValkyrieServer(_LISTEN_ADDRESS)
valkyrieServer.InjectValkyrieMessageQueue(messageQueue)
valkyrieServer.StartServer()
```

---


TODO

* Write Valkyrie consumer
* Add support for any kind of data
* Try using Non-Blocking server over **TFramedTransport**. 
 
 
 

 

 
 