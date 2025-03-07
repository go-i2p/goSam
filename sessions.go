package gosam

import (

	//	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// CreateSession creates a new Session of type style, with an optional destination.
// an empty destination is interpreted as "TRANSIENT"
// Returns the destination for the new Client or an error.
func (c *Client) CreateSession(style, dest string) (string, error) {
	if dest == "" {
		dest = "TRANSIENT"
	}
	//	c.id = id
	r, err := c.sendCmd(
		"SESSION CREATE STYLE=%s ID=%s DESTINATION=%s %s %s %s %s \n",
		style,
		c.ID(),
		dest,
		c.from(),
		c.to(),
		c.sigtype(),
		c.allOptions(),
	)
	if err != nil {
		return "", err
	}

	if !r.IsOk() {
		return "", ReplyError{ResultKeyNotFound, r}
	}
	c.destination = r.Pairs["DESTINATION"]
	return c.destination, nil
}

// CreateStreamSession creates a new STREAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateStreamSession(dest string) (string, error) {
	return c.CreateSession("STREAM", dest)
}

// CreateDatagramSession creates a new DATAGRAM Session.
// Returns the Id for the new Client.
func (c *Client) CreateDatagramSession(dest string) (string, error) {
	return c.CreateSession("DATAGRAM", dest)
}

// CreateRawSession creates a new RAW Session.
// Returns the Id for the new Client.
func (c *Client) CreateRawSession(dest string) (string, error) {
	return c.CreateSession("RAW", dest)
}
