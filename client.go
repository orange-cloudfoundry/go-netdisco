package netdisco

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Client struct {
	username      string
	password      string
	apiKey        string
	endpoint      string
	httpClient    *http.Client
	authTransport *AuthTransport
	mutexLogin    sync.Mutex
	lastLogin     time.Time
}

func NewClient(endpoint, username, password string, insecureSkipVerify bool) *Client {
	authTransport := NewTransport("", insecureSkipVerify)
	return &Client{
		username:      username,
		password:      password,
		authTransport: authTransport,
		httpClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Transport: authTransport,
			Timeout:   15 * time.Second,
		},
		endpoint: strings.TrimSuffix(endpoint, "/"),
	}
}

func (c *Client) Login() error {
	c.mutexLogin.Lock()
	defer c.mutexLogin.Unlock()
	// if last login below 3600 we already login
	if !c.lastLogin.IsZero() && time.Now().Sub(c.lastLogin) < (time.Second*3600) {
		return nil
	}
	req, err := c.NewRequest(http.MethodPost, "/login", nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.username, c.password)
	apiKeyStruct := struct {
		ApiKey string `json:"api_key"`
	}{}
	c.lastLogin = time.Now()
	err = c.DoWithRetry(req, &apiKeyStruct)
	if err != nil {
		return err
	}
	c.authTransport.ApiKey = apiKeyStruct.ApiKey
	return nil
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.endpoint+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("accept", "application/json")
	return req, nil
}

func (c *Client) Do(method, path string, query, value interface{}) error {
	var body io.Reader
	if query != nil {
		b, err := json.Marshal(query)
		if err != nil {
			return err
		}
		body = bytes.NewBuffer(b)
	}

	req, err := c.NewRequest(method, path, body)
	if err != nil {
		return err
	}

	return c.DoWithRetry(req, value)
}

func (c *Client) DoWithRetry(req *http.Request, value interface{}) error {
	var resp *http.Response
	var respErr error
	for i := 0; i < 3; i++ {
		resp, respErr = c.httpClient.Do(req)
		if respErr != nil {
			continue
		}
		if resp.StatusCode == 302 && req.URL.Path != "/login" {
			err := c.Login()
			if err != nil {
				return err
			}
			return c.DoWithRetry(req, value)
		}
		respErr = c.UnmarshalResponse(resp, value)
		if respErr != nil {
			continue
		}
		break
	}
	if respErr != nil {
		return respErr
	}
	return nil
}

func (c *Client) UnmarshalResponse(resp *http.Response, value interface{}, validResponseCode ...int) error {
	defer resp.Body.Close()
	if len(validResponseCode) == 0 {
		validResponseCode = []int{http.StatusOK}
	}
	valid := false
	for _, respCodeWanted := range validResponseCode {
		if respCodeWanted == resp.StatusCode {
			valid = true
			break
		}
	}
	if !valid {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("response was invalid (%d response code) and body cannot be read: %s", resp.StatusCode, err.Error())
		}
		return fmt.Errorf("response was invalid (%d response code), message: %s", resp.StatusCode, string(b))
	}
	err := json.NewDecoder(resp.Body).Decode(value)
	if err != nil {
		return fmt.Errorf("error during decode response: %s", err.Error())
	}
	return nil
}

func (c *Client) SearchDevice(query *SearchDeviceQuery) ([]Device, error) {
	devices := make([]Device, 0)
	err := c.Do(http.MethodGet, "/api/v1/search/device?"+query.Serialize().Encode(), nil, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
