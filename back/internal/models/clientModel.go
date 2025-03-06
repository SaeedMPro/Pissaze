package models

import "time"

type ClientAbstract interface {
	IsVIP() bool
	GetClient() Client
}

type Client struct {
	ClientID         int               `json:"client_id" db:"client_id"`
	PhoneNumber      string            `json:"phone_number" db:"phone_number"`
	FirstName        string            `json:"first_name" db:"first_name"`
	LastName         string            `json:"last_name" db:"last_name"`
	WalletBalance    float64           `json:"wallet_balance" db:"wallet_balance"`
	Timestamp        time.Time         `json:"timestamp" db:"time_stamp"`
	ReferralCode     string            `json:"referral_code" db:"referral_code"`
	NumberOfReferred int               `json:"number_of_referred"`
	Addresses        []AddressOfClient `json:"addresses"`
}

type VIPClient struct {
	Client         Client    `json:"client" db:"client"`
	ExpirationTime time.Time `json:"expiration_time" db:"expiration_time"`
}

type AddressOfClient struct {
	ClientID      int    `json:"client_id" db:"client_id"`
	Province      string `json:"province" db:"province"`
	RemainAddress string `json:"remain_address" db:"remain_address"`
}

func (c Client) IsVIP() bool       { return false }
func (c Client) GetClient() Client { return c }

func (v VIPClient) IsVIP() bool       { return true }
func (v VIPClient) GetClient() Client { return v.Client }
