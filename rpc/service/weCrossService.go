package service

import (
	"WeCross_Go_SDK/rpc/methods"
	"reflect"
)

type WeCrossService interface {
	Init() error
	Send(httpMethod string, uri string, request *methods.Request, p reflect.Type) error
	AsyncSend(httpMethod string, uri string, request *methods.Request, p reflect.Type, callback methods.Callback)
}
