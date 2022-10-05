package methods

import (
	"WeCross_Go_SDK/rpc/types"
)

type RequestData interface {
	GetType() string
	ToString() string
}

type Request struct {
	version         string
	data            RequestData
	ext             any
	weCrossCallback *types.WeCrossCallback
}

func NewRequest(data RequestData) *Request {
	newReq := &Request{
		data: data,
	}
	return newReq
}

func (r *Request) GetVersion() string {
	return r.version
}

func (r *Request) SetVersion(version string) {
	r.version = version
}

func (r *Request) GetData() RequestData {
	return r.data
}

func (r *Request) SetData(data RequestData) {
	r.data = data
}

func (r *Request) GetCallback() *types.WeCrossCallback {
	return r.weCrossCallback
}

func (r *Request) SetCallback(weCrossCallback *types.WeCrossCallback) {
	r.weCrossCallback = weCrossCallback
}

func (r *Request) GetExt() any {
	return r.ext
}

func (r *Request) SetExt(ext any) {
	r.ext = ext
}

func (r *Request) ToString() string {
	res := "Request{" + "version='" + r.version + "'"
	if r.data == nil {
		res += ""
	} else {
		res += ", data=" + r.data.ToString()
	}
	res += "}"
	return res
}
