package service

import (
	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func GetClientCart(clientID int)([]models.ShoppingCart, error){
	carst, err := repositories.GetShoppingCartByClientID(clientID)
	return carst, err
}