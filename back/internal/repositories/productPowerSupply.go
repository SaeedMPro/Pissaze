package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductPowerSupply() ([]models.ProductPowerSupply, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       ps.supported_wattage, ps.depth, ps.height, ps.width
		FROM product p
		JOIN product_power_supply ps ON p.id = rs.product_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductPowerSupply
	for rows.Next() {
		var product models.ProductPowerSupply
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.SupportedWattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductPowerSupplyByID(id int) (*models.ProductPowerSupply, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       ps.supported_wattage, ps.depth, ps.height, ps.width
		FROM product p
		JOIN product_power_supply c ON p.id = c.product_id
		WHERE p.id = $1`

	var product models.ProductPowerSupply
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.SupportedWattage, &product.Depth, &product.Height, &product.Width,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductPowerSupply(product models.ProductPowerSupply) error {
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

	powerSupplyQuery := `
		INSERT INTO product_power_supply(product_id, supported_wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5)`
	_, err = db.Exec(powerSupplyQuery, productID, product.SupportedWattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}
