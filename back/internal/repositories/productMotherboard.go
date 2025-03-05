package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductMotherboard() ([]models.ProductMotherboard, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       mb.chipset_name, mb.num_memory_slots, mb.memory_speed_range, mb.wattage, mb.depth, mb.height, mb.width
		FROM product p
		JOIN product_motherboard mb ON p.id = mb.product_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductMotherboard
	for rows.Next() {
		var product models.ProductMotherboard
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.ChipsetName, &product.NumMemorySlots, &product.MemorySpeedRange, &product.Wattage, &product.Depth, &product.Height, &product.Width,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetProductMotherboardByID(id int) (*models.ProductMotherboard, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       mb.chipset_name, mb.num_memory_slots, mb.memory_speed_range, mb.wattage, mb.depth, mb.height, mb.width
		FROM product p
		JOIN product_motherboard mb ON p.id = mb.product_id
		WHERE p.id = $1`

	var product models.ProductMotherboard
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.ChipsetName, &product.NumMemorySlots, &product.MemorySpeedRange, &product.Wattage, &product.Depth, &product.Height, &product.Width,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductMotherboard(product models.ProductMotherboard) error {
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

	motherboardQuery := `
		INSERT INTO product_ram_stick (product_id, chipset_name, num_memory_slots, memory_speed_range, wattage, depth, height, width)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err = db.Exec(motherboardQuery, productID, product.ChipsetName, product.NumMemorySlots, product.MemorySpeedRange, product.Wattage, product.Depth, product.Height, product.Width)
	if err != nil {
		return err
	}

	return nil
}
