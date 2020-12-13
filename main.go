package main

import (
	"bufio"
	"fmt"
	"os"
	"vending-machine/services"
)

func main() {
	// initialize vending machine
	vm := services.Init()

	fmt.Println("Select role:")
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("1. Seller")
	fmt.Println("2. User")
	fmt.Println("3. Exit")
	fmt.Println("------------------------------------------------------------------------")
	roleChoice := 0
	r := bufio.NewReader(os.Stdin)
	_, err := fmt.Fscanln(r, &roleChoice)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch roleChoice {
	case 1:
		// seller
		fmt.Println("------------------------------------------------------------------------")
		fmt.Println("1. Reset")
		fmt.Println("2. Exit")
		fmt.Println("------------------------------------------------------------------------")
		fmt.Print("Enter Your Choice:")
		choice := 0
		_, err = fmt.Fscanln(r, &choice)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch choice {
		case 1:
			// reset vending machine to default state
			services.Init()
		case 2:
			services.Exit()
		default:
			fmt.Println("Wrong input")
		}
	case 2:
		fmt.Println("------------------------------------------------------------------------")
		// user
		fmt.Println("Select product to purchase:")
		for key, item := range vm.Items {
			fmt.Println(key, "(Price:", item.Price, ")")
		}
		fmt.Println("------------------------------------------------------------------------")
		fmt.Print("Enter Your Choice:")
		productChoice := ""
		_, err = fmt.Fscanln(r, &productChoice)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Selected Item:", productChoice)
		// product choice validate
		_, ok := vm.Items[productChoice]
		if !ok {
			fmt.Println("ITEM_NOT_FOUND")
			return
		}
		fmt.Println("------------------------------------------------------------------------")
		fmt.Println("Acceptable Coins:")
		for _, coin := range vm.Coins {
			fmt.Println(coin.Type)
		}
		fmt.Println("------------------------------------------------------------------------")
		totalCoinsPrice := 0
		returningTotal := 0
		for _, coin := range vm.Coins {
			num := 0
			fmt.Print("How many coins of ", coin.Type, ": ")
			fmt.Fscanln(r, &num)
			totalCoinsPrice += coin.Type * num
			if totalCoinsPrice >= vm.Items[productChoice].Price {
				returningTotal = totalCoinsPrice - vm.Items[productChoice].Price
				break
			}
		}
		fmt.Println("returningTotal: ", returningTotal)
	case 3:
		services.Exit()
	default:
		fmt.Println("Wrong input")
	}
}
