package goSam

/*import (
	"crypto/sha256"
	"errors"
	"strings"
)

type I2PAddr string

// Returns the I2P destination (base32-encoded)
func (a I2PAddr) String() string {
	return string(a.Base32())
}

func (addr I2PAddr) Base32() (str string) {
	return addr.DestHash().String()
}

func (addr I2PAddr) DestHash() (h I2PDestHash) {
	hash := sha256.New()
	b, _ := addr.ToBytes()
	hash.Write(b)
	digest := hash.Sum(nil)
	copy(h[:], digest)
	return
}

// Turns an I2P address to a byte array. The inverse of NewI2PAddrFromBytes().
func (addr I2PAddr) ToBytes() ([]byte, error) {
	return i2pB64enc.DecodeString(string(addr))
}

// Returns "I2P"
func (a I2PAddr) Network() string {
	return "I2P"
}

type I2PDestHash [32]byte

// create a desthash from a string b32.i2p address
func DestHashFromString(str string) (dhash I2PDestHash, err error) {
	if strings.HasSuffix(str, ".b32.i2p") && len(str) == 60 {
		// valid
		_, err = i2pB32enc.Decode(dhash[:], []byte(str[:52]+"===="))
	} else {
		// invalid
		err = errors.New("invalid desthash format")
	}
	return
}

// create a desthash from a []byte array
func DestHashFromBytes(str []byte) (dhash I2PDestHash, err error) {
	if len(str) == 32 {
		// valid
		//_, err = i2pB32enc.Decode(dhash[:], []byte(str[:52]+"===="))
		copy(dhash[:], str)
	} else {
		// invalid
		err = errors.New("invalid desthash format")
	}
	return
}

// get string representation of i2p dest hash(base32 version)
func (h I2PDestHash) String() string {
	b32addr := make([]byte, 56)
	i2pB32enc.Encode(b32addr, h[:])
	return string(b32addr[:52]) + ".b32.i2p"
}

// get base64 representation of i2p dest sha256 hash(the 44-character one)
func (h I2PDestHash) Hash() string {
	hash := sha256.New()
	hash.Write(h[:])
	digest := hash.Sum(nil)
	buf := make([]byte, 44)
	i2pB64enc.Encode(buf, digest)
	return string(buf)
}

// Returns "I2P"
func (h I2PDestHash) Network() string {
	return "I2P"
}

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
		return "", "", err
	}
	var pub, priv string
	if priv = r.Pairs["PRIV"]; priv == "" {
		return "", "", errors.New("failed to generate private destination key")
	}
	if pub = r.Pairs["PUB"]; pub == "" {
		return priv, "", errors.New("failed to generate public destination key")
	}
	return priv, pub, nil
}
*/
