package models

import "time"

type DiscountCode struct {
	Code           int          `json:"code" db:"code"`
	Amount         float64      `json:"amount" db:"amount"`
	DiscountLimit  float64      `json:"discount_limit" db:"discount_limit"`
	UsageLimit     int          `json:"usage_limit" db:"usage_limit"`
	ExpirationTime time.Time    `json:"expiration_time" db:"expiration_time"`
	CodeType       DiscountEnum `json:"code_type" db:"code_type"`
}

type PrivateCode struct {
	Code      int       `json:"code" db:"code"`
	ClientID  int       `json:"client_id" db:"client_id"`
	Timestamp time.Time `json:"timestamp" db:"time_stamp"`
}