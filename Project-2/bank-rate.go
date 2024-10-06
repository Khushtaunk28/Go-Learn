package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Balance.txt"

// read the file data
func getbalanceFromFile() (float64, error) {
	//_ is used when there is and error with file
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to find balance")
	}
	//conv to string
	balanceText := string(data)
	//conv to float
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to convert to float")
	}
	return balance, nil
}

// write the data to file
func writeBalanceToFile(balance float64) {
	balancetext := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balancetext), 0644)
}
func main() {
	var accbal, err = getbalanceFromFile()
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic("Can't continue,Sorry")
	}
	fmt.Println("Welcome to the bank")
	for {
		presentOptions()
		var choice int
		fmt.Print("Enter Your Choice=")
		fmt.Scan(&choice)
		fmt.Println("Your Choice is= ", choice)

		switch choice {
		case 1:
			fmt.Println("your balance is=", accbal)

		case 2:
			{
				var depamt = 0.0
				fmt.Println("Enter the Amount to be deposited=")
				fmt.Scan(&depamt)
				if depamt <= 0 {
					fmt.Println("Invalide amt,must be greater that zero")
					continue
				}
				accbal += depamt
				writeBalanceToFile(accbal)
				fmt.Println("Balance updated!your balance is=", accbal)

			}
		case 3:
			{
				var withdamt = 0.0
				fmt.Println("Enter the Amount to be withdrawn=")
				fmt.Scan(&withdamt)
				if withdamt > accbal {
					fmt.Println("No sufficient balance")
					continue

				} else {
					accbal -= withdamt
					writeBalanceToFile(accbal)
					fmt.Println("Balance updated! your balance is=", accbal)
				}
			}
		default:
			{
				fmt.Println("Goodbye!")
				fmt.Println("thanks for choosing our bank")
				return
			}
		}
	}
}
