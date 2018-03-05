package client

import (
	"git.containerum.net/ch/kube-client/pkg/rest"

	"git.containerum.net/ch/kube-client/pkg/model"
)

const (
	getCheckToken  = "/token/{access_token}"
	getExtendToken = "/token/{refresh_token}"
	userAgent      = "kube-client"
)

// CheckToken -- consumes JWT token, user fingerprint
// If they're correct returns user access data:
// list of namespaces and list of volumes OR uninitialized structure AND error
func (client *Client) CheckToken(token string) (model.CheckTokenResponse, error) {
	var tokenResponse model.CheckTokenResponse
	err := client.re.Get(rest.Rq{
		Result: &tokenResponse,
		Path: rest.URL{
			Templ: client.APIurl + getCheckToken,
			Params: rest.P{
				"access_token": token,
			},
		},
	})
	return tokenResponse, err
}

// ExtendToken -- consumes refresh JWT token and user fingerprint
// If they're correct returns new extended access and refresh token OR void tokens AND error.
// Old access and refresh token become inactive.
func (client *Client) ExtendToken(refreshToken string) (model.Tokens, error) {
	var tokens model.Tokens
	err := client.re.Put(rest.Rq{
		Result: &tokens,
		Path: rest.URL{
			Templ: client.APIurl + getExtendToken,
			Params: rest.P{
				"refresh_token": refreshToken,
			},
		},
	})
	return tokens, err
}
