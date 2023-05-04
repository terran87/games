package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	low, high := 0, 100

	fmt.Println("Guess a number between " + strconv.Itoa(low) + " and " + strconv.Itoa(high))

	for {
		guess := (low + high) / 2
		fmt.Println("is your number higher or lower than " + strconv.Itoa(guess) + "?")
		fmt.Println("[a] - Higher")
		fmt.Println("[b] - Lower")
		fmt.Println("[c] - Correct")

		scan := bufio.NewReader(os.Stdin)
		char, _, err := scan.ReadRune()
		if err != nil {
			fmt.Println(err)
		}

		if char == 97 {
			low = guess + 1
		} else if char == 98 {
			high = guess - 1
		} else if char == 99 {
			fmt.Println("Winner winner, chicken dinner!")
			break
		} else {
			fmt.Println("try again")
		}
	}
}
