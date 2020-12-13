package services

import (
	"testing"
)

func TestConfirmTransactionSuccessCase(t *testing.T) {
	vm := Init()
	type args struct {
		productChoice string
	}
	tests := []struct {
		name          string
		productChoice string
		want          int
	}{
		{"t1", "Coke", 25},
		{"t2", "Pepsi", 60},
		{"t3", "Soda", 105},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			if got := vm.SellerBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			} else {
				t.Logf("PASS: got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfirmTransactionFailCase(t *testing.T) {
	vm := Init()
	type args struct {
		productChoice string
	}
	tests := []struct {
		name          string
		productChoice string
		want          int
	}{
		{"t1", "Coke", 125},
		{"t2", "Pepsi", 30},
		{"t3", "Soda", 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			if got := vm.SellerBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelTransactionSuccessCase(t *testing.T) {
	vm := Init()
	type args struct {
		productChoice string
	}
	tests := []struct {
		name          string
		productChoice string
		want          int
	}{
		{"t1", "Coke", 1000},
		{"t2", "Pepsi", 1000},
		{"t3", "Soda", 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			CancelTransaction(tt.productChoice)
			if got := vm.UserBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCancelTransactionFailCase(t *testing.T) {
	vm := Init()
	type args struct {
		productChoice string
	}
	tests := []struct {
		name          string
		productChoice string
		want          int
	}{
		{"t1", "Coke", 1025},
		{"t2", "Pepsi", 975},
		{"t3", "Soda", 25},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			CancelTransaction(tt.productChoice)
			if got := vm.UserBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnProductSuccessCase(t *testing.T) {
	vm := Init()
	tests := []struct {
		name           string
		productChoice  string
		returningTotal int
		want           int
	}{
		{"t1", "Coke", 0, 1000},
		{"t2", "Pepsi", 0, 1000},
		{"t3", "Soda", 0, 1000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			ReturnProduct(tt.productChoice, tt.returningTotal)
			if got := vm.UserBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestReturnProductFailCase(t *testing.T) {
	vm := Init()
	tests := []struct {
		name           string
		productChoice  string
		returningTotal int
		want           int
	}{
		{"t1", "Coke", 0, 975},
		{"t2", "Pepsi", 0, 965},
		{"t3", "Soda", 0, 955},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConfirmTransaction(tt.productChoice)
			ReturnProduct(tt.productChoice, tt.returningTotal)
			if got := vm.UserBalance; got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
