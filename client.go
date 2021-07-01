package mycommerce_sdk

import (
	"encoding/json"
	"github.com/codingXiang/gocron-sdk/auth"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"net/http"
)

type Client struct {
	Auth    *auth.Jwt
	baseUrl string
	//*http.Client
	*fasthttp.Client
}

func NewClient(config *viper.Viper) *Client {
	c := &Client{
		baseUrl: config.GetString("baseUrl"),
	}
	if _auth := config.GetStringMap("auth"); _auth != nil {
		c.Auth = auth.NewJwt(_auth["id"].(int), _auth["name"].(string), _auth["secret"].(string))
	}
	c.Client = &fasthttp.Client{}
	//c.Client = &http.Client{}
	return c
}
func (c *Client) request(endpoint string, body interface{}, contentType ...string) (*fasthttp.Request, error) {
	req := fasthttp.AcquireRequest()

	if c.Auth != nil {
		token, err := c.Auth.Get()
		if err != nil {
			return nil, err
		}
		req.Header.Add(auth.Authorization, token)
	}
	if contentType != nil {
		req.Header.SetContentType(contentType[0])
		req.SetBody(body.([]uint8))
	} else {
		req.Header.SetContentType("application/json")
		if body != nil {
			in, err := json.Marshal(body)
			if err != nil {
				return nil, err
			}
			req.SetBody(in)
		}
	}
	req.SetRequestURI(c.baseUrl + endpoint)

	return req, nil
}

func (c *Client) Get(endpoint string) (*fasthttp.Response, error) {
	req, err := c.request(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release

	resp := fasthttp.AcquireResponse()
	req.Header.SetMethod(http.MethodGet)
	err = c.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Post(endpoint string, body interface{}, contentType ...string) (*fasthttp.Response, error) {
	req, err := c.request(endpoint, body, contentType...)
	if err != nil {
		return nil, err
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	req.Header.SetMethod(http.MethodPost)
	err = c.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Put(endpoint string, body interface{}) (*fasthttp.Response, error) {
	req, err := c.request(endpoint, body)
	if err != nil {
		return nil, err
	}
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	req.Header.SetMethod(http.MethodPut)
	err = c.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

func (c *Client) Delete(endpoint string) (*fasthttp.Response, error) {
	req, err := c.request(endpoint, nil)
	if err != nil {
		return nil, err
	}
	defer fasthttp.ReleaseRequest(req) // <- do not forget to release
	resp := fasthttp.AcquireResponse()
	req.Header.SetMethod(http.MethodDelete)
	err = c.Do(req, resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
