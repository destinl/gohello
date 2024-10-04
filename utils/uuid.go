package utils

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func GenerateUUID() {
	// 创建
	u1 := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", u1)

	// 解析
	u2, err := uuid.FromString("5e39d96b-c06b-4016-a90b-64f218797206")
	if err != nil {
		fmt.Printf("Something gone wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)

}
