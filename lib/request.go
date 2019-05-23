package samplify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"time"
)

// APIResponse ...
type APIResponse struct {
	Body      json.RawMessage
	RequestID string
}

func sendRequest(host, method, url, accessToken string, body interface{}, timeout int) (*APIResponse, error) {
	// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendRequest", "URL": fmt.Sprintf("%s%s", host, url), "Method": method}).Info()
	jstr, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	dur := time.Duration(timeout)
	client := &http.Client{
		Timeout: time.Second * dur,
	}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, url), bytes.NewBuffer(jstr))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-type", "application/json")
	if len(accessToken) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyjson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ar := &APIResponse{
		RequestID: resp.Header.Get("x-request-id"),
	}
	if resp.StatusCode >= http.StatusBadRequest {
		t := time.Now()
		errPath := fmt.Sprintf("%s%s", host, url)
		err := &ErrorResponse{
			Timestamp:  &t,
			RequestID:  ar.RequestID,
			HTTPCode:   resp.StatusCode,
			HTTPPhrase: resp.Status,
			Path:       errPath,
			Errors:     []*Error{&Error{Path: errPath, Message: resp.Status}},
		}
		ar.Body = json.RawMessage(bodyjson)
		// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendRequest", "respBody": ar.Body}).Info(err)
		// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendRequest"}).Info(err)
		return ar, err
	}
	ar.Body = json.RawMessage(bodyjson)
	return ar, err
}

func sendFormData(host, method, path, accessToken string, file multipart.File, fileName string, message string, timeout int) (*APIResponse, error) {
	// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendFormData", "URL": fmt.Sprintf("%s%s", host, path), "Method": method}).Info()
	dur := time.Duration(timeout)
	client := &http.Client{
		Timeout: time.Second * dur,
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendFormData", "Method": method}).Error(err)
		return nil, err
	}
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		return nil, err
	}
	bodyWriter.WriteField("message", message)
	bodyWriter.Close()
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", host, path), bodyBuf)
	if len(accessToken) > 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	}
	req.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	bodyjson, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	ar := &APIResponse{
		RequestID: resp.Header.Get("x-request-id"),
	}
	if resp.StatusCode >= http.StatusBadRequest {
		t := time.Now()
		errPath := fmt.Sprintf("%s%s", host, path)
		err := &ErrorResponse{
			Timestamp:  &t,
			RequestID:  ar.RequestID,
			HTTPCode:   resp.StatusCode,
			HTTPPhrase: resp.Status,
			Path:       errPath,
			Errors:     []*Error{&Error{Path: errPath, Message: resp.Status}},
		}
		ar.Body = json.RawMessage(bodyjson)
		// log.WithFields(log.Fields{"module": "go-samplifyapi-client", "function": "sendFormData"}).Error(err)
		return ar, err
	}
	ar.Body = json.RawMessage(bodyjson)
	return ar, err
}
