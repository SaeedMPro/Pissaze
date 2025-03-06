package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductGPU() ([]models.ProductGPU, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       g.ram_size, g.clock_speed, g.num_fans, g.wattage, g.depth, g.height, g.width
		FROM product p
		JOIN product_gpu g ON p.id = g.product_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductGPU
	for rows.Next() {
		var product models.ProductGPU
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.RAMSize, &product.ClockSpeed, &product.NumFans, &product.Wattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductGPUByID(id int) (*models.ProductGPU, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       g.ram_size, g.clock_speed, g.num_fans, g.wattage, g.depth, g.height, g.width
		FROM product p
		JOIN product_gpu g ON p.id = g.product_id
		WHERE p.id = $1`

	var product models.ProductGPU
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.RAMSize, &product.ClockSpeed, &product.NumFans, &product.Wattage, &product.Depth, &product.Height, &product.Width,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductGPU(product models.ProductGPU) error {
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

	GPUQuery := `
		INSERT INTO product_gpu(product_id, ram_size, clock_speed, num_fans, wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.Exec(GPUQuery, productID, product.RAMSize, product.ClockSpeed, product.NumFans, product.Wattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}
