package utils

import "fmt"

func GetUserKey(hashString string) string {
	return fmt.Sprintf("u:%s:otp", hashString)
}