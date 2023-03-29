package middleware

import (
	"crypto/sha1"
	"encoding/hex"
)

func SHA1ForPassword(s string) string {
	o := sha1.New()

	o.Write([]byte(s))

	return hex.EncodeToString(o.Sum(nil))
}
