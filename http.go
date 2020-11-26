package utilities

import (
	"net/http"
	"time"
)

// DoWithRetry executes http.Request and retries in case of 500 range status code
//
func DoWithRetry(client *http.Client, request *http.Request, maxAttempts int, sleepSeconds int32) (*http.Response, error) {
	if client == nil || request == nil {
		return nil, nil
	}

	attempt := 1
	for attempt <= maxAttempts {
		res, err := client.Do(request)

		if res.StatusCode%100 == 5 { // retry in case of status 500 range (server error)
			attempt++
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		} else {
			return res, err
		}
	}

	// should never reach this
	return nil, nil
}
