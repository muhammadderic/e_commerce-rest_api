package auth

import "golang.org/x/crypto/bcrypt"

func GenerateHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ComparePasswords(passwd string, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwd), plain)
	return err == nil
}
