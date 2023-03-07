package main

import (
	"fmt"
	"hangman/game"
)

func main() {
	hangman := game.Hangman_new("happy")

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
		fmt.Println("Well done! You've found the hidden word.")
	} else {
		fmt.Printf("Unfortunately, you have not found the hidden word, which was '%s'\n.", hangman.Get_hidden_word())
	}
	fmt.Println("Rerun the program to play again")
}
