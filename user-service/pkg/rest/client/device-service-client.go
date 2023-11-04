package client

import (
	"io"
	"net/http"
)

func ExecuteNodeE1() ([]byte, error) {
	response, err := http.Get("device-service:4000/ping")

	if err != nil {
		return nil, err
	}

	return io.ReadAll(response.Body)
}
