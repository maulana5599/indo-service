package utils

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPassword []byte, password []byte) (error, bool) {
	// Bandingkan input password dengan hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return err, false
	} else {
		return nil, true
	}
}
