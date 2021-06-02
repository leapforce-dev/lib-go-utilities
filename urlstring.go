package utilities

import (
	"net/url"
)

const (
	modeAll     string = "all"
	modeInclude string = "include"
	modeExclude string = "exclude"
)

type URLString struct {
	URL string
}

func (_url *URLString) RemoveQueryParamsAll() (changed bool) {
	return _url.removeQueryParams(modeAll, []string{})
}

func (_url *URLString) RemoveQueryParamsExclude(params []string) (changed bool) {
	return _url.removeQueryParams(modeExclude, params)
}

func (_url *URLString) RemoveQueryParamsInclude(params []string) (changed bool) {
	return _url.removeQueryParams(modeInclude, params)
}

// RemoveQueryParams removes specified query parameters from query string
//
func (_url *URLString) removeQueryParams(mode string, params []string) (changed bool) {
	if _url == nil {
		return false
	}
	u, _ := url.Parse((*_url).URL)

	if u == nil {
		return false
	}

	query := u.Query()

	if len(query) > 0 {
		if mode == modeAll {
			u.RawQuery = ""
		} else if mode == modeExclude {
			for _, p := range params {
				query.Del(p)
			}
			u.RawQuery = query.Encode()
		} else if mode == modeInclude {
			for q := range query {
				remove := true
				for _, p := range params {
					if q == p {
						remove = false
						break
					}
				}
				if remove {
					query.Del(q)
				}
			}
			u.RawQuery = query.Encode()
		}
		_url1 := u.String()
		if _url.URL != _url1 {
			_url.URL = _url1
			return true
		}
	}

	return false
}
