package utils

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/pkoukk/tiktoken-go"
)

// 生成随机ID
func Make32ID() string {
	randomBytes := make([]byte, 16)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return "ABCDEFGHIJKLMNOPQRSTUVWXYZ123456"
	}

	return hex.EncodeToString(randomBytes)
}

// Token计算
func GetToken(data string) int {
	tkm, err := tiktoken.EncodingForModel("gpt-3.5-turbo")
	if err != nil {
		return 0
	}

	// encode
	token := tkm.Encode(data, nil, nil)
	return len(token)
}
