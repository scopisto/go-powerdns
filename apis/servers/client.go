package servers

import "github.com/scopisto/go-powerdns/pdnshttp"

type client struct {
	httpClient *pdnshttp.Client
}

func New(hc *pdnshttp.Client) Client {
	return &client{
		httpClient: hc,
	}
}
