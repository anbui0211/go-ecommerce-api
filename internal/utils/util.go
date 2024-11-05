package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetUserKey(hashString string) string {
	return fmt.Sprintf("u:%s:otp", hashString)
}

func GenerateUUID(userId int) string {
	newUUID := uuid.New()
	// covert UUID to string, remove "-" character
	uuidString := strings.ReplaceAll((newUUID).String(), "", "")
	return strconv.Itoa(userId) + "clitoken" + uuidString
}