package rpc

type Client struct {
	rpcURL string
	apiURL string
}

func NewClient(rpcURL, apiURL string) *Client {
	return &Client{
		rpcURL: rpcURL,
		apiURL: apiURL,
	}
}
