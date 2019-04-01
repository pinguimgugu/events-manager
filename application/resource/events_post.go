package resource

import (
	"net/http"

	"github.com/events-manager/domain/contract"
	"github.com/events-manager/domain/entity"
	"github.com/labstack/echo"
)

// EventsPost struct
type EventsPost struct {
	EventsRepository contract.EventsRepository
}

// Handle method
func (e *EventsPost) Handler(c echo.Context) error {
	eventEvelop := new(entity.EventEnvelop)

	if err := c.Bind(eventEvelop); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "invalid event content",
		})
	}

	err := e.EventsRepository.Create(eventEvelop)

	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusAccepted)
}
