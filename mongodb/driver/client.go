package driver

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/sunshineplan/database/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var _ mongodb.Client = &Client{}

var defaultTimeout = time.Minute

// Client contains mongodb basic configure.
type Client struct {
	Server     string
	Port       int
	Database   string
	Collection string
	Username   string
	Password   string
	SRV        bool

	client *mongo.Client
	coll   *mongo.Collection
}

func (c *Client) SetTimeout(d time.Duration) {
	defaultTimeout = d
}

// URI returns mongodb uri connection string
func (c *Client) URI() string {
	var prefix, auth, server string
	if c.SRV {
		prefix = "mongodb+srv"
	} else {
		prefix = "mongodb"
	}

	if c.Username != "" && c.Password != "" {
		auth = fmt.Sprintf("%s:%s@", c.Username, c.Password)
	}

	if c.SRV || c.Port == 27017 || c.Port == 0 {
		server = c.Server
	} else {
		server = fmt.Sprintf("%s:%d", c.Server, c.Port)
	}

	return fmt.Sprintf("%s://%s%s/%s", prefix, auth, server, c.Database)
}

// MongoClient opens a mongodb database return *mongo.Client.
func (c *Client) MongoClient() (*mongo.Client, error) {
	if c.client != nil {
		return c.client, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI()))
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	c.client = client

	return client, nil
}

func (c *Client) Connect() error {
	switch {
	case c.Database == "":
		return errors.New("database is required")
	case c.Collection == "":
		return errors.New("collection is required")
	}

	if c.client == nil {
		_, err := c.MongoClient()
		if err != nil {
			return err
		}
	}

	c.coll = c.client.Database(c.Database).Collection(c.Collection)

	return nil
}

func (c *Client) Close() error {
	if c.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
		defer func() {
			cancel()
			c.client = nil
		}()
		return c.client.Disconnect(ctx)
	}
	return nil
}

// Backup backups mongodb database to file.
func (c *Client) Backup(file string) error {
	args := []string{}
	args = append(args, fmt.Sprintf("--uri=%q", c.URI()))
	if c.Collection != "" {
		args = append(args, "-c"+c.Collection)
	}
	args = append(args, "--gzip")
	args = append(args, fmt.Sprintf("--archive=%q", file))

	command := exec.Command("mongodump", args...)
	var stderr bytes.Buffer
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to backup: %s\n%v", stderr.String(), err)
	}
	return nil
}

// Restore restores mongodb database collection from file.
func (c *Client) Restore(file string) error {
	args := []string{}
	args = append(args, fmt.Sprintf("--uri=%q", c.URI()))
	args = append(args, "--gzip")
	args = append(args, "--drop")
	args = append(args, fmt.Sprintf("--archive=%q", file))

	command := exec.Command("mongorestore", args...)
	var stderr bytes.Buffer
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		return fmt.Errorf("failed to restore: %s\n%v", stderr.String(), err)
	}
	return nil
}
