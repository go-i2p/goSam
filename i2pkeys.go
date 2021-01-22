package goSam

import (
	"errors"
)

// NewDestination generates a new I2P destination, creating the underlying
// public/private keys in the process. The public key can be used to send messages
// to the destination, while the private key can be used to reply to messages
func (c *Client) NewDestination(sigType ...string) (string, string, error) {
	var (
		sigtmp string
	)
	if len(sigType) > 0 {
		sigtmp = sigType[0]
	}
	r, err := c.sendCmd(
		"DEST GENERATE %s\n",
		sigtmp,
	)
	if err != nil {
		return "", err
	}
	var pub, priv string
	if priv = r.Pairs["PRIV"]; priv == "" {
		return "", errors.New("failed to generate private destination key")
	}
	if pub = r.Pairs["PUB"]; pub == "" {
		return priv, errors.New("failed to generate public destination key")
	}
	return priv, pub, nil
}
