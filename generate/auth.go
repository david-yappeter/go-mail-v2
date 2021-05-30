package generate

import (
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateAuth(key string) (string, error) {
	timeNow := time.Now().UTC().Unix()

	hash, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%d%s%d", timeNow, key, timeNow)), 10)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return base64.RawStdEncoding.EncodeToString([]byte(fmt.Sprintf("%d_%s", timeNow, string(hash)))), nil
}
