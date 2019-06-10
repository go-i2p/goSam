package samkeys

import (
	"github.com/eyedeekay/sam3/i2pkeys"
)

func DestToKeys(dest string) (i2pkeys.I2PKeys, error) {
	addr, err := i2pkeys.NewI2PAddrFromString(dest)
	if err != nil {
		return i2pkeys.I2PKeys{}, err
	}
	return i2pkeys.NewKeys(addr, dest), nil
}
