package advisorsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func get(url string) (res AdvisorResponse, err error) {
	getResp, getErr := http.Get(url)
	if getErr != nil {
		return nil, getErr
	}
	defer getResp.Body.Close()

	body, bodyErr := io.ReadAll(getResp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var bodyToJson map[string]interface{}
	jsonParserErr := json.Unmarshal(body, &bodyToJson)
	if jsonParserErr != nil {
		return nil, jsonParserErr
	}

	_, keyExists := bodyToJson["error"]
	if keyExists {
		return nil, fmt.Errorf("%v", bodyToJson)
	}

	return bodyToJson, nil
}

func getImage(url string) (imageBody io.ReadCloser, err error) {
	getResp, getErr := http.Get(url)
	if getErr != nil {
		return nil, getErr
	}

	return getResp.Body, nil
}

func post(url string, body []byte) (res AdvisorResponse, err error) {
	postResp, postErr := http.Post(url, "application/json", bytes.NewBuffer(body))
	if postErr != nil {
		return nil, postErr
	}
	defer postResp.Body.Close()

	body, bodyErr := io.ReadAll(postResp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	var bodyToJson map[string]interface{}
	jsonParserErr := json.Unmarshal(body, &bodyToJson)
	if jsonParserErr != nil {
		return nil, jsonParserErr
	}

	return bodyToJson, nil
}

func postGeometry(url string, body []byte) (res AdvisorResponse, err error) {
	postResp, postErr := http.Post(url, "application/json", bytes.NewBuffer(body))
	if postErr != nil {
		return nil, postErr
	}
	defer postResp.Body.Close()

	body, bodyErr := io.ReadAll(postResp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	if postResp.StatusCode == 200 {
		var bodyToJson []interface{}

		jsonParserErr := json.Unmarshal(body, &bodyToJson)
		if jsonParserErr != nil {
			return nil, jsonParserErr
		}

		return bodyToJson, nil
	}

	var errorToJson map[string]interface{}
	jsonParserErr := json.Unmarshal(body, &errorToJson)
	if jsonParserErr != nil {
		return nil, jsonParserErr
	}

	return nil, fmt.Errorf("%v", errorToJson)
}

func retryRequest(
	funcName string,
	retries uint8,
	delay time.Duration,
	url string,
	body []byte,
) (res interface{}, err error) {
	for attempts := retries + 1; attempts > 0; attempts-- {
		switch funcName {
		case "get":
			res, err = get(url)
		case "post":
			res, err = post(url, body)
		case "postGeometry":
			res, err = postGeometry(url, body)
		}

		if err == nil {
			break
		}

		if attempts > 1 {
			time.Sleep(delay)
		}
	}

	return res, err
}

func retryGetImage(retries uint8, delay time.Duration, url string) (res io.ReadCloser, err error) {
	for attempts := retries + 1; attempts > 0; attempts-- {
		res, err = getImage(url)

		if err == nil {
			break
		}

		if attempts > 1 {
			time.Sleep(delay)
		}
	}

	return res, err
}
