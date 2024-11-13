package gosam

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
)

// DialContext implements the net.DialContext function and can be used for http.Transport
func (c *Client) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	errCh := make(chan error, 1)
	connCh := make(chan net.Conn, 1)
	go func() {
		if conn, err := c.DialContextFree(network, addr); err != nil {
			errCh <- err
		} else if ctx.Err() != nil {
			log.Println(ctx)
			errCh <- ctx.Err()
		} else {
			connCh <- conn
		}
	}()
	select {
	case err := <-errCh:
		return nil, err
	case conn := <-connCh:
		return conn, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}

}

func (c *Client) Dial(network, addr string) (net.Conn, error) {
	return c.DialContext(context.TODO(), network, addr)
}

// DialContextFree implements the net.Dial function and can be used for http.Transport
func (c *Client) DialContextFree(network, addr string) (net.Conn, error) {
	if network == "tcp" || network == "tcp6" || network == "tcp4" {
		return c.DialStreamingContextFree(addr)
	}
	if network == "udp" || network == "udp6" || network == "udp4" {
		return c.DialDatagramContextFree(addr)
	}
	if network == "raw" || network == "ip" {
		return c.DialDatagramContextFree(addr)
	}
	return c.DialStreamingContextFree(addr)
}

// DialDatagramContextFree is a "Dialer" for "Client-Like" Datagram connections.
// It is also not finished. If you need datagram support right now, use sam3.
func (c *Client) DialDatagramContextFree(addr string) (*DatagramConn, error) {
	portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, err := c.Lookup(addr)
	if err != nil {
		log.Printf("LOOKUP DIALER ERROR %s %s", addr, err)
		return nil, err
	}

	return nil, fmt.Errorf("Datagram support is not finished yet, come back later`")
}

// DialStreamingContextFree is a "Dialer" for "Client-Like" Streaming connections.
func (c *Client) DialStreamingContextFree(addr string) (net.Conn, error) {
	portIdx := strings.Index(addr, ":")
	if portIdx >= 0 {
		addr = addr[:portIdx]
	}
	addr, err := c.Lookup(addr)
	if err != nil {
		log.Printf("LOOKUP DIALER ERROR %s %s", addr, err)
		return nil, err
	}

	if c.destination == "" {
		c.destination, err = c.CreateStreamSession(c.destination)
		if err != nil {
			return nil, err
		}
	}

	d, err := c.NewClient(c.NewID())
	if err != nil {
		return nil, err
	}
	err = d.StreamConnect(addr)
	if err != nil {
		return nil, err
	}
	return d.SamConn, nil
}
