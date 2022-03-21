package main

import "fmt"

func main() {
	word := []string{"abc", "car", "ada", "racecar", "cool"}
	fmt.Println(firstPalindrome(word))
}

func firstPalindrome(words []string) string {
	n := len(words)
	for i := 0; i < n; i++ {
		if isPal(words[i]) {
			return words[i]
		}
	}
	return ""
}
func isPal(word string) bool {
	byteWord := []byte(word)
	for i := 0; i < len(word); i++ {
		if byteWord[i] != byteWord[len(word)-i-1] {
			return false
		}
	}
	return true
}
