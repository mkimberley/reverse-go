package main

import (
	"fmt"
	"log"
	"os"
)

func ReverseFile(filname string) (string, error) {
	contents, err := os.ReadFile(filname)
	if err != nil {
		return "", err
	}
	return Reverse(string(contents)), err
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		log.Fatal("Filename not specified")
		os.Exit(1)
	}
	filename := arguments[1]
	ReversedText, _ := ReverseFile(filename)
	fmt.Printf("Text reversed: %s\n", ReversedText)

}
