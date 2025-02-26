package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductCooler() ([]models.ProductCooler, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       c.cooling_method, c.fan_size, c.max_rotational_speed, c.wattage, c.depth, c.height, c.width
		FROM product p
		JOIN product_cooler c ON p.id = c.product_id`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductCooler
	for rows.Next() {
		var product models.ProductCooler
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.CoolingMethod, &product.FanSize, &product.MaxRotationalSpeed, &product.Wattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetCoolerByID(id int) (*models.ProductCooler, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       c.cooling_method, c.fan_size, c.max_rotational_speed, c.wattage, c.depth, c.height, c.width
		FROM product p
		JOIN product_cooler c ON p.id = c.product_id
		WHERE p.id = $1`
	var product models.ProductCooler
	err :=db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.CoolingMethod, &product.FanSize, &product.MaxRotationalSpeed, &product.Wattage, &product.Depth, &product.Height, &product.Width,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductCooler(product models.ProductCooler) error {
	db := storage.GetDB()

	productQuery := `
		INSERT INTO product (brand, model, current_price, stock_count, category, product_image)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	var productID int
	err := db.QueryRow(productQuery, product.Brand, product.Model, product.CurrentPrice, product.StockCount, product.Category, product.ProductImage).Scan(&productID)
	if err != nil {
		return err
	}

	coolerQuery := `
		INSERT INTO product_cooler (product_id, cooling_method, fan_size, max_rotational_speed, wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.Exec(coolerQuery, productID, product.CoolingMethod, product.FanSize, product.MaxRotationalSpeed, product.Wattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}

func insertDatasetCooler() {
	
	coolers := []models.ProductCooler{
		{
			Product: models.Product{
				Brand:        "Noctua",
				Model:        "NH-D15",
				CurrentPrice: 100,
				StockCount:   50,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodAir,
			FanSize:            140,
			MaxRotationalSpeed: 1500,
			Wattage:            10,
			Depth:              165,
			Height:             165,
			Width:              150,
		},
		{
			Product: models.Product{
				Brand:        "Cooler Master",
				Model:        "Hyper 212",
				CurrentPrice: 40,
				StockCount:   100,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodAir,
			FanSize:            120,
			MaxRotationalSpeed: 2000,
			Wattage:            8,
			Depth:              120,
			Height:             158,
			Width:             120,
		},
		{
			Product: models.Product{
				Brand:        "Corsair",
				Model:        "H100i",
				CurrentPrice: 150,
				StockCount:   30,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodLiquid,
			FanSize:            120,
			MaxRotationalSpeed: 2400,
			Wattage:            12,
			Depth:              277,
			Height:             120,
			Width:             120,
		},
		{
			Product: models.Product{
				Brand:        "NZXT",
				Model:        "Kraken X63",
				CurrentPrice: 160,
				StockCount:   25,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodLiquid,
			FanSize:            140,
			MaxRotationalSpeed: 1800,
			Wattage:            15,
			Depth:              315,
			Height:             143,
			Width:             140,
		},
		{
			Product: models.Product{
				Brand:        "be quiet!",
				Model:        "Dark Rock Pro 4",
				CurrentPrice: 90,
				StockCount:   40,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodAir,
			FanSize:            135,
			MaxRotationalSpeed: 1500,
			Wattage:            9,
			Depth:              145,
			Height:             163,
			Width:             135,
		},
		{
			Product: models.Product{
				Brand:        "Thermaltake",
				Model:        "Water 3.0",
				CurrentPrice: 130,
				StockCount:   20,
				Category:     "Cooler",
			},
			CoolingMethod:      models.CoolingMethodLiquid,
			FanSize:            120,
			MaxRotationalSpeed: 2200,
			Wattage:            14,
			Depth:              276,
			Height:             120,
			Width:             120,
		},
	}

	// Insert each cooler into the database
	for _, cooler := range coolers {
		err := InsertProductCooler(cooler)
		if err != nil {
			panic(err)
		}
	}
}