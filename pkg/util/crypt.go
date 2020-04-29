package util

import (
	"crypto/sha1"
	"fmt"
)

var (
	SHA1_SALT = "@ur39cjeyx#@76549.x48x#@256xw3."
)

func GSha1(key string) string {
	data := []byte(Concat(key, SHA1_SALT))
	return (fmt.Sprintf("%x", sha1.Sum(data)))
}
