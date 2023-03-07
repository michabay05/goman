package game

import (
	"errors"
	"fmt"
	"unicode"
)

type Hangman struct {
	attempts      int
	hidden_word   string
	current_guess []rune
	wrong_letters []rune
}

func Hangman_new(hidden_word string) Hangman {
	h := Hangman{
		attempts:      6,
		hidden_word:   hidden_word,
		current_guess: []rune{},
		wrong_letters: []rune{},
	}
	for i := 0; i < len(hidden_word); i++ {
		h.current_guess = append(h.current_guess, '-')
	}
	return h
}

func (h *Hangman) Update(ltr rune) error {
	if !unicode.IsLetter(ltr) {
		return errors.New("Not a letter")
	}
	found_letter := false
	for idx, val := range h.hidden_word {
		if ltr == val {
			h.current_guess[idx] = val
			found_letter = true
		}
	}
	if !found_letter {
		h.attempts--
		h.wrong_letters = append(h.wrong_letters, ltr)
	}
	return nil
}

func (h Hangman) Check_state() bool {
	for _, val := range h.current_guess {
		if val == '-' {
			return false
		}
	}
	return true
}

func (h Hangman) Print_state() {
	for _, val := range h.current_guess {
		fmt.Printf("%c", val)
	}
	fmt.Println()
}

func (h Hangman) Print_wrong_letters() {
	for _, val := range h.wrong_letters {
		fmt.Printf("%c, ", val)
	}
	fmt.Println()
}

func (h Hangman) Has_attempts() bool {
	return h.attempts > 0
}

func (h Hangman) Get_hidden_word() string {
	return h.hidden_word
}

func (h Hangman) Print_info() {
	fmt.Printf("Attempts left: %d\n", h.attempts)
	fmt.Print("Wrong letters: ")
	h.Print_wrong_letters()
	fmt.Print("Current state: ")
	h.Print_state()
}
