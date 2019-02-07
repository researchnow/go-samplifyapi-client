package samplify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	log.Printf("Method:%v; URL:%v", method, fmt.Sprintf("%s%s", host, url))
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
		log.Printf("HTTP error:%v; Details:%s", resp.Status, bodyjson)
		return ar, err
	}
	ar.Body = json.RawMessage(bodyjson)
	return ar, err
}

func sendFormData(host, method, path, accessToken string, file multipart.File, fileName string, message string, timeout int) (*APIResponse, error) {
	log.Printf("Method:%v; URL:%v", method, fmt.Sprintf("%s%s", host, path))
	dur := time.Duration(timeout)
	client := &http.Client{
		Timeout: time.Second * dur,
	}
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile("file", fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
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
	//log.Println(bodyBuf)
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
		log.Printf("HTTP error:%v; Details:%s", resp.Status, bodyjson)
		return ar, err
	}
	ar.Body = json.RawMessage(bodyjson)
	return ar, err
}

