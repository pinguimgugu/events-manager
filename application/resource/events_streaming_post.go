package resource

import (
	"encoding/json"
	"net/http"

	"github.com/events-manager/domain/contract"
	"github.com/events-manager/domain/entity"
	"github.com/labstack/echo"
)

// EventsStreamingPost struct
type EventsStreamingPost struct {
	EventsRepository contract.EventsRepository
}

// Handle method
func (e *EventsStreamingPost) Handler(c echo.Context) error {
	eventStreaming := new(entity.EventStreaming)

	if err := c.Bind(eventStreaming); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"description": "invalid event straming data",
		})
	}

	interruptStream := make(chan bool)

	go func(interruptStream chan bool) {
		if <-c.Response().CloseNotify() {
			interruptStream <- true
		}
	}(interruptStream)

	events := e.EventsRepository.CreateStreaming(eventStreaming, interruptStream)
	for content := range events {
		eventJson, _ := json.Marshal(content)
		c.Response().Header().Set("Content-type", "application/json")
		c.Response().Write([]byte(string(eventJson) + "\n"))
		c.Response().Flush()
	}

	return nil
}
