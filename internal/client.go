package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"reflect"
)

type Client struct {
	BaseUrl    func() string
	Req        func(*http.Request)
	Err        func(resp *http.Response) error
	httpClient *http.Client
}

func NewClient(baseUrl func() string, req func(req *http.Request), err func(resp *http.Response) error) *Client {
	client := new(Client)
	client.BaseUrl = baseUrl
	client.Req = req
	client.Err = err
	client.httpClient = &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}}
	return client
}

func (c *Client) Send(method, endpoint string, data any, out any) error {
	var buf = &bytes.Buffer{}
	if data != nil {
		if err := json.NewEncoder(buf).Encode(data); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, c.BaseUrl()+endpoint, buf)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	c.Req(req)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != 200 {
		return c.Err(resp)
	}

	rv := reflect.ValueOf(out)
	if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.String {
		bs, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		rv.Elem().SetString(string(bs))
		return nil
	}

	return json.NewDecoder(resp.Body).Decode(out)
}
