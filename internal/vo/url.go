package vo

import "strings"

type Url struct {
	url string
}

func NewUrl(url string) Url {
	if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
		return Url{""}
	}
	return Url{
		url: url,
	}
}

func (u Url) GetValue() string {
	return u.url
}
