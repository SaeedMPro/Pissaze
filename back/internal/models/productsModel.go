package models

type ProductInterface interface {
	GetType() string
}

type Product struct {
	ID           int    `json:"id" db:"id"`
	Brand        string `json:"brand" db:"brand"`
	Model        string `json:"model" db:"model"`
	CurrentPrice int    `json:"current_price" db:"current_price"`
	StockCount   int    `json:"stock_count" db:"stock_count"`
	Category     string `json:"category" db:"category"`
	ProductImage []byte `json:"product_image" db:"product_image"`
}

type ProductHDD struct {
	Product
	Capacity        float64 `json:"capacity" db:"capacity"`
	RotationalSpeed int     `json:"rotational_speed" db:"rotational_speed"`
	Wattage         int     `json:"wattage" db:"wattage"`
	Depth           float64 `json:"depth" db:"depth"`
	Height          float64 `json:"height" db:"height"`
	Width           float64 `json:"width" db:"width"`
}

type ProductCooler struct {
	Product
	CoolingMethod      CoolingMethodEnum `json:"cooling_method" db:"cooling_method"`
	FanSize            int               `json:"fan_size" db:"fan_size"`
	MaxRotationalSpeed int               `json:"max_rotational_speed" db:"max_rotational_speed"`
	Wattage            int               `json:"wattage" db:"wattage"`
	Depth              float64           `json:"depth" db:"depth"`
	Height             float64           `json:"height" db:"height"`
	Width              float64           `json:"width" db:"width"`
}

type ProductCPU struct {
	Product
	Generation        string  `json:"generation" db:"generation"`
	Microarchitecture string  `json:"microarchitecture" db:"microarchitecture"`
	NumCores          int     `json:"num_cores" db:"num_cores"`
	NumThreads        int     `json:"num_threads" db:"num_threads"`
	BaseFrequency     float64 `json:"base_frequency" db:"base_frequency"`
	BoostFrequency    float64 `json:"boost_frequency" db:"boost_frequency"`
	MaxMemoryLimit    int     `json:"max_memory_limit" db:"max_memory_limit"`
	Wattage           int     `json:"wattage" db:"wattage"`
}

type ProductRAMStick struct {
	Product
	Generation string  `json:"generation" db:"generation"`
	Capacity   float64 `json:"capacity" db:"capacity"`
	Frequency  float64 `json:"frequency" db:"frequency"`
	Wattage    int     `json:"wattage" db:"wattage"`
	Depth      float64 `json:"depth" db:"depth"`
	Height     float64 `json:"height" db:"height"`
	Width      float64 `json:"width" db:"width"`
}

type ProductCase struct {
	Product
	ProductType string  `json:"product_type" db:"product_type"`
	Color       string  `json:"color" db:"color"`
	Material    string  `json:"material" db:"material"`
	FanSize     int     `json:"fan_size" db:"fan_size"`
	NumFans     int     `json:"num_fans" db:"num_fans"`
	Wattage     int     `json:"wattage" db:"wattage"`
	Depth       float64 `json:"depth" db:"depth"`
	Height      float64 `json:"height" db:"height"`
	Width       float64 `json:"width" db:"width"`
}

type ProductPowerSupply struct {
	Product
	SupportedWattage int     `json:"supported_wattage" db:"supported_wattage"`
	Depth            float64 `json:"depth" db:"depth"`
	Height           float64 `json:"height" db:"height"`
	Width            float64 `json:"width" db:"width"`
}

type ProductGPU struct {
	Product
	RAMSize    int     `json:"ram_size" db:"ram_size"`
	ClockSpeed float64 `json:"clock_speed" db:"clock_speed"`
	NumFans    int     `json:"num_fans" db:"num_fans"`
	Wattage    int     `json:"wattage" db:"wattage"`
	Depth      float64 `json:"depth" db:"depth"`
	Height     float64 `json:"height" db:"height"`
	Width      float64 `json:"width" db:"width"`
}

type ProductSSD struct {
	Product
	Capacity float64 `json:"capacity" db:"capacity"`
	Wattage  int     `json:"wattage" db:"wattage"`
}

type ProductMotherboard struct {
	Product
	ChipsetName      string  `json:"chipset_name" db:"chipset_name"`
	NumMemorySlots   int     `json:"num_memory_slots" db:"num_memory_slots"`
	MemorySpeedRange float64 `json:"memory_speed_range" db:"memory_speed_range"`
	Wattage          int     `json:"wattage" db:"wattage"`
	Depth            float64 `json:"depth" db:"depth"`
	Height           float64 `json:"height" db:"height"`
	Width            float64 `json:"width" db:"width"`
}

type Compatible struct {
	FirstProduct ProductInterface
	SecondProduct ProductInterface
}

func (p ProductCPU) GetType() string         { return "CPU" }
func (p ProductCooler) GetType() string      { return "Cooler" }
func (p ProductHDD) GetType() string         { return "HDD" }
func (p ProductCase) GetType() string        { return "Case" }
func (p ProductGPU) GetType() string         { return "GPU" }
func (p ProductMotherboard) GetType() string { return "Motherboard" }
func (p ProductSSD) GetType() string         { return "SSD" }
func (p ProductPowerSupply) GetType() string { return "Power Supply" }
func (p ProductRAMStick) GetType() string    { return "RAM stick" }
//TODO: p*
