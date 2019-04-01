package factory

import(
	"github.com/streadway/amqp"
	"sync"
)

func GetRabbitConnection() *amqp.Connection {
	var once sync.Once
	var rabbitConn *amqp.Connection

	once.Do(func() {
		rabbitConn, _ = amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	})
	
	return rabbitConn
}