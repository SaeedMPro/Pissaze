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

func GETAllMotherBoards()([]models.ProductMotherboard,error){
	motherBoard, err := repositories.GetAllProductMotherboard()
	return motherBoard,err
}

func GETAllRam()([]models.ProductRAMStick,error){
	ram, err := repositories.GetAllProductRAMStick()
	return ram, err
}

func GETAllGPU()([]models.ProductGPU,error){
	gpu, err := repositories.GetAllProductGPU()
	return gpu, err
}

func GETAllPowerSupply()([]models.ProductPowerSupply,error){
	power, err := repositories.GetAllProductPowerSupply()
	return power, err
}

func GETAllCase()([]models.ProductCase,error){
	pCase, err := repositories.GetAllProductCase()
	return pCase, err
}

func GETAllSSD()([]models.ProductSSD,error){
	ssd, err := repositories.GetAllProductSSD()
	return ssd, err
}

func GETAllProducts()([]models.ProductInterface,error){
	var list []models.ProductInterface

	cpu, err := GetAllCPU()
	if err != nil {
		return nil, err
	}
	for _, tmp := range cpu {
		list = append(list, tmp)
	}

	cooler, err := GetAllCooler()
	if err != nil {
		return nil, err
	}
	for _, tmp := range cooler {
		list = append(list, tmp)
	}

	HDD, err := GetAllHDD()
	if err != nil {
		return nil, err
	}
	for _, tmp := range HDD {
		list = append(list, tmp)
	}

	SSD, err := GETAllSSD()
	if err != nil {
		return nil, err
	}
	for _, tmp := range SSD {
		list = append(list, tmp)
	}

	cases, err := GETAllCase()
	if err != nil {
		return nil, err
	}
	for _, tmp := range cases {
		list = append(list, tmp)
	}

	powerSupply, err := GETAllPowerSupply()
	if err != nil {
		return nil, err
	}
	for _, tmp := range powerSupply {
		list = append(list, tmp)
	}

	rams, err := GETAllRam()
	if err != nil {
		return nil, err
	}
	for _, tmp := range rams {
		list = append(list, tmp)
	}

	return list, nil
}

