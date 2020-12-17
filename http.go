package utilities

import (
	"fmt"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

// DoWithRetry executes http.Request and retries in case of 500 range status code
//
func DoWithRetry(client *http.Client, request *http.Request, maxRetries uint, secondsBetweenRetries uint32) (*http.Response, *errortools.Error) {
	if client == nil || request == nil {
		return nil, nil
	}

	attempt := uint(1)
	maxAttempts := maxRetries + 1

	for attempt <= maxAttempts {
		if attempt > 1 {
			fmt.Printf("Starting attempt %v for %s %s\n", attempt, request.Method, request.URL.String())
			time.Sleep(time.Duration(secondsBetweenRetries) * time.Second)
		}

		response, err := client.Do(request)
		statusCode := 0
		if response != nil {
			statusCode = response.StatusCode
		}

		if statusCode/100 == 5 && attempt < maxAttempts { // retry in case of status 500 range (server error)
			attempt++
		} else {
			if err == nil && (statusCode/100 == 4 || statusCode/100 == 5) {
				err = fmt.Errorf("Server returned statuscode %v", statusCode)
			}

			if err != nil {
				e := new(errortools.Error)
				e.SetRequest(request)
				e.SetResponse(response)
				e.SetMessage(err.Error())

				return response, e
			}

			return response, nil
		}
	}

	// should never reach this
	return nil, nil
}
