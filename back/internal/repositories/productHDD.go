package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductHDD() ([]models.ProductHDD, error) {
	db := storage.GetDB()
	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       h.capacity, h.rotational_speed, h.wattage, h.depth, h.height, h.width
		FROM product p
		JOIN product_hdd h ON p.id = h.product_id`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductHDD
	for rows.Next() {
		var product models.ProductHDD
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.Capacity, &product.RotationalSpeed, &product.Wattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetHDDByID(id int) (*models.ProductHDD, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		       h.capacity, h.rotational_speed, h.wattage, h.depth, h.height, h.width
		FROM product p
		JOIN product_hdd h ON p.id = h.product_id
		WHERE p.id = $1`

	var product models.ProductHDD
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.Capacity, &product.RotationalSpeed, &product.Wattage, &product.Depth, &product.Height, &product.Width,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	
	return &product, nil
}

func InsertProductHDD(product models.ProductHDD) error {
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

	hddQuery := `
		INSERT INTO product_hdd (product_id, capacity, rotational_speed, wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err = db.Exec(hddQuery, productID, product.Capacity, product.RotationalSpeed, product.Wattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}

func insertDatasetHDD() {

	hdds := []models.ProductHDD{
		{
			Product: models.Product{
				Brand:        "Seagate",
				Model:        "Barracuda",
				CurrentPrice: 50,
				StockCount:   100,
				Category:     "HDD",
			},
			Capacity:        1.0,
			RotationalSpeed: 7200,
			Wattage:         6,
			Depth:           5.0,
			Height:          1.0,
			Width:           3.5,
		},
		{
			Product: models.Product{
				Brand:        "Western Digital",
				Model:        "Blue",
				CurrentPrice: 60,
				StockCount:   80,
				Category:     "HDD",
			},
			Capacity:        2.0,
			RotationalSpeed: 5400,
			Wattage:         5,
			Depth:           5.0,
			Height:          1.0,
			Width:           3.5,
		},
		{
			Product: models.Product{
				Brand:        "Toshiba",
				Model:        "P300",
				CurrentPrice: 70,
				StockCount:   60,
				Category:     "HDD",
			},
			Capacity:        3.0,
			RotationalSpeed: 7200,
			Wattage:         7,
			Depth:           5.0,
			Height:          1.0,
			Width:           3.5,
		},
		{
			Product: models.Product{
				Brand:        "HGST",
				Model:        "Ultrastar",
				CurrentPrice: 120,
				StockCount:   40,
				Category:     "HDD",
			},
			Capacity:        4.0,
			RotationalSpeed: 7200,
			Wattage:         8,
			Depth:           5.0,
			Height:          1.0,
			Width:           3.5,
		},
	}

	for _, hdd := range hdds {
		err := InsertProductHDD(hdd)
		if err != nil {
			panic(err)
		}
	}
}