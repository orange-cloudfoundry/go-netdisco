package netdisco

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

type AuthTransport struct {
	ApiKey        string
	WrapTransport http.RoundTripper
}

func NewTransport(apiKey string, insecureSkipVerify bool) *AuthTransport {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	return &AuthTransport{
		ApiKey: apiKey,
		WrapTransport: &http.Transport{
			Proxy:                 http.ProxyFromEnvironment,
			DialContext:           dialer.DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: insecureSkipVerify},
		},
	}
}

func (t AuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t.ApiKey != "" {
		req.Header.Set("Authorization", t.ApiKey)
	}
	return t.WrapTransport.RoundTrip(req)
}
