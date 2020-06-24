package service

import (
	"log"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func CreateToken(length int) string {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	tokenBytes := make([]byte, length)
	for i := range tokenBytes {
		tokenBytes[i] = charset[newRand.Intn(len(charset))]
	}
	return string(tokenBytes)
}

func isTokenValid(token string, ident TreeIdent) bool {
	return token == ident.token
}

func CheckIDAndToken2(ok bool, id int32, token string, ident TreeIdent) bool {
	if ok {
		if !isTokenValid(token, ident) {
			log.Printf("Tree with id %v does not correspond to token %v\n", id, token)
			log.Printf("ID: %v, Token: %v\n", ident.id, ident.token)
			return false
		}
		return true
	}
	log.Printf("Tree with id %v does not exist\n", id)
	return false
}
