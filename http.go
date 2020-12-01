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

	attempt := 1
	for attempt <= maxAttempts {
		res, err := client.Do(request)

		if res.StatusCode/100 == 5 { // retry in case of status 500 range (server error)
			attempt++
			fmt.Printf("Starting attempt %v for %s\n", attempt, request.URL.String())
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		} else {
			if err == nil && (res.StatusCode/100 == 4 || res.StatusCode/100 == 5) {
				err = fmt.Errorf("Server returned statuscode %v", res.StatusCode)
			}

			return res, errortools.ErrorMessage(err)
		}
	}

	// should never reach this
	return nil, nil
}
