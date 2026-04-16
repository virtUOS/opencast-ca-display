// This package provides a basic API interface for the most important public Opencast endpoints.
// There is no logic in here which combines the results from multiple endpoints to enhance readability.

package opencast

import "net/url"

var opencastClient OpencastRequester

func Init(url url.URL, username string, password string, timeout int) {
	opencastClient = OpencastRequester{
		URL:      url,
		Username: username,
		Password: password,
		Timeout:  timeout,
	}
}
