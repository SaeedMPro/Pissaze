package service

import (
	"errors"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/repositories"
)

// GetClientByPhoneNumber fetches a client by phone number, including VIP status and addresses
func GetClientByPhoneNumber(phoneNumber string) (models.ClientAbstract, error) {
	
	if phoneNumber == "" {
		return nil, errors.New("phone number is required")
	}

	client, err := repositories.GetClientByPhoneNumber(phoneNumber)
	if err != nil {
		return nil, err
	}

	client.Addresses, err = repositories.GetClientAddressByClientID(client.ClientID)
	if err != nil {
		return nil, err
	}

	clientVIP, err := repositories.GetVIPClientByID(client.ClientID)
	if err != nil {
		return nil, err
	}else if clientVIP != nil{
		clientVIP.Client = *client
		return clientVIP, nil
	}

	return client, nil
}
