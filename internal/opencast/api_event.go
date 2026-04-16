package opencast

import (
	"fmt"
	"net/url"
	"time"
)

func (events) Get(ID string) (Event, error) {
	var eventData Event
	err := opencastClient.GetJSON(fmt.Sprintf("/api/evets/%s", ID), &eventData)

	if err != nil {
		return Event{}, err
	}

	return eventData, nil
}

func (events) GetAllEvents() ([]Event, error) {
	var eventData []Event

	err := opencastClient.GetJSON("/api/events", &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}

func (events) GetEventsByCaptureAgent(agentID string) ([]Event, error) {
	var eventData []Event

	err := opencastClient.GetJSON(fmt.Sprintf("/api/events?withscheduling=1&filter=agentid:%s", agentID), &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}

func (events) GetUpcommingEvents(agentID string, duration time.Duration) ([]Event, error) {
	start := time.Now()
	end := start.Add(duration)

	var eventData []Event

	query := url.Values{}

	query.Add("withscheduling", "true")
	query.Add("filter", fmt.Sprintf("agent_id:%s,start:%s/%s", agentID, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339)))

	err := opencastClient.GetJSONwithQuery("/api/events", query, &eventData)
	if err != nil {
		return nil, err
	}

	return eventData, nil
}
