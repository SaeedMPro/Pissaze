package repositories

import (
	"database/sql"
	"errors"
	"time"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)



func GetClientByPhoneNumber( phoneNumber string) (models.ClientAbstract, error) {
	
	var client models.Client
	query := `
		SELECT client_id, phone_number, first_name, last_name, 
		       wallet_balance, time_stamp, referral_code
		FROM client
		WHERE phone_number = $1`

	db := storage.GetDB()

	err := db.QueryRow( query, phoneNumber).Scan(
		&client.ClientID, &client.PhoneNumber, &client.FirstName, 
		&client.LastName, &client.WalletBalance, &client.Timestamp, 
		&client.ReferralCode,
	)


	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("client not found")
		}
		return nil, errors.New("undfine error")
	}

	return client, nil
}

func GetClientAddressByClientID(ClientID int)(address []models.AddressOfClient, err error){
	addressQuery := `
		SELECT province, remain_address
		FROM address_of_client
		WHERE client_id = $1`

	db := storage.GetDB()
	rows, err := db.Query(addressQuery, ClientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var addresses []models.AddressOfClient
	for rows.Next() {
		var address models.AddressOfClient
		if err := rows.Scan(&address.Province, &address.RemainAddress); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}
	return addresses, nil
}

func GetVIPClientByID(ClientID int)(client *models.VIPClient,err error){
	
	var expirationTime time.Time 
	vipQuery := `
		SELECT expiration_time
		FROM vip_client
		WHERE client_id = $1`

	db := storage.GetDB()
	err = db.QueryRow(vipQuery, ClientID).Scan(&expirationTime)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, nil 
	case err != nil:
		return nil, err
	default:
		return &models.VIPClient{
			ExpirationTime: expirationTime,
		}, nil
	}
}


// ----- validators -----
// func validate(client *client.ClientAbstract)(err error){
// 	//TODO:
// 	return
// }