package cryptx

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func PasswordEncrypt(salt, password string) string{
	dk, _:=scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	return fmt.Sprintf("%x", base64.StdEncoding.EncodeToString(dk))
}