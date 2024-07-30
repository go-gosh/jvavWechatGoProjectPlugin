package main

import (
	"log/slog"
	"wechat-hub-plugin/hub"
)

type httpResult[T any] struct {
	Code int    `json:"code"` // 0表示成功
	Msg  string `json:"msg"`  //
	Data T      `json:"data"`
}

type Service struct {
	sender *Sender
}

func NewService(sender *Sender) *Service {
	return &Service{
		sender: sender,
	}
}

func (s *Service) Handle(message *hub.Message) error {
	slog.Info("receive message", "type", message.MsgType, "content", message.Content)
	if "#ping3" == message.Content {
		return s.sender.SendText(message.GID, "pong")
	}
	return nil
}
