package plugins

import (
	"fmt"
	"log/slog"
	"math/rand"
	"wechat-hub-plugin/hub"
)

type Fake struct {
}

func NewFake() hub.Plugin {
	return &Fake{}
}

func (p *Fake) Handle(message *hub.Message, sender hub.SenderInterface) error {
	if rand.Intn(100) < 3 {
		slog.Info("coming")
		return sender.SendText(message.GID, fmt.Sprintf("【随机抽奖】恭喜%s这个B获得%d个西八币，集齐3000000000000个西八币可兑换1积分。", message.Username, 318*rand.Intn(100)))
	}
	return nil
}
