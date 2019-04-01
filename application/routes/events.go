package routes

import (
	"fmt"
	"time"

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

	e.echo.GET("/events/v1/streaming/:event_name/:consumer_name/", func(c echo.Context) error {
		fmt.Println(c.Param("event_name"))
		for {
			c.Response().Header().Set("Content-type", "application/json")
			c.Response().Write([]byte("{\"teste\":\"viado\"}\n"))
			c.Response().Flush()
			<-time.After(time.Second * 4)
		}

		return nil
	})
}
