package vo

type HashedUrl struct {
	url string
}

func NewHashedUrl(url string) HashedUrl {
	return HashedUrl{
		url: url,
	}
}

func (h HashedUrl) GetValue() string {
	return h.url
}
