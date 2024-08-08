package plugins

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"wechat-hub-plugin/hub"
)

type Police struct {
	mutex sync.Mutex
	data  map[string]int
}

func NewPolice() hub.Plugin {
	return &Police{
		mutex: sync.Mutex{},
		data:  make(map[string]int),
	}
}

var notify = []string{
	"在休息时间高强度水群，小助手代表群主强烈谴责！",
	"休息时间请保持群内安静，小助手代表群主发出警告！",
	"休息时间的群内活跃度过高，小助手代表群主提出批评！",
	"休息时间请尊重他人休息，小助手代表群主表示不满！",
	"休息时间的群内信息量过大，小助手代表群主发出警告！",
	"休息时间请避免打扰他人，小助手代表群主发出谴责！",
	"休息时间的群内讨论过于频繁，小助手代表群主表示不满！",
	"休息时间请降低群内噪音，小助手代表群主提出批评！",
	"休息时间请减少群内发言，小助手代表群主发出警告！",
	"休息时间的群内信息干扰，小助手代表群主发出谴责！",
	"休息时间请避免群内刷屏，小助手代表群主提出批评！",
	"休息时间的群内交流过于频繁，小助手代表群主表示不满！",
	"休息时间请降低群内信息量，小助手代表群主发出警告！",
	"休息时间的群内互动过于频繁，小助手代表群主发出谴责！",
	"休息时间请减少群内发言频率，小助手代表群主提出批评！",
	"休息时间请降低群内信息干扰，小助手代表群主表示不满！",
	"休息时间的群内互动过于活跃，小助手代表群主发出警告！",
	"休息时间请避免群内信息量过大，小助手代表群主发出谴责！",
}

func (p *Police) Handle(message *hub.Message, sender hub.SenderInterface) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	p.data[message.UID]++
	if p.data[message.UID]%30 == 0 {
		slog.Info("coming")
		return sender.SendText(message.GID, fmt.Sprintf("%s %s", message.Username, notify[rand.Intn(len(notify))]))
	}
	return nil
}
