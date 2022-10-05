package methods

import (
	"WeCross_Go_SDK/common"
	"strconv"
)

type ResponseData interface {
	GetType() string
	ToString() string
}

type Response struct {
	version   string
	errorCode common.ErrorCode
	message   string
	data      ResponseData
}

func NewResponse(errorCode common.ErrorCode, message string, data ResponseData) *Response {
	newRes := &Response{
		errorCode: errorCode,
		message:   message,
		data:      data,
	}
	return newRes
}

func (r *Response) GetVersion() string {
	return r.version
}

func (r *Response) SetVersion(version string) {
	r.version = version
}

func (r *Response) GetErrorCode() common.ErrorCode {
	return r.errorCode
}

func (r *Response) SetErrorCode(errorCode common.ErrorCode) {
	r.errorCode = errorCode
}

func (r *Response) GetMessage() string {
	return r.message
}

func (r *Response) SetMessage(message string) {
	r.message = message
}

func (r *Response) GetData() ResponseData {
	return r.data
}

func (r *Response) SetData(data ResponseData) {
	r.data = data
}

func (r *Response) ToString() string {
	res := "Response{" + "version='" + r.version + "'" + ", errorCode="
	res += strconv.Itoa(int(r.errorCode)) + ", message='" + r.message + "'" + ", data="
	res += r.data.ToString() + "}"
	return res
}
