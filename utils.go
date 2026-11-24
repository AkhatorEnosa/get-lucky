package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetName() string {
	name := ""

	fmt.Println("Welcome to Gosab's World!!!")

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your name: ")

	if scanner.Scan() {
		name := scanner.Text()
		fmt.Printf("Hello %s, welcome to my world. Let's play ", name)
	}

	return name
}

func GetBet(balance uint) uint {
	var bet uint
	
	for true {
		fmt.Printf("Enter your bet, or 0 to exit. Your balance is %d: ", balance)
		fmt.Scan(&bet)

		if bet > balance {
			fmt.Println("You cannot bet more than your balance.")
		} else {
			break
		}
	}

	return bet
}
