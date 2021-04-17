package calc

//Configuration stores the neccessary RabbitMQ configuration
type Configuration struct {
	AMQPConnectionURL string
}

//Config variable stores the configuration of the current RabbitMQ
var Config = Configuration{
	AMQPConnectionURL: "amqp://guest:guest@localhost:5672/",
}

//AddTask represents struct form of message
type AddTask struct {
	Number1 int
	Number2 int
}
