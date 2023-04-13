package quakeLog

import (
	"fmt"
	"io/ioutil"
	"net/http"

)

type QuakeLog interface {
	Get() (string, error)
}

type quakeLog struct {
	Url string
}

func NewQuakeLog(url string) *quakeLog {
	return &quakeLog{Url: url}
}

func (ref quakeLog) Get() (string, error) {
	res, err := http.Get(ref.Url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return fmt.Sprint(string(body)), nil
}
