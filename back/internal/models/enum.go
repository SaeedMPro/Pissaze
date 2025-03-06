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


type CategoryProductEnum string
const (
    CategoryMotherboard  CategoryProductEnum = "Motherboard"
    CategoryCPU          CategoryProductEnum = "CPU"
    CategoryGPU         CategoryProductEnum = "GPU"
    CategoryRAMStick    CategoryProductEnum = "RAM Stick"
    CategoryCooler      CategoryProductEnum = "Cooler"
    CategoryPowerSupply CategoryProductEnum = "Power Supply"
    CategoryCase        CategoryProductEnum = "Case"
    CategorySSD         CategoryProductEnum = "SSD"
    CategoryHDD         CategoryProductEnum = "HDD"
)
