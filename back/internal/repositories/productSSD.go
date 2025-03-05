package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductSSD() ([]models.ProductSSD, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
			s.capacity, s.wattage
		FROM product p
		JOIN product_ssd s ON p.id = s.product_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductSSD
	for rows.Next() {
		var product models.ProductSSD
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.Capacity, &product.Wattage,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductSSDByID(id int) (*models.ProductSSD, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
			s.capacity, s.wattage
		FROM product p
		JOIN product_ssd s ON p.id = s.product_id
		WHERE p.id = $1`

	var product models.ProductSSD
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.Capacity, &product.Wattage,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductSSD(product models.ProductSSD) error {
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

	ssdQuery := `
		INSERT INTO product_ssd(product_id, capacity, wattage)
		VALUES ($1, $2, $3)`
	_, err = db.Exec(ssdQuery, productID, product.Capacity, product.Wattage)
	if err != nil {
		return err
	}

	return nil
}
