package lagou

// message defines the multi-channel communication message
type message struct {
	data      interface{}
	error     error
	interrupt bool
}

func newMsgData(data interface{}) *message {
	return &message{
		data: data,
	}
}

func newMsgError(err error) *message {
	return &message{
		error: err,
	}
}

//func newMsgInterrupt() *message {
//	return &message{
//		interrupt: true,
//	}
//}

func (m message) hasData() bool {
	return m.data != nil
}

func (m message) hasError() bool {
	return m.error != nil
}

func (m message) isInterrupted() bool {
	return m.interrupt
}
