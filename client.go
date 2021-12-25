package solanasdk

type Client struct {
	client *RpcClient
}

func New(rpcEndpoint string) *Client {
	rpcClient := NewRpc(rpcEndpoint)
	return &Client{
		client: rpcClient,
	}
}
