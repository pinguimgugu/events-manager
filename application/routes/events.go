package routes

import (
	"github.com/events-manager/application/resource"
	"github.com/events-manager/domain/contract"

	"github.com/labstack/echo"
)

// Events struct
type Events struct {
	echo             *echo.Echo
	EventsRepository contract.EventsRepository
}

//NewEvents func
func NewEvents(e *echo.Echo, ev contract.EventsRepository) *Events {
	return &Events{e, ev}
}

// Handler
func (e *Events) Handler() {

	e.echo.POST("/events/v1/events/", func(c echo.Context) error {
		action := &resource.EventsPost{
			EventsRepository: e.EventsRepository,
		}

		return action.Handler(c)
	})

	e.echo.POST("/events/v1/streaming/", func(c echo.Context) error {
		action := &resource.EventsStreamingPost{
			EventsRepository: e.EventsRepository,
		}

		return action.Handler(c)
	})
}
