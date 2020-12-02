package utilities

import (
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

// DoWithRetry executes http.Request and retries in case of 500 range status code
//
func DoWithRetry(client *http.Client, request *http.Request, maxAttempts int, sleepSeconds int32) (*http.Response, *errortools.Error) {
	if client == nil || request == nil {
		return nil, nil
	}

	e := new(errortools.Error)
	e.SetRequest(request)

	attempt := 1
	for attempt <= maxAttempts {
		response, err := client.Do(request)

		if response.StatusCode/100 == 5 { // retry in case of status 500 range (server error)
			attempt++
			fmt.Printf("Starting attempt %v for %s\n", attempt, request.URL.String())
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		} else {
			if err == nil && (response.StatusCode/100 == 4 || response.StatusCode/100 == 5) {
				err = fmt.Errorf("Server returned statuscode %v", response.StatusCode)
			}

			e.SetResponse(response)
			e.SetMessage(err.Error())
			return response, e
		}
	}

	// should never reach this
	return nil, nil
}
