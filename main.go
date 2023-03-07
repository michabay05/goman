package main

import (
	"fmt"
	"hangman/game"
	"io/ioutil"
	"log"
	"math/rand"
	"regexp"
)

func get_rand_word(path string) string {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal()
	}
	words := string(content)
	regex := regexp.MustCompile(`\s`)
	split := regex.Split(words, -1)
	word_set := []string{}

	for i := range split {
		word_set = append(word_set, split[i])
	}
	rand_ind := rand.Intn(len(word_set))
	return word_set[rand_ind]
}

func main() {
	hidden_word := get_rand_word("words.txt")
	hangman := game.Hangman_new(hidden_word)
	fmt.Println(hidden_word)

	found_word := false
	for hangman.Has_attempts() {
		if hangman.Check_state() {
			found_word = true
			break
		}
		var input string
		hangman.Print_info()
		fmt.Print("Input: ")
		fmt.Scanln(&input)
		if input == "$q" {
			break
		}
		hangman.Update([]rune(input)[0])
		fmt.Println()
	}
	if found_word {
		fmt.Printf("Well done! You've found the hidden word - '%s'.\n", hangman.Get_hidden_word())
	} else {
		fmt.Printf("Unfortunately, you have not found the hidden word, which was '%s'\n.", hangman.Get_hidden_word())
	}
	fmt.Println("Rerun the program to play again")
}
