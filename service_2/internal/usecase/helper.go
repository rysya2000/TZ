package usecase

import (
	"crypto/md5"
)

func Hash(Salt, Pass string) string {
	passWithSalt := []byte(Pass + Salt)
	b := md5.Sum(passWithSalt)
	return string(b[:])
}
