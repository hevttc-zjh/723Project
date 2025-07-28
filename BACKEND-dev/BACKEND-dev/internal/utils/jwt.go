package utils

import (
	"time"

	"risk-insight-system/config"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateJWT 生成JWT token
func GenerateJWT(userId, policeId, role string) (string, error) {
	secret := config.GetString("jwt.secret")
	expireHours := config.GetInt("jwt.expire_hours")

	claims := jwt.MapClaims{
		"user_id":   userId,
		"police_id": policeId,
		"role":      role,
		"exp":       time.Now().Add(time.Duration(expireHours) * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
