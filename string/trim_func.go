package main

import (
	"strings"
	"unicode"
)

func main() {

}

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	stars := strings.Repeat("*", numStarsPerLine)
	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage1(oldMsg string) string {

	first := strings.TrimRightFunc(oldMsg, func(r rune) bool {
		return unicode.IsSpace(r) || string(r) == "*"
	})
	second := strings.TrimLeftFunc(first, func(r rune) bool {
		return unicode.IsSpace(r) || string(r) == "*"
	})

	return second
}

func CleanupMessageOther(oldMsg string) string {
	return strings.Trim(oldMsg, "* \n")
}
