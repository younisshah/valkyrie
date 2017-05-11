# Valkyrie

##### Using Apache Thrift with RabbitMQ message broker!

This is a simple demonstration of using [RabbitMQ message broker](https://rabbitmq.com) 
with [Apache Thrift](https://thrift.apache.org/).

> Thrift client is in Golang using **TCompactProtocol** over **TSocket** transport. 


> Thrift server is in Golang using **TCompactProtocol** for efficient encoding of data.

---

* Client sends the message and the RabbitMQ queue name to publish the data on like this 
 ```go
  status, err := vclient.Send("Hello World", "myqueue")
```
*  If the client skips the queue name the default Valkyrie queue `valkyrie_queue` will be used.

TODO

* Write Valkyrie consumer
* Add support for any kind of data
* Try using Non-Blocking server over **TFramedTransport**. 
 
 
 

 

 
 