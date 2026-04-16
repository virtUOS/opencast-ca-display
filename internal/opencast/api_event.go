package opencast

import (
	"fmt"
	"time"
)

func (event OpencastEventAPI) GET(ID string) (Event, error) {
	var eventData Event
	err := event.requester.GetJSON(fmt.Sprintf("/api/evets/%s", ID), &eventData)

	if err != nil {
		return Event{}, err
	}

	return eventData, nil
}

func (event OpencastEventAPI) GetAllEvents() ([]Event, error) {
	var eventData []Event

	err := event.requester.GetJSON("/api/events", &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}

func (event OpencastEventAPI) GetEventsByCaptureAgent(agentID string) ([]Event, error) {
	var eventData []Event

	err := event.requester.GetJSON(fmt.Sprintf("/api/events?withscheduling=1&filter=agentid:%s", agentID), &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}

func (event OpencastEventAPI) GetUpcommingEvents(agentID string, duration time.Duration) ([]Event, error) {
	start := time.Now()
	end := start.Add(duration)

	var eventData []Event

	err := event.requester.GetJSON(fmt.Sprintf("/api/events?withscheduling=1&filter=agentid:%s,start:%s/%s", agentID, start.Format(time.RFC3339), end.Format(time.RFC3339)), &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}
