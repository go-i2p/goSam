package gosam

import (
	"fmt"
	"strconv"
	"strings"
)

func validateKindInner(kind string) string {
	if strings.HasPrefix(kind, "SIGNATURE_TYPE=") {
		return kind
	}
	return "SIGNATURE_TYPE=" + kind
}

func validateKind(kind string) (string, error) {
	//convert kind to int
	kint, err := strconv.Atoi(kind)
	if err != nil {
		for _, k := range SAMsigTypes {
			if strings.HasSuffix(k, kind) {
				return validateKindInner(kind), nil
			}
		}
	}
	if kint >= 0 && kint <= 7 {
		return validateKindInner(kind), nil
	}
	return "SIGNATURE_TYPE=7", fmt.Errorf("Invalid sigType: %s", kind)
}

// Generate a new destination and return the base64 encoded string
func (c *Client) NewDestination(kind ...string) (string, string, error) {
	if len(kind) == 0 {
		kind = []string{"7"}
	} else {
		var err error
		kind[0], err = validateKind(kind[0])
		if err != nil {
			if c.debug {
				fmt.Println(err)
			}
		}
	}
	r, err := c.sendCmd("DEST GENERATE %s\n", kind[0])
	if err != nil {
		return "", "", err
	}
	if r.Topic != "DEST" {
		return "", "", fmt.Errorf("NewDestination Unknown Reply: %+v\n", r)
	}
	return r.Pairs["PRIV"], r.Pairs["PUB"], nil

}
