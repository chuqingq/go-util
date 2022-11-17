package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
)

func HttpRequestJson(method string, url string, headers map[string]string, req interface{}, resp interface{}) error {
	reqBody, _ := json.Marshal(req)
	request, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("http.NewReuqest error: %v", err)
		return err
	}
	// // content-type
	// if contentType != "" {
	// 	request.Header.Add("Content-Type", contentType)
	// }
	// headers
	for headerKey, headerValue := range headers {
		request.Header.Add(headerKey, headerValue)
	}

	httpResp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("req error: %v", err)
		return err
	}
	defer httpResp.Body.Close()

	respBody, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return err
	}
	log.Printf("respBody: %v", string(respBody))

	err = json.Unmarshal(respBody, &resp)
	if err != nil {
		log.Printf("resp unmarshal error: %v", err)
		return err
	}
	return nil
}

// HttpRequest 发送request
// func SendRequest(url string, body io.Reader, addHeaders map[string]string, method string) ([]byte, error) {
func HttpRequest(method, url string, addHeaders map[string]string, body io.Reader) ([]byte, error) {
	// 1、创建req
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// req.Header.Add("Content-Type", "application/json")

	// 2、设置headers
	for k, v := range addHeaders {
		req.Header.Add(k, v)
	}

	// 3、发送http请求
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		err = errors.New("http status err")
		log.Println(err)
		return nil, err
	}

	// 4、结果读取
	resp, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// log.Println(string(resp))
	return resp, nil
}
