package models

import "time"

type CartInterface interface {
	GetStatus() CartStatusEnum
	GetCart() ShoppingCart
}

type ShoppingCart struct {
	CartNumber int            `json:"cart_number" db:"cart_number"`
	ClientID   int            `json:"client_id" db:"client_id"`
	CartStatus CartStatusEnum `json:"cart_status" db:"cart_status"`
}

type LockedShoppingCart struct {
	ShoppingCart    					   `json:"shopping_cart"`
	Products         []ProductShoppingCart `json:"products"`
	TotalPrice       float64               `json:"total_price" db:"total_price"`
	LockedCartNumber int                   `json:"locked_cart_number" db:"locked_cart_number"`
	TimeStamp        time.Time             `json:"timestamp" db:"time_stamp"`
}

type ProductShoppingCart struct {
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity" db:"quantity"`
	CartPrice float64 `json:"cart_price" db:"cart_price"`
}

func (s ShoppingCart) GetCart() ShoppingCart     { return s }
func (s ShoppingCart) GetStatus() CartStatusEnum { return s.CartStatus }

func (ls LockedShoppingCart) GetCart() ShoppingCart              { return ls.ShoppingCart }
func (ls LockedShoppingCart) GetStatus() CartStatusEnum          { return ls.CartStatus }
func (ls LockedShoppingCart) GetProducts() []ProductShoppingCart { return ls.Products }
