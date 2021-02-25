package goSam

import (
	"fmt"
	//	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateSession creates a new STREAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateSession(id int32, style, dest string) (string, error) {
	if dest == "" {
		dest = "TRANSIENT"
	}
	c.id = id
	r, err := c.sendCmd(
		"SESSION CREATE STYLE=%s ID=%d DESTINATION=%s %s %s %s %s \n",
		style,
		c.id,
		dest,
		c.from(),
		c.to(),
		c.sigtype(),
		c.allOptions(),
	)
	if err != nil {
		return "", err
	}

	// TODO: move check into sendCmd()
	if r.Topic != "SESSION" || r.Type != "STATUS" {
		return "", fmt.Errorf("Session Unknown Reply: %+v\n", r)
	}

	result := r.Pairs["RESULT"]
	if result != "OK" {
		return "", ReplyError{ResultKeyNotFound, r}
	}
	c.destination = r.Pairs["DESTINATION"]
	return c.destination, nil
}

// CreateStreamSession creates a new STREAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateStreamSession(id int32, dest string) (string, error) {
	return c.CreateSession(id, "STREAM", dest)
}

// CreateDatagramSession creates a new DATAGRAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateDatagramSession(id int32, dest string) (string, error) {
	return c.CreateSession(id, "DATAGRAM", dest)
}

// CreateRawSession creates a new RAW Session.
// Returns the Id for the new Client.
func (c *Client) CreateRawSession(id int32, dest string) (string, error) {
	return c.CreateSession(id, "RAW", dest)
}
