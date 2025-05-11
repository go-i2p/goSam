package gosam

// SetupAuth sends the AUTH ENABLE command and immediately sets up a new Username and
// Password from the arguments
func (c *Client) SetupAuth(user, password string) error {
	_, err := c.sendCmd("AUTH ENABLE\n")
	if err != nil {
		return err
	}
	_, err = c.sendCmd("AUTH %s %s\n", user, password)
	if err != nil {
		return err
	}
	return nil
}

// TeardownAuth sends the AUTH DISABLE command but does not remove the Username and
// Password from the client PasswordManager
func (c *Client) TeardownAuth() error {
	_, err := c.sendCmd("AUTH DISABLE\n")
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) RemoveAuthUser(user string) error {
	_, err := c.sendCmd("AUTH REMOVE %s\n", user)
	if err != nil {
		return err
	}
	return nil
}
