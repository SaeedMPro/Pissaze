package models

type CoolingMethodEnum string
const (
	CoolingMethodLiquid CoolingMethodEnum = "liquid"
	CoolingMethodAir    CoolingMethodEnum = "air"
)

type DiscountEnum string
const (
	DiscountPublic  DiscountEnum = "public"
	DiscountPrivate DiscountEnum = "private"
)


type TransactionStatusEnum string
const (
	TransactionStatusSuccessful     TransactionStatusEnum = "Successful"
	TransactionStatusSemiSuccessful TransactionStatusEnum = "semi-successful"
	TransactionStatusUnsuccessful   TransactionStatusEnum = "unsuccessful"
)


type TransactionTypeEnum string
const (
	TransactionTypeBank   TransactionTypeEnum = "bank"
	TransactionTypeWallet TransactionTypeEnum = "wallet"
)


type CartStatusEnum string
const (
	CartStatusLocked  CartStatusEnum = "locked"
	CartStatusBlocked CartStatusEnum = "blocked"
	CartStatusActive  CartStatusEnum = "active"
)