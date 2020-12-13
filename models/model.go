package models

// VendingMachine model for Vending Machine
type VendingMachine struct {
	Available     bool
	Coins         map[string]*Coin
	Items         map[string]*Item
	UserBalance   int
	SellerBalance int
}

// Coin model for coin
type Coin struct {
	Choice int
	Type   int
}

// Item model for item
type Item struct {
	Type           string
	Price          int
	AvailableItems int
}
