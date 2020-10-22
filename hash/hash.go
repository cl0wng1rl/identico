package hash

import (
	"encoding/hex"
	"strconv"

	"golang.org/x/crypto/sha3"
)

func Hash(input string, length int) string {
	msg := []byte(input)
	out := make([]byte, length)
	c1 := sha3.NewCShake256([]byte("NAME"), []byte("Partition1"))
	c1.Write(msg)
	c1.Read(out)
	return hex.EncodeToString(out)
}

func ConvertHashToBase(hash string, base int) string {
	baseHash := ""
	index := 0
	for len(hash[index:]) >= 16 {
		baseHash += hashBase(hash[index:index+16], base)
		index += 16
	}
	baseHash += hashBase(hash[index:], base)
	return baseHash
}

func hashBase(hash string, base int) string {
	decimal, _ := strconv.ParseInt(hash, 16, 64)
	based := strconv.FormatInt(decimal, base)
	return based
}
