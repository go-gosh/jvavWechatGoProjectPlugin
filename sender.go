package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"wechat-hub-plugin/hub"
)

type Sender struct {
	apiHost  string
	username string
	password string
	sendFn   func(msg hub.SendMsgCommand) error
	client   *http.Client
}

func NewSender(apiHost string, username string, password string, sendFn func(msg hub.SendMsgCommand) error) *Sender {
	return &Sender{
		apiHost:  apiHost,
		username: username,
		password: password,
		sendFn:   sendFn,
		client:   &http.Client{},
	}
}
func (s *Sender) SendText(gid string, content string) error {
	return s.sendFn(hub.SendMsgCommand{
		Gid:  gid,
		Type: 1,
		Body: content,
	})
}

func (s *Sender) SendNetworkImg(gid string, src string) error {
	return s.sendFn(hub.SendMsgCommand{
		Gid:  gid,
		Type: 2,
		Body: src,
	})
}

func (s *Sender) SendImg(gid string, filename string, file io.Reader) error {
	src, err := s.upload(filename, file)
	if err != nil {
		return err
	}
	return s.sendFn(hub.SendMsgCommand{
		Gid:      gid,
		Type:     2,
		Body:     src,
		Filename: filename,
	})
}

func (s *Sender) upload(filename string, file io.Reader) (string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if part, err := writer.CreateFormFile("file", filename); err != nil {
		return "", err
	} else {
		if _, err = io.Copy(part, file); err != nil {
			return "", err
		}
	}

	if err := writer.WriteField("filename", filename); err != nil {
		return "", err
	}
	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", apiHost+"/upload", body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(username+":"+password)))

	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(resp.Status)
	}
	result := httpResult[string]{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	if result.Code != 0 {
		return "", fmt.Errorf(result.Msg)
	}
	return result.Data, nil

}
