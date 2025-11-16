package main

import (
	"fmt"
	"os"

	_ "github.com/mukailasam/ayika"
)

func main() {
	input := "turtle race"
	secret := os.Getenv("SECRET")
	if input != secret {
		fmt.Println("wrong input, the secet code is ->", secret)
	} else {
		fmt.Println("stww")
	}

}
