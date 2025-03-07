package service

import (

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func GetClientCart(clientID int)([]models.ShoppingCart, error){
	carst, err := repositories.GetShoppingCartByClientID(clientID)
	return carst, err
}

func GetClientSummuryOfCarts(clientID, count int)([]models.LockedShoppingCart, error) {
	carts, err := repositories.GetLockedShoppingCartByClientID(clientID, count)
	if err != nil {
		return nil, err
	}

	for i := range carts {
		err = repositories.GetProductsInLockedShoppingCart(&carts[i])
		if err != nil {
			return nil, err
		}
	}
	return carts, nil
}