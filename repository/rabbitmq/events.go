package rabbitmq

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/events-manager/domain/entity"
	"github.com/streadway/amqp"
)

type Events struct {
	conn *amqp.Connection
}

func NewEvents(conn *amqp.Connection) *Events {
	return &Events{conn: conn}
}

func (e *Events) Create(eventEnvelop *entity.EventEnvelop) error {
	ch, _ := e.conn.Channel()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		eventEnvelop.Name,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)

	body, _ := json.Marshal(eventEnvelop)
	err = ch.Publish(
		eventEnvelop.Name,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		fmt.Println("error to publisher event", body)
		return errors.New("error to publisher event")
	}

	return nil
}

func (e *Events) CreateStreaming(eventStreaming *entity.EventStreaming) chan *entity.EventEnvelop {
	ch, _ := e.conn.Channel()

	ch.ExchangeDeclare(
		eventStreaming.EventName, // name
		"fanout",                 // type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)

	q, _ := ch.QueueDeclare(
		eventStreaming.ConsumerName, // name
		true,  // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)

	ch.QueueBind(
		q.Name, // queue name
		"",     // routing key
		eventStreaming.EventName, // exchange
		false,
		nil)

	msg, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	chanEvent := make(chan *entity.EventEnvelop)

	go func() {
		for content := range msg {
			marshedMsg := &entity.EventEnvelop{}
			json.Unmarshal(content.Body, &marshedMsg)
			chanEvent <- marshedMsg
		}
	}()

	return chanEvent
}
