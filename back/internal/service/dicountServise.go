package service

import (
	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

func GetClientPrivateCode(client models.ClientAbstract, day int)([]models.PrivateCode,error){
	codes, err := repositories.GetPrivateCodesWithLessThenIntervalDayByClientID(
		client.GetClient().ClientID, day,
	)
	return codes,err
}

func NumberOfGiftedCode(client models.ClientAbstract)(int,error){
	count, err := repositories.GetNumberOfGiftedDiscountCode(client.GetClient().ClientID)
	return count, err
}