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

	fmt.Println("Event published", string(body))

	return nil
}
