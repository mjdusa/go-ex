package ext

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

// MD5 returns the MD5 checksum of the data.
func MD5(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// SHA256 returns the SHA256 checksum of the data.
func SHA256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// SHA512 returns the SHA512 checksum of the data.
func SHA512(data string) string {
	hash := sha512.New()
	hash.Write([]byte(data))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
