package utilities

import (
	"net/url"
)

const (
	modeAll     string = "all"
	modeInclude string = "include"
	modeExclude string = "exclude"
)

type UrlString struct {
	Url string
}

func (_url *UrlString) RemoveQueryParamsAll() (changed bool) {
	return _url.removeQueryParams(modeAll, []string{})
}

func (_url *UrlString) RemoveQueryParamsExclude(params []string) (changed bool) {
	return _url.removeQueryParams(modeExclude, params)
}

func (_url *UrlString) RemoveQueryParamsInclude(params []string) (changed bool) {
	return _url.removeQueryParams(modeInclude, params)
}

// RemoveQueryParams removes specified query parameters from query string
//
func (_url *UrlString) removeQueryParams(mode string, params []string) (changed bool) {
	if _url == nil {
		return false
	}
	u, _ := url.Parse((*_url).Url)

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
		if _url.Url != _url1 {
			_url.Url = _url1
			return true
		}
	}

	return false
}
