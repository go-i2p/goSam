package goSam

import "fmt"

// SetupAuth sends the AUTH ENABLE command and immediately sets up a new Username and
// Password from the arguments
func (c *Client) SetupAuth(user, password string) error {
	r, err := c.sendCmd("AUTH ENABLE\n")
	if err != nil {
		return err
	}
	if r.Topic != "AUTH" {
		return fmt.Errorf("SetupAuth Unknown Reply: %+v\n", r)
	}
	r, err = c.sendCmd("AUTH %s %s\n", user, password)
	if err != nil {
		return err
	}
	if r.Topic != "AUTH" {
		return fmt.Errorf("SetupAuth Unknown Reply: %+v\n", r)
	}
	return nil
}

// TeardownAuth sends the AUTH DISABLE command but does not remove the Username and
// Password from the client PasswordManager
func (c *Client) TeardownAuth() error {
	r, err := c.sendCmd("AUTH DISABLE\n")
	if err != nil {
		return err
	}
	if r.Topic != "AUTH" {
		return fmt.Errorf("TeardownAuth Unknown Reply: %+v\n", r)
	}
	return nil
}

func (c *Client) RemoveAuthUser(user string) error {
	r, err := c.sendCmd("AUTH REMOVE %s\n", user)
	if err != nil {
		return err
	}
	if r.Topic != "AUTH" {
		return fmt.Errorf("RemoveAuthUser Unknown Reply: %+v\n", r)
	}
	return nil
}
