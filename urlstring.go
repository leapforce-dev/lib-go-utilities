package utilities

import (
	"net/url"
)

type URLString struct {
	URL string
}

// RemoveQueryParams removes specified query parameters from query string
//
func (_url *URLString) RemoveQueryParams(params []string) {
	if _url == nil {
		return
	}
	u, _ := url.Parse((*_url).URL)

	query := u.Query()

	for _, p := range params {
		query.Del(p)
	}

	u.RawQuery = query.Encode()

	_url1 := u.String()
	_url.URL = _url1
}
