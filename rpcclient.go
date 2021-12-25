package solanasdk

import (
	"context"
	"errors"

	"github.com/ybbus/jsonrpc/v2"
)

type RpcClient struct {
	Client jsonrpc.RPCClient
}

func NewRpc(url string) *RpcClient {
	Client := jsonrpc.NewClient(url)

	return &RpcClient{Client}
}

// CallComand  - call command rpc
func (r *RpcClient) CallComand(ctx context.Context, action string, params ...interface{}) (*jsonrpc.RPCResponse, error) {
	response, err := r.Client.Call(action, params)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, errors.New(response.Error.Message)
	}
	if response.Result == nil {
		return nil, errors.New("result is null")
	}

	return response, nil
}
