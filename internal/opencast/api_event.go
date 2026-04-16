package opencast

import (
	"errors"
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

func (events) GetEventsBeween(agentID string, start time.Time, end time.Time) ([]Event, error) {
	if start.UnixMilli() > end.UnixMilli() {
		return nil, errors.New("The start time can´t be after the end time.")
	}

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

func (event events) GetUpcommingEvents(agentID string, duration time.Duration) ([]Event, error) {
	start := time.Now()
	end := start.Add(duration)

	return event.GetEventsBeween(agentID, start, end)
}

func (event events) GetCurrentEvent(agentID string) (bool, Event, error) {
	start := time.Now().Add(-time.Hour * 24) // Has to be set to another value because otherwise events that are curretnly running are not captured
	end := time.Now().Add(time.Minute * 10)  // the maximum can be now because every started event has started by now (10 minutes are just for a buffer)

	events, err := event.GetEventsBeween(agentID, start, end)
	if err != nil {
		return false, Event{}, err
	}

	for _, eventData := range events {
		eventStart, err := time.Parse(time.RFC3339, eventData.Scheduling.Start)
		if err != nil {
			return false, Event{}, err
		}
		eventEnd, err := time.Parse(time.RFC3339, eventData.Scheduling.End)
		if err != nil {
			return false, Event{}, err
		}

		if eventStart.Unix() <= time.Now().Unix() && eventEnd.Unix() >= time.Now().Unix() {
			// there can be only one event which is currently scheduled
			return true, eventData, nil
		}
	}

	return false, Event{}, nil
}

func (event events) GetNextEvent(agentID string) (bool, Event, error){
	start := time.Now().Add(-time.Hour * 24) // Has to be set to another value because otherwise events that are curretnly running are not captured
	end := time.Now().Add(time.Minute * 10)  // the maximum can be now because every started event has started by now (10 minutes are just for a buffer)

	events, err := event.GetEventsBeween(agentID, start, end)
	if err != nil {
		return false, Event{}, err
	}

	for _, eventData := range events {
		eventStart, err := time.Parse(time.RFC3339, eventData.Scheduling.Start)
		if err != nil {
			return false, Event{}, err
		}
		eventEnd, err := time.Parse(time.RFC3339, eventData.Scheduling.End)
		if err != nil {
			return false, Event{}, err
		}

		if eventStart.Unix() <= time.Now().Unix() && eventEnd.Unix() >= time.Now().Unix() {
			// there can be only one event which is currently scheduled
			return true, eventData, nil
		}
	}

	return false, Event{}, nil
}