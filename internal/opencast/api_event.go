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

func (event events) GetCurrentEvent(agentID string) (*Event, error) {
	start := time.Now().Add(-time.Hour * 24) // All future events can start at the earlist at the current time.
	end := time.Now().Add(time.Minute * 10)  // Only events within the next 24 hours are considered

	var eventsList []Event

	query := url.Values{}

	query.Add("withscheduling", "true")
	query.Add("filter", fmt.Sprintf("agent_id:%s,start:%s/%s", agentID, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339)))
	query.Add("sort", "start_date:DESC")
	query.Add("limit", "1")

	err := opencastClient.GetJSONwithQuery("/api/events", query, &eventsList)
	if err != nil {
		return nil, err
	}

	for _, eventData := range eventsList {
		// fmt.Println(eventData)
		eventStart, err := time.Parse(time.RFC3339, eventData.Scheduling.Start)
		if err != nil {
			println("error parsing the date")
			return nil, err
		}
		eventEnd, err := time.Parse(time.RFC3339, eventData.Scheduling.End)
		if err != nil {
			println("error parsing the date")
			return nil, err
		}

		if eventStart.Unix() <= time.Now().Unix() && eventEnd.Unix() >= time.Now().Unix() {
			// there can be only one event which is currently scheduled
			return &eventData, nil
		}
	}

	return nil, nil
}

func (event events) GetNextEvent(agentID string) (*Event, error) {
	start := time.Now()                    // All future events can start at the earlist at the current time.
	end := time.Now().Add(time.Hour * 720) // Only events within the next 24 hours are considered

	var eventsList []Event

	query := url.Values{}

	query.Add("withscheduling", "true")
	query.Add("filter", fmt.Sprintf("agent_id:%s,start:%s/%s", agentID, start.UTC().Format(time.RFC3339), end.UTC().Format(time.RFC3339)))
	query.Add("sort", "start_date:ASC")
	query.Add("limit", "1")

	err := opencastClient.GetJSONwithQuery("/api/events", query, &eventsList)
	if err != nil {
		return nil, err
	}

	for _, eventData := range eventsList {
		// fmt.Println(eventData)
		// eventStart, err := time.Parse(time.RFC3339, eventData.Scheduling.Start)
		// if err != nil {
		// 	println("error parsing the date")
		// 	return false, Event{}, err
		// }
		// eventEnd, err := time.Parse(time.RFC3339, eventData.Scheduling.End)
		// if err != nil {
		// 	println("error parsing the date")
		// 	return false, Event{}, err
		// }

		// if eventStart.Unix() <= time.Now().Unix() && eventEnd.Unix() >= time.Now().Unix() {
		// 	// there can be only one event which is currently scheduled
		// 	return true, eventData, nil
		// }
		return &eventData, nil
	}

	return nil, nil
}
