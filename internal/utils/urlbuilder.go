package utils

import "net/url"

type HttpURLBuilder struct {
	raw   string
	u     *url.URL
	args  url.Values
	opcnt int
}

func URLBuild(rawUrl string) *HttpURLBuilder {
	var err error
	var u *url.URL
	u, err = url.Parse(rawUrl)
	if err != nil {
		panic(err)
	}
	return &HttpURLBuilder{
		raw:   rawUrl,
		u:     u,
		args:  url.Values{},
		opcnt: 0,
	}
}

func (u *HttpURLBuilder) Build() string {
	if u.opcnt == 0 {
		return u.raw
	}
	u.u.RawQuery = u.args.Encode()
	return u.u.String()
}

func (u *HttpURLBuilder) BuildUrl() *url.URL {
	u.u.RawQuery = u.args.Encode()
	return u.u
}

func (u *HttpURLBuilder) AddArg(key, value string) *HttpURLBuilder {
	u.opcnt++
	u.args.Add(key, value)
	return u
}

func (u *HttpURLBuilder) SetArg(key, value string) *HttpURLBuilder {
	u.opcnt++
	u.args.Set(key, value)
	return u
}
