package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	DataSource string
	Database   string
	Collection string
	AppID      string
	Key        string
}

func (c *Client) auth() (m M) {
	b, _ := json.Marshal(
		struct {
			DataSource string `json:"dataSource"`
			Database   string `json:"database"`
			Collection string `json:"collection"`
		}{c.DataSource, c.Database, c.Collection},
	)
	json.Unmarshal(b, &m)
	return
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

	m := c.auth()
	b, err := json.Marshal(action)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(b, &m); err != nil {
		return err
	}
	b, _ = json.Marshal(m)

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(base, c.AppID)+endpoint,
		bytes.NewReader(b),
	)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("api-key", c.Key)

	resp, err := http.DefaultClient.Do(req)
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
	case 200:

	default:
		return fmt.Errorf("unknown status code: %d", resp.StatusCode)
	}

	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, data)
}
