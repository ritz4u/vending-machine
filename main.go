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
	case 3:
		services.Exit()
	default:
		fmt.Println("Wrong input")
	}
}
