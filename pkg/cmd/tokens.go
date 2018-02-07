package cmd

import (
	"net"

	"git.containerum.net/ch/json-types/auth"
	user "git.containerum.net/ch/json-types/user-manager"
)

const (
	getCheckToken  = "/token/{access_token}"
	getExtendToken = "/token/{access_token}"
	userAgent      = "kube-client"
)

// CheckToken -- consumes JWT token, user fingerprint and IP
// returns user access data: list of namespaces and list of volumes
func (client *Client) CheckToken(token, userFingerprint string, clientIP net.IP) (auth.CheckTokenResponse, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": token,
		}).
		SetResult(auth.CheckTokenResponse{}).
		SetHeaders(map[string]string{
			user.FingerprintHeader: userFingerprint,
			user.ClientIPHeader:    clientIP.String(),
			user.UserAgentHeader:   userAgent,
		}).Get(client.serverURL + getCheckToken)
	if err != nil {
		return auth.CheckTokenResponse{}, err
	}
	return *resp.Result().(*auth.CheckTokenResponse), nil
}

// ExtendToken -- consumes refresh JWT token and user fingerprint
// If they're correct returns new extended access and refresh token OR void tokens AND error.
// Old access and refresh token become inactive.
func (client *Client) ExtendToken(refreshToken, userFingerprint string) (auth.ExtendTokenResponse, error) {
	resp, err := client.Request.
		SetPathParams(map[string]string{
			"access_token": refreshToken,
		}).
		SetResult(auth.ExtendTokenResponse{}).
		SetHeaders(map[string]string{
			user.FingerprintHeader: userFingerprint,
		}).Put(client.serverURL + getExtendToken)
	if err != nil {
		return auth.ExtendTokenResponse{}, err
	}
	return *resp.Result().(*auth.ExtendTokenResponse), nil
}
