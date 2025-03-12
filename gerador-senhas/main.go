package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"

func generatePassword(length int) string {
	password := make([]byte, length)
	for i := range password {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		password[i] = charset[num.Int64()]
	}
	return string(password)
}

func main() {
	var length int
	fmt.Print("Digite o tamanho da senha: ")
	fmt.Scan(&length)

	if length <= 0 {
		fmt.Println("O tamanho precisa ser maior que zero!")
		return
	}

	password := generatePassword(length)
	fmt.Println("Senha gerada:", password)
}
