package domain

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateUUID() string {
	uuidBytes := make([]byte, 16)
	_, err := rand.Read(uuidBytes)
	if err != nil {
		panic(err)
	}

	uuidString := hex.EncodeToString(uuidBytes)

	return fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		uuidString[:8],
		uuidString[8:12],
		uuidString[12:16],
		uuidString[16:20],
		uuidString[20:],
	)
}
