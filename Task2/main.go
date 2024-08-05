package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func WordFrequencyCount(text string) map[string]int {
	text = strings.ToLower(text)

	var cleanedText strings.Builder

	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
			cleanedText.WriteRune(char)
		}
	}

	words := strings.Fields(cleanedText.String())

	frequencyMap := make(map[string]int)

	for _, word := range words {
		frequencyMap[word]++
	}

	return frequencyMap
}

func palindromeChecker(word string) bool {
	word = strings.ToLower(word)

	
	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}

	return true
}


func main() {
	reader := bufio.NewReader(os.Stdin)

	// Word Count
	fmt.Println("Enter a word or sentence to count the frequency of each word:")
	var text string
	text, _ = reader.ReadString('\n')
	fmt.Println(WordFrequencyCount(text))

	fmt.Println()

	// Palindrome Checker
	fmt.Println("Enter a word to check if it is a palindrome:")
	var word string
	word, _ = reader.ReadString('\n')
	word = strings.TrimSpace(word)

	if palindromeChecker(word) {
		fmt.Println("The word is a palindrome")
	} else {
		fmt.Println("The word is not a palindrome")
	}
	

}