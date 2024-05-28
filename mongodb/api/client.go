package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/sunshineplan/database/mongodb"
)

var _ mongodb.Client = &Client{}

const defaultVersion = "v1"

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
	Version    string
}

func (c *Client) SetTimeout(d time.Duration) {
	defaultTimeout = d
}

func (c *Client) Connect() error {
	switch {
	case c.DataSource == "":
		return fmt.Errorf("DataSource is required")
	case c.Database == "":
		return fmt.Errorf("database is required")
	case c.Collection == "":
		return fmt.Errorf("collection is required")
	case c.AppID == "":
		return fmt.Errorf("AppID is required")
	case c.Key == "":
		return fmt.Errorf("API key is required")
	case c.Version == "":
		c.Version = defaultVersion
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

func (c *Client) body(data any) (io.Reader, error) {
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

func (c *Client) Request(endpoint string, action, data any) error {
	if rv := reflect.ValueOf(data); rv.Kind() != reflect.Pointer {
		return &mongodb.InvalidDecodeError{Type: reflect.TypeOf(data)}
	}

	if err := c.Connect(); err != nil {
		return err
	}

	body, err := c.body(action)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		"POST",
		fmt.Sprintf(base, c.AppID, c.Version)+endpoint,
		body,
	)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.Key)

	resp, err := defaultHTTPClient.Do(req)
	if err != nil {
		return err
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

	return json.NewDecoder(resp.Body).Decode(data)
}
