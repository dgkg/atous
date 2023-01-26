package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func Password(pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	return hex.EncodeToString(hash.Sum(nil))
}
