package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Client interface {
	Get(ctx context.Context, url string) (response *http.Response, err error)
	Post(ctx context.Context, url string, body interface{}) (response *http.Response, err error)
	Patch(ctx context.Context, url string, body interface{}) (response *http.Response, err error)
}

type client struct {
	c     *http.Client
	url   string
	token string
}

func (c client) Get(ctx context.Context, url string) (response *http.Response, err error) {
	reqUrl := c.url + url

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	log.Println(http.MethodGet, reqUrl)
	response, err = c.c.Do(req)
	return
}

func (c client) Post(ctx context.Context, url string, body interface{}) (response *http.Response, err error) {
	reqUrl := c.url + url
	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, reqUrl, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	response, err = c.c.Do(req)
	return
}

func (c client) Patch(ctx context.Context, url string, body interface{}) (response *http.Response, err error) {
	reqUrl := c.url + url
	b, err := json.Marshal(body)
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, reqUrl, bytes.NewBuffer(b))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	response, err = c.c.Do(req)
	return
}

func NewClient(token string) Client {
	return &client{
		c:     &http.Client{},
		url:   "https://timecrowd.net/api/v1",
		token: token,
	}
}
