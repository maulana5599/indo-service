package utils

func ComparePassword(hashedPassword string, password string) bool {
	return hashedPassword == password
}
