package rpc

type Client struct {
	rpcJsonURL string
	apiURL     string
}

func NewClient(rpcJsonURL, apiURL string) *Client {
	return &Client{
		rpcJsonURL: rpcJsonURL,
		apiURL:     apiURL,
	}
}
