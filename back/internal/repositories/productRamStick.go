package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductRAMStick() ([]models.ProductRAMStick, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       rs.generation, rs.capacity, rs.frequency, rs.wattage, rs.depth, rs.height, rs.width
		FROM product p
		JOIN product_ram_stick rs ON p.id = rs.product_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductRAMStick
	for rows.Next() {
		var product models.ProductRAMStick
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.Generation, &product.Capacity, &product.Frequency, &product.Wattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductRAmStickByID(id int) (*models.ProductRAMStick, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       c.cooling_method, c.fan_size, c.max_rotational_speed, c.wattage, c.depth, c.height, c.width
		FROM product p
		JOIN product_ram_stick c ON p.id = c.product_id
		WHERE p.id = $1`

	var product models.ProductRAMStick
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.Generation, &product.Capacity, &product.Frequency, &product.Wattage, &product.Depth, &product.Height, &product.Width,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductRAMStick(product models.ProductRAMStick) error {
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

	ramStickQuery := `
		INSERT INTO product_ram_stick(product_id, cooling_method, fan_size, max_rotational_speed, wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5, $6, $7,$8)`
	_, err = db.Exec(ramStickQuery, productID, product.Generation, product.Capacity, product.Frequency, product.Wattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}
