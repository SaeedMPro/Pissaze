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
	//TODO:
	return nil, nil
}

func GETAllRam()([]models.ProductRAMStick,error){
	//TODO:
	return nil, nil
}

func GETAllGPU()([]models.ProductGPU,error){
	//TODO:
	return nil, nil
}

func GETAllPowerSuply()([]models.ProductPowerSupply,error){
	//TODO:
	return nil, nil
}

func GETAllCase()([]models.ProductCase,error){
	//TODO:
	return nil, nil
}

func GETAllSSD()([]models.ProductSSD,error){
	//TODO:
	return nil, nil
}

func GETAllProducts()([]models.ProductInterface,error){
	var list []models.ProductInterface

	cpus, err := GetAllCPU()
	if err != nil {
		return nil, err
	}
	for _, tmp := range cpus {
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

	powerSuply, err := GETAllPowerSuply()
	if err != nil {
		return nil, err
	}
	for _, tmp := range powerSuply {
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

