package util_net

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func ParseQuery(r *http.Request) (url.Values, error) {
	m, err := url.ParseQuery(r.URL.RawQuery)
	return m, err
}

func ExecGet(urlstr string) ([]byte, error) {

	resp, err := http.Get(urlstr)

	var msg []byte

	if err == nil {
		defer resp.Body.Close()
		msg, err = ioutil.ReadAll(resp.Body)
	}

	return msg, err
}
