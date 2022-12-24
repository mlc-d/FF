package hash

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math/rand"
	"time"
	"unsafe"
)

const (
	letterBytes   = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

var src = rand.NewSource(time.Now().UnixNano())

// GenerateRandomString returns a random string of l length
func GenerateRandomString(l int) string {

	// Thanks to @icza for this solution
	// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go

	b := make([]byte, l)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := l-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

// Md5Sum generates a md5sum from src
func Md5Sum(file io.Reader) (md5sum string, err error) {
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	md5sum = hex.EncodeToString(hash.Sum(nil))
	return md5sum, nil
}
