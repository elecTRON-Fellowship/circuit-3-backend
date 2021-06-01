package bcrypt

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPasswd takes a password as input and returns a hashed value
func HashPasswd(passwd string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return string(hashedPass), nil
}

// VerifyPasswd takes a hashed password and a normal password and verify
// its integrity
func VerifyPasswd(hashedPass, passwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(passwd))
}
