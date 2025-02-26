package service

import (
	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func GetAllCPU()([]models.ProductCPU,error){
	CPU, err := repositories.GetAllProductCPU()
	return CPU, err
}

func GetAllCooler()([]models.ProductCooler,error){
	cooler, err := repositories.GetAllProductCooler()
	return cooler, err

}

func GetAllHDD()([]models.ProductHDD,error){
	HDD, err := repositories.GetAllProductHDD()
	return HDD, err
}

