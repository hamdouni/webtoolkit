package secure

import "golang.org/x/crypto/bcrypt"

// Hash password using Bcrypt
func HashPassword(pw string) (encryptedPassword string, err error) {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hashedPasswordBytes), nil
}

// CheckPassword return true if hash(pw)=hashed
func CheckPassword(pw, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pw))
	return err == nil
}
