package opencast

import "net/url"

var Agents agents

var Events events

type OpencastRequester struct {
	URL      url.URL
	Username string
	Password string
	Timeout  int // milliseconds of timeout
}

type OpencastAPI struct {
	requester OpencastRequester

	Events events
}

type events struct {
}

type agents struct {
}
