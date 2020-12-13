package services

import (
	"fmt"
	"os"
	"vending-machine/models"
)

var vm = models.VendingMachine{}

// Exit exit the application
func Exit() {
	os.Exit(0)
}

// Init initialize the vending machine
func Init() *models.VendingMachine {
	vm.Available = true
	vm.Coins = getCoins()
	vm.Items = getItems()
	vm.UserBalance = 1000
	return &vm
}

func getCoins() map[string]*models.Coin {
	coinMap := make(map[string]*models.Coin)
	coinMap["1"] = &models.Coin{1, 1}
	coinMap["5"] = &models.Coin{2, 5}
	coinMap["10"] = &models.Coin{3, 10}
	coinMap["25"] = &models.Coin{4, 25}
	return coinMap
}

func getItems() map[string]*models.Item {
	itemMap := make(map[string]*models.Item)
	itemMap["Coke"] = &models.Item{"Coke", 25, 10}
	itemMap["Pepsi"] = &models.Item{"Pepsi", 35, 10}
	itemMap["Soda"] = &models.Item{"Soda", 45, 10}
	return itemMap
}

// DisplayBalance displays current balance for seller and user
func DisplayBalance() {
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("Seller balance: ", vm.SellerBalance)
	fmt.Println("User balance: ", vm.UserBalance)
	fmt.Println("------------------------------------------------------------------------")
}

// CancelTransaction cancels the transaction & refund money to user
func CancelTransaction(productChoice string) {
	// refund money
	vm.UserBalance += vm.Items[productChoice].Price
	// debit the amount from seller account
	vm.SellerBalance -= vm.Items[productChoice].Price
}

// ConfirmTransaction confirms the transaction & debit user balance
func ConfirmTransaction(productChoice string) {
	// decrease available items
	vm.Items[productChoice].AvailableItems--
	// deduct balance
	vm.UserBalance -= vm.Items[productChoice].Price
	// credit the amount to seller account
	vm.SellerBalance += vm.Items[productChoice].Price
}

// ReturnProduct returns product & remaining change to seller
func ReturnProduct(productChoice string, returningTotal int) {
	// increase available items
	vm.Items[productChoice].AvailableItems++
	// debit the amount from seller account
	vm.SellerBalance -= vm.Items[productChoice].Price
	// credit the returned change to seller account
	vm.SellerBalance += returningTotal
	// refund money to user
	vm.UserBalance += (vm.Items[productChoice].Price + returningTotal)
}
