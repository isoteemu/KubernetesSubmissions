package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func printString(s string) {
	timestamp := time.Now().UTC().Format(time.RFC3339Nano)
	fmt.Printf("%s: %s\n", timestamp, s)
}

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	randomID := randString(10)

	printString(randomID)
	for range ticker.C {
		printString(randomID)
	}
}
