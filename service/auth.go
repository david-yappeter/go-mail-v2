package service

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func AuthorizationCheck(token string) bool {
	timeNow := time.Now().UTC().Unix()

	decodedString, err := base64.RawStdEncoding.DecodeString(token)

	if err != nil {
		fmt.Println("this error", err)
		return false
	}

	arrStr := strings.Split(string(decodedString), "_")

	if len(arrStr) != 2 {
		fmt.Println(err)
		return false
	}

	timestamp, err := strconv.Atoi(arrStr[0])
	if err != nil || timeNow-int64(timestamp) >= 360 {
		fmt.Println(err)
		return false
	}

	if err = bcrypt.CompareHashAndPassword([]byte(arrStr[1]), []byte(fmt.Sprintf("%d%s%d", timestamp, os.Getenv("SECRET"), timestamp))); err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
