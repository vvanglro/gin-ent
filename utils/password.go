package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(p string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func DecryptPassword(hashPass, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil)[:])
}
