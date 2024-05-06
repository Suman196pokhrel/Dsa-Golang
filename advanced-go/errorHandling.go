package advancedgo

import (
	"errors"
	"fmt"
	"os"
)

// Unlike languages with exceptions, Go doesn’t have special constructs to detect if an
// error was returned. Whenever a function returns, use an if statement to check the
// error variable to see if it is non-nil:

func calcRemainderAndMod(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numerator / denominator, numerator % denominator, nil
}

func ErrorsInGo() {
	numerator := 20
	denominator := 0
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(remainder, mod)
}
