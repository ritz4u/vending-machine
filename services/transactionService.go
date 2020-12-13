package services

import (
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
