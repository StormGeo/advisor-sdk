package advisorsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func retryReq(
	method string,
	retries uint8,
	delay time.Duration,
	url string,
	body []byte,
) (res *http.Response, err error) {
	for retryNumber := retries + 1; retryNumber > 0; retryNumber-- {
		if method == "GET" {
			res, err = http.Get(url)
		} else if method == "POST" {
			res, err = http.Post(url, "application/json", bytes.NewBuffer(body))
		}

		if res.StatusCode < 500 && res.StatusCode != 429 {
			return res, err
		}

		if retryNumber > 0 {
			time.Sleep(delay)
		}
	}

	return res, err
}

func resToJson(res *http.Response, resError error) (data any, err error) {
	if resError != nil {
		return nil, resError
	}
	defer res.Body.Close()

	body, bodyErr := io.ReadAll(res.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var destiny any

	jsonParserErr := json.Unmarshal(body, &destiny)
	if jsonParserErr != nil {
		return nil, jsonParserErr
	}

	destinyMap, ok := destiny.(map[string]any)
	if ok {
		_, keyExists := destinyMap["error"]
		if keyExists {
			return nil, fmt.Errorf("%v", destiny)
		}
	}

	return destiny, nil
}
