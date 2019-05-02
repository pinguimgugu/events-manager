package contract

import "github.com/events-manager/domain/entity"

type EventsRepository interface {
	Create(*entity.EventEnvelop) error
	CreateStreaming(*entity.EventStreaming, chan bool) chan *entity.EventEnvelop
}
