package gosam

// StreamConnect asks SAM for a TCP-Like connection to dest, has to be called on a new Client
func (c *Client) StreamConnect(dest string) error {
	if dest == "" {
		return nil
	}
	r, err := c.sendCmd("STREAM CONNECT ID=%s DESTINATION=%s %s %s\n", c.ID(), dest, c.from(), c.to())
	if err != nil {
		return err
	}

	if !r.IsOk() {
		return ReplyError{r.GetResult(), r}
	}

	return nil
}

// StreamAccept asks SAM to accept a TCP-Like connection
func (c *Client) StreamAccept() (*Reply, error) {
	r, err := c.sendCmd("STREAM ACCEPT ID=%s SILENT=false\n", c.ID())
	if err != nil {
		return nil, err
	}

	if !r.IsOk() {
		return nil, ReplyError{r.GetResult(), r}
	}

	return r, nil
}
