package utils

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func GenerateToken() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(time.Now().String())))
}
