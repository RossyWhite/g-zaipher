package zaipher

import (
	"fmt"
	"time"
	"crypto/hmac"
	"encoding/hex"
	"crypto/sha512"
)

type AuthConfig struct {
	APIkey    string
	APISecret string
}

func GetNonce() string {
	return fmt.Sprintf("%.6f", float64(time.Now().UnixNano())/float64(time.Second))
}

func MakeHMAC(msg, key string) string {
	mac := hmac.New(sha512.New, []byte(key))
	mac.Write([]byte(msg))
	return hex.EncodeToString(mac.Sum(nil))
}
