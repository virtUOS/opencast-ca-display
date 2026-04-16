package endpoints

import (
	"fmt"
	"net/http"
	"opencast-ca-display/internal/opencast"
	"time"

	"github.com/gin-gonic/gin"
)

type Event struct {
	Title string `json:"title"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type CalendarWorkflowProperties struct {
	StraightToPublishing string `json:"straightToPublishing"`
}

type CalenderAgentConfig struct {
	CaptureDeviceNames                 string `json:"capture.device.names"`
	WorkflowDefinition                 string `json:"org.opencastproject.workflow.definition"`
	WorkflowConfigStraightToPublishing string `json:"org.opencastproject.workflow.config.straightToPublishing"`
	EventLocation                      string `json:"event.location"`
	EventTitle                         string `json:"event.title"`
}

type CalendarRecording struct {
}

type CalendarData struct {
	EventID            string                     `json:"eventId"`
	AgentID            string                     `json:"agentId"`
	StartDate          int                        // Verwenden Sie time.Time statt string
	EndDate            int                        // Verwenden Sie time.Time statt string
	Presenters         []string                   `json:"presenters"`
	WorkflowProperties CalendarWorkflowProperties `json:"workflowProperties"`
	AgentConfig        CalenderAgentConfig        `json:"agentConfig"`
	Recording          CalendarRecording          `json:"recording"`
}

type CalendarEntry struct {
	Data              CalendarData `json:"data"`
	EpisodeDublinCore string       `json:"episode-dublincore"`
}

func calendarEndpoint(c *gin.Context) {
	allEvents, err := opencast.Events.GetUpcommingEvents(localConfig.Opencast.Agent, time.Duration(time.Hour*720))
	if err != nil {
		fmt.Printf("Error while getting upcomming events: %s\n", err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": "Internal server error"})
		return
	}

	var events []Event
	for _, eventData := range allEvents {
		start, err := time.Parse(time.RFC3339, eventData.Scheduling.Start)
		if err != nil {
			println("Failed to parse start time")
			return
		}
		end, err := time.Parse(time.RFC3339, eventData.Scheduling.End)
		if err != nil {
			println("Failed to parse end time")
			return
		}
		title := eventData.Title
		e := Event{Title: title, Start: int(start.UnixMilli()), End: int(end.UnixMilli())}
		events = append(events, e)
	}

	if len(events) > 0 {
		c.JSON(http.StatusOK, events)
	} else {
		c.JSON(http.StatusOK, "")
	}
}
