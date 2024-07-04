package cache

import "github.com/scopisto/go-powerdns/pdnshttp"

type client struct {
	httpClient *pdnshttp.Client
}

// New creates a new Cache client
func New(hc *pdnshttp.Client) Client {
	return &client{
		httpClient: hc,
	}
}
