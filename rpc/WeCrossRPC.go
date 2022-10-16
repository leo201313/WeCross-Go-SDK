package rpc

import "WeCross-Go-SDK/rpc/eles/account"

type WeCrossRPC interface {
	Test() *RemoteCall
	SupportedStubs() *RemoteCall
	QueryPub() *RemoteCall
	QueryAuthCode() *RemoteCall
	ListAccount() *RemoteCall
	ListResources(ignoreRemote bool) *RemoteCall
	Detail(path string) *RemoteCall
	Call(path, method string, args ...string) *RemoteCall
	SendTransaction(path, method string, args ...string) *RemoteCall
	Invoke(path, method string, args ...string) *RemoteCall
	CallXA(transactionID, path, method string, args ...string) *RemoteCall
	SendXATransaction(transactionID, path, method string, args ...string) *RemoteCall
	StartXATransaction(transactionID string, paths []string) *RemoteCall
	CommitXATransaction(transactionID string, paths []string) *RemoteCall
	RollbackXATransaction(transactionID string, paths []string) *RemoteCall
	GetXATransaction(transactionID string, paths []string) *RemoteCall
	CustomCommand(command string, path string, args ...any) *RemoteCall
	ListXATransactions(size int) *RemoteCall
	Register(name, password string) *RemoteCall
	Login(name, password string) *RemoteCall
	Logout() *RemoteCall
	AddChainAccount(chainType string, chainAccount *account.ChainAccount) *RemoteCall
	SetDefaultAccount(chainType string, chainAccount *account.ChainAccount, keyID int) *RemoteCall // chainAccount and keyID should be input at least one, and first use the chainAccount

	GetCurrentTransactionID() string
}
