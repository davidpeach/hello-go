package main

import (
	"os"
)

func main() {
	buffer := []byte("The is the content.\n")
	os.WriteFile("/tmp/lol", buffer, 0644)

	file, _ := os.OpenFile("/tmp/lol", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

	file.WriteString("This comes after.\n")
}
