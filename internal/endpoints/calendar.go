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

	// client := &http.Client{Timeout: time.Duration(localConfig.Timeout * int(time.Millisecond))}
	// // Cutoff is set to 3 day from now; TODO: set back to 24 Hours
	// cutoff := time.Now().Add(time.Hour * 720).UnixMilli()
	// url := localConfig.Opencast.URL + "/recordings/calendar.json?agentid=" + localConfig.Opencast.Agent + "&cutoff=" + fmt.Sprint(cutoff) + "&timestamp=true"
	// req, err := http.NewRequest("GET", url, nil)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusBadGateway, nil)
	// 	return
	// }
	// req.SetBasicAuth(localConfig.Opencast.Username, localConfig.Opencast.Password)
	// resp, err := client.Do(req)
	// if err != nil {
	// 	if os.IsTimeout(err) {
	// 		log.Println("Request timed out:", err)
	// 		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "Request timed out"})
	// 		metrics.UpdateState("request_timed_out")
	// 	} else {
	// 		log.Println(err)
	// 		c.JSON(http.StatusBadGateway, gin.H{"error": "Internal server error"})
	// 		// stateCollector.WithLabelValues("internal_server_error").Set(1)
	// 		metrics.UpdateState("internal_server_error")
	// 	}
	// 	return
	// }
	// if resp.StatusCode != 200 {
	// 	log.Println(resp)
	// 	c.JSON(resp.StatusCode, nil)
	// 	return
	// }

	// bodyText, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	c.JSON(http.StatusBadGateway, nil)
	// 	return
	// }
	// s := string([]byte(bodyText))

	// var allEvents []CalendarEntry
	// json_err := json.Unmarshal([]byte(s), &allEvents)
	// if json_err != nil {
	// 	log.Fatal(json_err)
	// }

	// var events []Event
	// for _, eventData := range allEvents {
	// 	start := eventData.Data.StartDate
	// 	end := eventData.Data.EndDate
	// 	title := eventData.Data.AgentConfig.EventTitle
	// 	e := Event{Title: title, Start: start, End: end}
	// 	events = append(events, e)
	// }

	var events []Event
	for _, eventData := range allEvents {
		fmt.Println(eventData.Scheduling.Start)
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
