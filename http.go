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
	for attempt <= maxRetries+1 {
		response, err := client.Do(request)

		if response.StatusCode/100 == 5 { // retry in case of status 500 range (server error)
			attempt++
			fmt.Printf("Starting attempt %v for %s %s\n", attempt, request.Method, request.URL.String())
			time.Sleep(time.Duration(secondsBetweenRetries) * time.Second)
		} else {
			if err == nil && (response.StatusCode/100 == 4 || response.StatusCode/100 == 5) {
				err = fmt.Errorf("Server returned statuscode %v", response.StatusCode)
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
