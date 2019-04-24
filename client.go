/*
 * Copyright (c) 2019. Dabank Authors
 * All rights reserved.
 */

package nuls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"time"
)

const (
	rootPath = "/api"
)

// Client nuls API client
// Clients are safe for concurrent use by multiple goroutines.
type Client struct {
	coon    *http.Client
	baseURL string

	// Verbose when true, talks a lot
	Verbose bool

	// BeforeRequest fires before every request,
	// can be used in rate limit.
	// request is aborted if there is an error
	BeforeRequest func(requestType string, option string, outcome interface{}) error
}

// New initializes a new nuls API client
func New(baseURL string) *Client {
	return &Client{
		coon: &http.Client{
			Timeout: time.Second * 30,
		},
		baseURL: baseURL,
	}
}

// call does almost all the dirty work.
func (c *Client) call(requestMethod, requestType, option string, param map[string]interface{}, outcome interface{}) (err error) {
	// fire hooks if in need
	if c.BeforeRequest != nil {
		err = c.BeforeRequest(requestType, option, outcome)
		if err != nil {
			err = wrapErr(err, "BeforeRequest error")
			return
		}
	}

	// recover if there shall be an panic
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	URL, body := c.craftRequest(requestMethod, requestType, option, param)
	req, err := http.NewRequest(requestMethod, URL, body)
	if err != nil {
		err = wrapErr(err, "http.NewRequest")
		return
	}
	req.Header.Set("User-Agent", "nuls-api-client(Go)")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if c.Verbose {
		var reqDump []byte
		reqDump, err = httputil.DumpRequestOut(req, true)
		if err != nil {
			err = wrapErr(err, "verbose mode req dump failed")
			return
		}

		fmt.Printf("\n%s\n\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.coon.Do(req)
	if err != nil {
		err = wrapErr(err, "sending request")
		return
	}
	defer res.Body.Close()

	if c.Verbose {
		var resDump []byte
		resDump, err = httputil.DumpResponse(res, true)
		if err != nil {
			err = wrapErr(err, "verbose mode res dump failed")
			return
		}

		fmt.Printf("%s\n", resDump)
	}

	var content bytes.Buffer
	if _, err = io.Copy(&content, res.Body); err != nil {
		err = wrapErr(err, "reading response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", res.StatusCode, res.Status, content.String())
		return
	}

	var errRes ErrorResponse
	err = json.Unmarshal(content.Bytes(), &errRes)
	if err != nil {
		err = wrapErr(err, "json unmarshal errRes")
		return
	}
	if !errRes.Success && (errRes.Data.Code != "" || errRes.Data.Msg != "") {
		err = fmt.Errorf("nuls server: [err code %v] %s", errRes.Data.Code, errRes.Data.Msg)
		return
	}

	err = json.Unmarshal(content.Bytes(), outcome)
	if err != nil {
		err = wrapErr(err, "json unmarshal outcome")
		return
	}

	return
}

// craftRequest returns desired URL via param provided
func (c *Client) craftRequest(requestMethod, requestType, option string, param map[string]interface{}) (URL string, body io.Reader) {
	URL = c.baseURL + requestType + option
	if len(param) == 0 {
		return
	}

	if requestMethod == http.MethodPost {
		jsonData, _ := json.Marshal(param)
		body = bytes.NewReader(jsonData)
	}
	return
}
