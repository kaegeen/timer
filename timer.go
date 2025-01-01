package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Simple Timer")
	fmt.Println("=============")
	fmt.Println("Enter the number of seconds for the timer or type 'exit' to quit.")

	for {
		fmt.Print("Set timer (seconds): ")
		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		var seconds int
		_, err := fmt.Sscanf(input, "%d", &seconds)
		if err != nil || seconds <= 0 {
			fmt.Println("Please enter a valid positive number.")
			continue
		}

		fmt.Printf("Timer set for %d seconds. Starting...\n", seconds)
		for i := seconds; i > 0; i-- {
			fmt.Printf("Time remaining: %d seconds\r", i)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("\nTime's up!\n")
	}
}
