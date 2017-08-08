package lagou

type Msg struct {
	data      interface{}
	error     error
	interrupt bool
}

func MsgData(data interface{}) *Msg {
	return &Msg{
		data: data,
	}
}

func MsgError(err error) *Msg {
	return &Msg{
		error: err,
	}
}

func MsgInterrupt() *Msg {
	return &Msg{
		interrupt: true,
	}
}

func (m Msg) HasData() bool {
	return m.data != nil
}

func (m Msg) HasError() bool {
	return m.error != nil
}

func (m Msg) IsInterrupted() bool {
	return m.interrupt
}
