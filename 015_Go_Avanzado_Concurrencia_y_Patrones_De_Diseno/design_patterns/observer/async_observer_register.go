package observer

type AsyncRegisterTopic struct {
	AsyncTopic[string]
}

func NewAsyncRegisterTopic() *AsyncRegisterTopic {
	return &AsyncRegisterTopic{
		NewAsyncTopic[string](),
	}
}
