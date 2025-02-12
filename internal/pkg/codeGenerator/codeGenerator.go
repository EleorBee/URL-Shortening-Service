package codeGenerator

import (
	"math/rand"
)

const base64chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

//
//type codeGenerator struct {
//	storage *sqlite.Storage
//}

func GenerateCode() string {
	code := ""
	for i := 0; i < 6; i++ {
		code += string(base64chars[rand.Intn(len(base64chars))])
	}
	return code
}
