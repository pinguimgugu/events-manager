package contract

import "github.com/events-manager/domain/entity"

type EventsRepository interface {
	Create(*entity.EventEnvelop) error
	CreateStreaming(eventName *entity.EventStreaming) chan *entity.EventEnvelop
}
