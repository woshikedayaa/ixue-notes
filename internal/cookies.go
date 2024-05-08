package internal

import "net/http"

type Cookies []*http.Cookie

func (c Cookies) Len() int {
	return len(c)
}

func (c Cookies) Find(key string) *http.Cookie {
	for i := 0; i < len(c); i++ {
		if c[i].Name == key {
			return c[i]
		}
	}
	return nil
}
