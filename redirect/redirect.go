package redirect

type (
	MessageRedirector interface {
		SendMessage([]byte) error
	}
	MessageReceiver interface {
		OnMessage(OnMessage)
	}
	OnMessage func([]byte) error
)
