package repositories

import (
	"database/sql"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllDiscountCodeNotExpired() ([]models.DiscountCodeInterface, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			d.code, d.amount, d.discount_limit, d.usage_limit, d.expiration_time, d.code_type
			p.client_id, p.time_stamp
		FROM discount_code d
		RIGHT JOIN private_code p ON d.code = p.code
		WHERE expiration_time > NOW()`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var discountCodes []models.DiscountCodeInterface
	for rows.Next() {
		var code models.DiscountCode
		var clientID sql.NullInt64
		var timestamp sql.NullTime

		err := rows.Scan(
			&code.Code, &code.Amount, &code.DiscountLimit, &code.UsageLimit, &code.ExpirationTime, &code.CodeType,
			&clientID, &timestamp,
		)

		if err != nil {
			return nil, err
		}

		if clientID.Valid {
			privateCode := models.PrivateCode{
				DiscountCode: code,
				ClientID:     int(clientID.Int64),
				Timestamp:    timestamp.Time,
			}
			discountCodes = append(discountCodes, &privateCode)
		} else {
			discountCodes = append(discountCodes, &code)
		}
	}

	return discountCodes, nil
}

func GetPrivateCodesWithLessThenIntervalDayByClientID(clientID int, day int) ([]models.PrivateCode, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			d.code, d.amount, d.discount_limit, d.usage_limit, d.expiration_time, d.code_type,
			p.client_id, p.time_stamp
		FROM discount_code d
		INNER JOIN private_code p ON d.code = p.code
		WHERE p.client_id = $12 AND d.expiration_time BETWEEN NOW() AND NOW() + INTERVAL $1 DAY`

	rows, err := db.Query(query, day, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privateCodes []models.PrivateCode
	for rows.Next() {
		var code models.PrivateCode
		err := rows.Scan(
			&code.Code, &code.Amount, &code.DiscountLimit, &code.UsageLimit, &code.ExpirationTime, &code.CodeType,
			&code.ClientID, &code.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		privateCodes = append(privateCodes, code)
	}

	return privateCodes, nil
}

func GetPrivateCodesByClientIDNotExpire(clientID int) ([]models.PrivateCode, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			d.code, d.amount, d.discount_limit, d.usage_limit, d.expiration_time, d.code_type,
			p.client_id, p.time_stamp
		FROM discount_code d
		INNER JOIN private_code p ON d.code = p.code
		WHERE p.client_id = $1 AND expiration_time > NOW()`

	rows, err := db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var privateCodes []models.PrivateCode
	for rows.Next() {
		var code models.PrivateCode
		err := rows.Scan(
			&code.Code, &code.Amount, &code.DiscountLimit, &code.UsageLimit, &code.ExpirationTime, &code.CodeType,
			&code.ClientID, &code.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		privateCodes = append(privateCodes, code)
	}

	return privateCodes, nil
}

func GetNumberOfGiftedDiscountCode(clientID int) (int, error){
	db := storage.GetDB()

	query := `
	WITH RECURSIVE referral_chain AS (
		SELECT rr.referee_id, rr.referrer_id
        FROM refers rr
		JOIN client c ON rr.referrer_id = c.referral_code
        WHERE c.client_id = $1
		
		UNION ALL

        SELECT r.referee_id, r.referrer_id
        FROM referral_chain rc
        JOIN refers r ON rc.referee_id = r.referrer_id
    )
    SELECT COUNT(*) AS gifted_codes_count
    FROM referral_chain`

	var numberOfGiftedCode int
	err := db.QueryRow(query, clientID).Scan(&numberOfGiftedCode)
	if err != nil {
		return 0, err
	}

	return numberOfGiftedCode, nil
}