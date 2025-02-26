package repositories

import (
	"database/sql"
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetAllProductCPU() ([]models.ProductCPU, error) {
	db := storage.GetDB()

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       c.generation, c.microarchitecture, c.num_cores, c.num_threads, c.base_frequency, c.boost_frequency, c.max_memory_limit, c.wattage
		FROM product p
		JOIN product_cpu c ON p.id = c.product_id`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductCPU
	for rows.Next() {
		var product models.ProductCPU
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
			&product.Generation, &product.Microarchitecture, &product.NumCores, &product.NumThreads, &product.BaseFrequency, &product.BoostFrequency, &product.MaxMemoryLimit, &product.Wattage,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func GetCPUByID(id int) (*models.ProductCPU, error) {
	db := storage.GetDB()
	
	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image,
		       c.generation, c.microarchitecture, c.num_cores, c.num_threads, c.base_frequency, c.boost_frequency, c.max_memory_limit, c.wattage
		FROM product p
		JOIN product_cpu c ON p.id = c.product_id
		WHERE p.id = $1`
	var product models.ProductCPU
	err := db.QueryRow(query, id).Scan(
		&product.ID, &product.Brand, &product.Model, &product.CurrentPrice, &product.StockCount, &product.Category, &product.ProductImage,
		&product.Generation, &product.Microarchitecture, &product.NumCores, &product.NumThreads, &product.BaseFrequency, &product.BoostFrequency, &product.MaxMemoryLimit, &product.Wattage,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found")
		}
		return nil, err
	}
	return &product, nil
}

func InsertProductCPU(product models.ProductCPU) error {
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

	cpuQuery := `
		INSERT INTO product_cpu (product_id, generation, microarchitecture, num_cores, num_threads, base_frequency, boost_frequency, max_memory_limit, wattage)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err = db.Exec(cpuQuery, productID, product.Generation, product.Microarchitecture, product.NumCores, product.NumThreads, product.BaseFrequency, product.BoostFrequency, product.MaxMemoryLimit, product.Wattage)
	if err != nil {
		return err
	}

	return nil
}


func insertDatasetCPU() error{
	
	cpus := []models.ProductCPU{
		{
			Product: models.Product{
				Brand:        "Intel",
				Model:        "i9-13900K",
				CurrentPrice: 599,
				StockCount:   10,
				Category:     "CPU",

			},
			Generation:        "13th Gen",
			Microarchitecture: "Raptor Lake",
			NumCores:          24,
			NumThreads:        32,
			BaseFrequency:     3.0,
			BoostFrequency:    5.8,
			MaxMemoryLimit:    128,
			Wattage:           253,
		},
		{
			Product: models.Product{
				Brand:        "AMD",
				Model:        "Ryzen 9 7950X",
				CurrentPrice: 699,
				StockCount:   15,
				Category:     "CPU",
			},
			Generation:        "Zen 4",
			Microarchitecture: "Zen 4",
			NumCores:          16,
			NumThreads:        32,
			BaseFrequency:     4.5,
			BoostFrequency:    5.7,
			MaxMemoryLimit:    128,
			Wattage:           230,
		},
		{
			Product: models.Product{
				Brand:        "Intel",
				Model:        "i7-13700K",
				CurrentPrice: 409,
				StockCount:   20,
				Category:     "CPU",
			},
			Generation:        "13th Gen",
			Microarchitecture: "Raptor Lake",
			NumCores:          16,
			NumThreads:        24,
			BaseFrequency:     3.4,
			BoostFrequency:    5.4,
			MaxMemoryLimit:    128,
			Wattage:           190,
		},
		{
			Product: models.Product{
				Brand:        "AMD",
				Model:        "Ryzen 7 7700X",
				CurrentPrice: 399,
				StockCount:   25,
				Category:     "CPU",
			},
			Generation:        "Zen 4",
			Microarchitecture: "Zen 4",
			NumCores:          8,
			NumThreads:        16,
			BaseFrequency:     4.5,
			BoostFrequency:    5.4,
			MaxMemoryLimit:    128,
			Wattage:           105,
		},
	}

	for _, cpu := range cpus {
		err := InsertProductCPU(cpu)
		if err != nil {
			panic(err) 
		}
		
	}

	return nil
}