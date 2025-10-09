package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DoPostRequest(
	url string,
	data map[string]string,
	header map[string]string,
) (string, error) {
	var req *http.Request
	var err error

	contentType, _ := header["Content-Type"]
	switch contentType {
	case "application/json":
		jsonData, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	default:
		return "", nil
	}

	for k, v := range header {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("read response body error: %w", err)
	}

	return string(result), err
}
