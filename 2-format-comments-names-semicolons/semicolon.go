package main

// Similar to C in that it uses semicolons
// But you won't see them... Thanks to the lexer

/*
Rule for adding Semicolons:
	"If the last token before a newline is an identifier (which includes words like int and float64),
	a basic literal such as a number or string constant, or one of the tokens"
	- Effective Go
*/

func semicolon() {
	// Incorrect
	// if 4 < 2
	// {
	// 	// ...
	// }

	// Correct
	if 4 < 2 {
		// ...
	}
}
