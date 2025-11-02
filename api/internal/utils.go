package internal

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
)

func Env(key string) string {
	return os.Getenv(key)
}

func HashPassword(key string) string {
	hashedKey, err := bcrypt.GenerateFromPassword([]byte(key), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Error encrypting key: %v\n", err)
		return ""
	}
	return string(hashedKey)
}

func CheckPasswordHash(password, hash string) bool {	
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("No .env file found or error loading .env file")
	}
}