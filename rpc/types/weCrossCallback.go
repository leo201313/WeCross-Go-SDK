package types

type WeCrossCallback struct {
	status  int
	message string
	data    any
}

func (wcc *WeCrossCallback) GetStatus() int {
	return wcc.status
}

func (wcc *WeCrossCallback) SetStatus(status int) {
	wcc.status = status
}

func (wcc *WeCrossCallback) GetMessage() string {
	return wcc.message
}

func (wcc *WeCrossCallback) SetMessage(message string) {
	wcc.message = message
}

func (wcc *WeCrossCallback) GetData() any {
	return wcc.data
}

func (wcc *WeCrossCallback) SetData(data any) {
	wcc.data = data
}
