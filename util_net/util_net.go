package util_net

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func GetQueryValues(r *http.Request, keys ...interface{}) ([][]byte, error) {
	m, err := url.ParseQuery(r.URL.RawQuery)

	if err != nil {
		return nil, nil
	}

	var resultArray [][]byte

	for i := 0; i < len(keys); i++ {
		value := m.Get(keys[i].(string))
		resultArray = append(resultArray, []byte(value))
	}
	return resultArray, err
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
