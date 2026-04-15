package opencast

import "net/url"

type OpencastRequester struct {
	URL      url.URL
	Username string
	Password string
	Timeout  int // milliseconds of timeout
}

type OpencastAPI struct {
	requester OpencastRequester

	Events OpencastEventAPI
}

type OpencastEventAPI struct {
	requester *OpencastRequester
}
