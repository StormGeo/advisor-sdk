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
	delay uint8,
	url string,
	body []byte,
	header http.Header,
) (res *http.Response, err error) {
	for retryNumber := retries + 1; retryNumber > 0; retryNumber-- {
		var req *http.Request

		if method == "GET" {
			req, err = http.NewRequest("GET", url, nil)
		} else if method == "POST" {
			req, err = http.NewRequest("POST", url, bytes.NewBuffer(body))
		}

		req.Header = header
		client := &http.Client{}

		res, err = client.Do(req)
		if (res.StatusCode < 500) || retryNumber == 0 {
			return res, err
		}

		if res != nil {
			defer res.Body.Close()
		}

		if retryNumber > 0 {
			time.Sleep(time.Second * time.Duration(delay))
		}
	}

	return res, err
}

func formatResponse(res *http.Response, resError error) (data any, err error) {
	if resError != nil {
		return nil, resError
	}

	var destiny any

	if res != nil {
		defer res.Body.Close()

		body, bodyErr := io.ReadAll(res.Body)
		if bodyErr != nil {
			return nil, bodyErr
		}

		if res.Request.Header.Get("Accept") != "application/json" {
			if res.StatusCode > 400 {
				return nil, fmt.Errorf("%s", string(body))
			}

			return string(body), nil
		}

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
	}

	return destiny, nil
}
