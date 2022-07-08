package utils

import "golang.org/x/crypto/bcrypt"

// CheckPassword 检查密码
func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

// EncryptPassword 加密密码
func EncryptPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err == nil {
		return string(bytes)
	}

	return ""
}
