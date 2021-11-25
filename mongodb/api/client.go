package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/sunshineplan/database/mongodb"
)

var _ mongodb.Client = &Client{}

var defaultHTTPClient = http.DefaultClient
var defaultTimeout = time.Minute

func DefaultHTTPClient(client *http.Client) {
	if client == nil {
		panic("client is nil")
	}
	defaultHTTPClient = client
}

type Client struct {
	DataSource string
	Database   string
	Collection string
	AppID      string
	Key        string
}

func (c *Client) SetTimeout(d time.Duration) {
	defaultTimeout = d
}

func (c *Client) Connect() error {
	switch {
	case c.DataSource == "":
		return fmt.Errorf("dataSource is required")
	case c.Database == "":
		return fmt.Errorf("database is required")
	case c.Collection == "":
		return fmt.Errorf("collection is required")
	case c.AppID == "":
		return fmt.Errorf("app ID is required")
	case c.Key == "":
		return fmt.Errorf("api key is required")
	}
	return nil
}

func (c *Client) Close() error { return nil }

type rawData struct {
	DataSource  string           `json:"dataSource"`
	Database    string           `json:"database"`
	Collection  string           `json:"collection"`
	Document    *json.RawMessage `json:"document,omitempty"`
	Documents   *json.RawMessage `json:"documents,omitempty"`
	Filter      *json.RawMessage `json:"filter,omitempty"`
	Update      *json.RawMessage `json:"update,omitempty"`
	Replacement *json.RawMessage `json:"replacement,omitempty"`
	Pipeline    *json.RawMessage `json:"pipeline,omitempty"`
	Projection  *json.RawMessage `json:"projection,omitempty"`
	Sort        *json.RawMessage `json:"sort,omitempty"`
	Limit       int64            `json:"limit,omitempty"`
	Skip        int64            `json:"skip,omitempty"`
	Upsert      bool             `json:"upsert,omitempty"`
}

func (c *Client) body(data interface{}) (io.Reader, error) {
	var raw rawData
	b, _ := json.Marshal(c)
	json.Unmarshal(b, &raw)
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, &raw)
	b, _ = json.Marshal(raw)
	return bytes.NewReader(b), nil
}

func (c *Client) Request(endpoint string, action, data interface{}) error {
	switch {
	case c.DataSource == "":
		return fmt.Errorf("dataSource is required")
	case c.Database == "":
		return fmt.Errorf("database is required")
	case c.Collection == "":
		return fmt.Errorf("collection is required")
	case c.AppID == "":
		return fmt.Errorf("app ID is required")
	case c.Key == "":
		return fmt.Errorf("api key is required")
	}

	body, err := c.body(action)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(base, c.AppID)+endpoint,
		body,
	)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.Key)

	t := time.NewTimer(defaultTimeout)
	var resp *http.Response
	ch := make(chan error, 1)
	go func() {
		resp, err = defaultHTTPClient.Do(req)
		ch <- err
	}()
	select {
	case <-t.C:
		return context.DeadlineExceeded
	case err := <-ch:
		if !t.Stop() {
			<-t.C
		}
		if err != nil {
			return err
		}
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 400:
		return fmt.Errorf("the request was invalid")
	case 401:
		return fmt.Errorf("the request did not include an authorized and enabled Data API Key")
	case 404:
		return fmt.Errorf("the request was sent to an endpoint that does not exist")
	case 500:
		return fmt.Errorf("the Data API encountered an internal error and could not complete the request")
	case 200, 201:

	default:
		return fmt.Errorf("unknown status code: %d", resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, data)
}
