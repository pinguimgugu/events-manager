package factory

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func GetRabbitConnection() *amqp.Connection {
	chErr := make(chan *amqp.Error)
	conn := make(chan *amqp.Connection)
	go func() {
		for {
			rabbitConn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")

			if err != nil {
				continue
			}

			conn <- rabbitConn

			rabbitConn.NotifyClose(chErr)
			<-chErr
			fmt.Println("reconnect to rabbit ...")
			time.Sleep(time.Second * 1)
		}
	}()

	return <-conn
}
