package service

import (
	"sort"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func FindCompatibleWithCarts(listID []int)([]models.Product, error){
	if len(listID) <= 0 {return nil, nil}

	final, err := findCompatibleWithProduct(listID[0])
	if err != nil {
		return nil, err
	}

	for _, productID := range listID {
		compatible, err := findCompatibleWithProduct(productID)
		if err != nil {
			return nil, err
		}
		final = intersectionSlices(final, compatible)
	}
	return final, nil
}

func FilterBy(base []models.Product, filter []models.CategoryProductEnum)(ans []models.Product){
	if len(filter) <= 0 {
		return base
	}
	
	filterMap := make(map[string]bool)
	for _, val := range filter {
		filterMap[string(val)] = val.IsValid()
	}

	for _,product := range base {
		if filterMap[product.Category] {
			ans = append(ans, product)
		}
	}
	return ans
}

//----------------------- helper ------------------------------------------
func findCompatibleWithProduct(productId int)([]models.Product, error){
	compatible, err := repositories.GetCompatibleByID(productId)
	return compatible, err 
}

func intersectionSlices(base, new []models.Product)(ans []models.Product){
	sort.Slice(base, func(i, j int) bool {
		return base[i].ID < base[j].ID
	})
	sort.Slice(new, func(i, j int) bool {
		return new[i].ID < new[j].ID
	})

	basePtr, newPtr := 0, 0
	for basePtr < len(base) && newPtr < len(new) {
		if base[basePtr].ID == new[newPtr].ID {
			ans = append(ans, base[basePtr])
			basePtr++
			newPtr++

		} else if base[basePtr].ID > new[newPtr].ID {
			newPtr++

		}else { // base[basePtr].ID < new[newPtr].ID
			basePtr++
		}
	}
	return ans
}

