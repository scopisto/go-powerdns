package search

import "github.com/scopisto/go-powerdns/pdnshttp"

type client struct {
	httpClient *pdnshttp.Client
}

// New creates a new Search client
func New(hc *pdnshttp.Client) Client {
	return &client{
		httpClient: hc,
	}
}
