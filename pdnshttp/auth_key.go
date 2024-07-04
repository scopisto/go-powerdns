package pdnshttp

import "net/http"

type APIKeyAuthenticator struct {
	APIKey string
}

func (a *APIKeyAuthenticator) OnRequest(r *http.Request) error {
	r.Header.Set("X-API-Key", a.APIKey)
	r.Header.Set("X-Auth-Token", a.APIKey)
	return nil
}

func (a *APIKeyAuthenticator) OnConnect(*http.Client) error {
	return nil
}
