package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetClientByPhoneNumber( phoneNumber string) (*models.Client, error) {
	
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

	return &client, nil
}

func GetClientAddressByClientID(ClientID int)(address []models.AddressOfClient, err error){
	addressQuery := `
		SELECT client_id, province, remain_address
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
		if err := rows.Scan(&address.ClientID,&address.Province, &address.RemainAddress); err != nil {
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
		WHERE client_id = $1 AND expiration_time >= NOW()`

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

func GetNumberOfReferredByClient(clientID int) (int, error){
	db := storage.GetDB()

	query := `
		SELECT COALESCE(COUNT(*), 0)
		FROM refers r
		JOIN client c ON r.referrer_id = c.referral_code
		WHERE c.client_id = $1`

	var numberOfReferred int
	err := db.QueryRow(query, clientID).Scan(&numberOfReferred)

	if err != nil {
		return 0, fmt.Errorf("failed to get referral count: %w", err)
	}
	return numberOfReferred, nil
}

func GetCurrentMonthVIPProfit(clientID int) (float64, error) {
    db := storage.GetDB()

    query := `
        SELECT COALESCE(SUM(adt.cart_price) * 0.15, 0)
        FROM vip_client vc
        JOIN issued_for ifo ON vc.client_id = ifo.client_id
        JOIN transaction t ON ifo.tracking_code = t.tracking_code
        JOIN added_to adt ON ifo.client_id = adt.client_id 
            AND ifo.cart_number = adt.cart_number 
            AND ifo.locked_number = adt.locked_number
        WHERE vc.client_id = $1
          AND t.transaction_status = 'successful'
          AND t.time_stamp >= DATE_TRUNC('month', CURRENT_DATE)
          AND t.time_stamp <= NOW()
          AND vc.expiration_time >= t.time_stamp`

    var cashback float64
    err := db.QueryRow(query, clientID).Scan(&cashback)
    
    if err != nil {
        return 0, fmt.Errorf("failed to calculate VIP profit: %w", err)
    }
    
    return cashback, nil
}

func InsertClient(client models.Client) (int, error) {
	db := storage.GetDB()
	var clientID int
	query := `
		INSERT INTO client (phone_number, first_name, last_name, referral_code)
		VALUES ($1, $2, $3, $4)
		RETURNING client_id`
	err := db.QueryRow(query, client.PhoneNumber, client.FirstName, client.LastName, client.ReferralCode).Scan(&clientID)
	if err != nil {
		return 0, err
	}
	return clientID, nil
}


func InsertAddress(addr models.AddressOfClient) error {
	db := storage.GetDB()
	query := `
		INSERT INTO address_of_client (client_id, province, remain_address)
		VALUES ($1, $2, $3)`
	_, err := db.Exec(query, addr.ClientID, addr.Province, addr.RemainAddress)
	return err
}

func InsertVIPClient(vip models.VIPClient) error {
	db := storage.GetDB()
	query := `
		INSERT INTO vip_client (client_id, expiration_time)
		VALUES ($1, $2)`
	_, err := db.Exec(query, vip.Client.ClientID, vip.ExpirationTime)
	return err
}
