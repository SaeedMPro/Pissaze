package service

import (
	"fmt"
	"sort"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func FindcompatibleWithCarts(listID []int)([]models.Product, error){
	if len(listID) <= 0 {return nil, nil}

	final, err:= findcompatibleWithProduct(listID[0])
	if err != nil {
		return nil, err
	}
	fmt.Println("final is ->", final)

	for _, prouductID := range listID {
		compatible, err := findcompatibleWithProduct(prouductID)
		if err != nil {
			return nil, err
		}
		final = intersectionSlices(final, compatible)
		fmt.Println("final is ->", final)
	}
	return final, nil
}
//----------------------- helper ------------------------------------------
func findcompatibleWithProduct(prouductId int)([]models.Product, error){
	compatible, err := repositories.GetCompatibleByID(prouductId)
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

