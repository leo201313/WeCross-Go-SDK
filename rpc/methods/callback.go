package methods

type Callback interface {
	OnSuccess(response any)
	OnFailed(err error)
	CallOnSuccess(response any)
	CallOnFailed(err error)
}
