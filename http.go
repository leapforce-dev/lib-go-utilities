package utilities

import (
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"time"

	errortools "github.com/leapforce-libraries/go_errortools"
)

// DoWithRetryOLD executes http.Request and retries in case of 500 range status code
//
func DoWithRetryOLD(client *http.Client, request *http.Request, maxRetries uint, secondsBetweenRetries uint32) (*http.Response, *errortools.Error) {
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

// DoWithRetry executes http.Request and retries in case of 500 range status code
// see: https://developers.google.com/analytics/devguides/config/mgmt/v3/errors#handling_500_or_503_responses
func DoWithRetry(client *http.Client, request *http.Request) (*http.Response, *errortools.Error) {
	if client == nil || request == nil {
		return nil, nil
	}

	retry := 0
	maxRetries := 5

	for retry <= maxRetries {
		if retry > 0 {
			fmt.Printf("Starting retry %v for %s %s\n", retry, request.Method, request.URL.String())
			waitSeconds := math.Pow(2, float64(retry-1))
			waitMilliseconds := int(rand.Float64() * 1000)
			time.Sleep(time.Duration(waitSeconds)*time.Second + time.Duration(waitMilliseconds)*time.Millisecond)
		}

		response, err := client.Do(request)
		statusCode := 0
		if response != nil {
			statusCode = response.StatusCode
		}

		if (statusCode == 500 || statusCode == 503) && retry < maxRetries { // retry in case of status 500/503 (server error)
			retry++
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
