package models

import "time"

type CartInterface interface {
	GetStatus()
	GetCart()
}

type ShoppingCart struct {
	CartNumber int
	ClientID   int
	CartStatus CartStatusEnum
}

type LockedShoppingCart struct {
	ShoppingCart
	Products         []ProductShoppingCart
	LockedCartNumber int
	TimeStamp        time.Time
}

type ProductShoppingCart struct {
	Product   Product
	Quantity  int
	CartPrice float64
}

func (s ShoppingCart) GetCart() ShoppingCart     { return s }
func (s ShoppingCart) GetStatus() CartStatusEnum { return s.CartStatus }

func (ls LockedShoppingCart) GetCart() ShoppingCart              { return ls.ShoppingCart }
func (ls LockedShoppingCart) GetStatus() CartStatusEnum          { return ls.CartStatus }
func (ls LockedShoppingCart) GetProducts() []ProductShoppingCart { return ls.Products }
