package repositories

import (
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func CompatibleWithCpu(id int)(compatibleProducts []models.Product,err error){ // ok
	db := storage.GetDB()

	cpu, err:= GetCPUByID(id)
	if err != nil {
		return nil, err
	}

	query := `
		-- mother board -> on socket 
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_mc_socket cmc ON p.id = cmc.motherboard_id
		WHERE cmc.cpu_id = $1

		UNION

		-- memory limit and cpu frequcy limit : more ram
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_ram_stick prs ON p.id = prs.product_id
		WHERE prs.capacity >= $2 AND prs.frequency BETWEEN $3 AND $4

		UNION

		-- cooler only on soket
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_cc_socket ccc ON p.id = ccc.cooler_id
		WHERE ccc.cpu_id = $1

		UNION

		-- supoer volatage
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $5

		UNION

		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		WHERE p.category IN ('GPU', 'SSD', 'HDD');
	`

	rows, err := db.Query(query, cpu.ID, cpu.MaxMemoryLimit, cpu.BaseFrequency, cpu.BoostFrequency, cpu.Wattage)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Brand, &p.Model, &p.CurrentPrice, &p.StockCount, &p.Category, &p.ProductImage)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibleProducts, nil
}

func CompatibleWithMotherboard(id int) (compatibleProducts []models.Product,err error) {
	db := storage.GetDB()

	mb, err := GetProductMotherboardByID(id)
	if err != nil {
		return nil, err
	}

	query := `
		-- cpu only on soket 
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_mc_socket cmc ON p.id = cmc.cpu_id
		WHERE cmc.motherboard_id = $1

		UNION

		-- generation suport and frequnce
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_rm_slot crm ON p.id = crm.ram_id
		JOIN product_ram_stick prs ON p.id = prs.product_id
		WHERE prs.frequency <= $2

		UNION

		-- sokcet
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_gm_slot cgm ON p.id = cgm.gpu_id
		WHERE cgm.motherboard_id = $1

		UNION

		-- sokcet
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_sm_slot csm ON p.id = csm.ssd_id
		WHERE csm.motherboard_id = $1

		UNION

		-- power
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $3

		UNION

		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		WHERE p.category IN ('Cooler', 'HDD')
	`

	rows, err := db.Query(query, mb.ID, mb.MemorySpeedRange, mb.Wattage)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return compatibleProducts, nil
}


func CompatibleWithRAM(id int) (compatibleProducts []models.Product, err error) {
	db := storage.GetDB()

	ram, err := GetProductRAmStickByID(id)
	if err != nil {
		return nil, err
	}

	query := `
		-- mother board -> slot and frequncy less than
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_motherboard pm ON p.id = pm.product_id
		JOIN compatible_rm_slot crm ON pm.product_id = crm.motherboard_id
		WHERE crm.ram_id = $1 
		AND pm.memory_speed_range >= $2

		UNION

		-- capcit bigger than cpu ad frecuency less 
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_cpu pc ON p.id = pc.product_id
		WHERE pc.max_memory_limit <= $3 
		AND $4 BETWEEN pc.base_frequency AND pc.boost_frequency

		UNION

		-- lest voltage
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $5

		UNION

		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		WHERE p.category IN ('GPU', 'Cooler', 'SSD', 'HDD')
	`

	rows, err := db.Query(query,ram.ID,ram.Frequency,ram.Capacity,ram.Frequency,ram.Wattage)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return compatibleProducts, nil
}


func CompatibleWithCooler(id int) (compatibleProducts []models.Product,err error) {
	db := storage.GetDB()
	cooler, err := GetCoolerByID(id)
	if err != nil {
		return nil, err
	}
	
	query := `
		-- cpu with socket
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_cc_socket ccs ON p.id = ccs.cpu_id
		WHERE ccs.cooler_id = $1

		UNION

		-- power suplay energy
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $2

		UNION

		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		WHERE p.category IN ('Motherboard', 'HDD', 'SSD', 'GPU', 'RAM Stick')
	`

	rows, err := db.Query(query, cooler.ID, cooler.Wattage)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}
	
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return compatibleProducts, nil
}


func CompatibleWithSSD(id int) (compatibleProducts []models.Product, err error) {
	db := storage.GetDB()

	ssd, err := GetProductSSDByID(id)
	if err != nil {
		return nil, err
	}

	query := `
		-- with table
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_sm_slot csm ON p.id = csm.motherboard_id
		WHERE csm.ssd_id = $1

		UNION

		--  for battery
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $2

		UNION

		SELECT id, brand, model, current_price, stock_count, category, product_image
		FROM product
		WHERE category IN ('Motherboard', 'HDD', 'CPU', 'GPU', 'RAM Stick')
	`

	rows, err := db.Query(query, ssd.ID, ssd.Wattage)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}

	return compatibleProducts, nil
}

func CompatibleWithGPU(id int) (compatibleProducts []models.Product,err error) {
	db := storage.GetDB()

	gpu, err := GetProductGPUByID(id)
	if err != nil {
		return nil, err
	}

	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_gm_slot cgm ON p.id = cgm.motherboard_id
		WHERE cgm.gpu_id = $1

		UNION

		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN compatible_gp_connector cgp ON p.id = cgp.power_supply_id
		WHERE cgp.gpu_id = $1

		UNION


		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $2

		UNION

		SELECT id, brand, model, current_price, stock_count, category, product_image
		FROM product
		WHERE category IN ('CPU', 'SSD', 'HDD', 'RAM Stick')
	`

	rows, err := db.Query(query, gpu.ID, gpu.Wattage)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return compatibleProducts, nil
}


func CompatibleWithHDD(id int) (compatibleProducts []models.Product,err error) {
	db := storage.GetDB()
	
	hdd, err := GetHDDByID(id)
	if err != nil {
		return nil, err
	}
	query := `
		SELECT p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, p.product_image
		FROM product p
		JOIN product_power_supply pps ON p.id = pps.product_id
		WHERE pps.supported_wattage >= $1

		UNION

		SELECT id, brand, model, current_price, stock_count, category, product_image
		FROM product
		WHERE category IN ('CPU', 'SSD', 'HDD', 'RAM Stick')
	`

	rows, err := db.Query(query, hdd.Wattage)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p models.Product
		err := rows.Scan(
			&p.ID,
			&p.Brand,
			&p.Model,
			&p.CurrentPrice,
			&p.StockCount,
			&p.Category,
			&p.ProductImage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product: %v", err)
		}
		compatibleProducts = append(compatibleProducts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}

	return compatibleProducts, nil
}

func GetProductInterfaceByID(id int) (models.ProductInterface, error) {
	baseProduct, err := GetProductByID(id)
	if err != nil {
		return nil, err
	}

	switch baseProduct.Category {
	case string(models.CategoryCPU):
		return GetCPUByID(baseProduct.ID)

	case string(models.CategoryGPU):
		return GetProductGPUByID(baseProduct.ID)

	case string(models.CategoryMotherboard):
		return GetProductMotherboardByID(baseProduct.ID)

	case string(models.CategoryRAMStick):
		return GetProductRAmStickByID(baseProduct.ID)

	case string(models.CategoryCooler):
		return GetCoolerByID(baseProduct.ID)

	case string(models.CategoryPowerSupply):
		return GetProductPowerSupplyByID(baseProduct.ID)

	case string(models.CategoryCase):
		return GetProductCaseByID(baseProduct.ID)

	case string(models.CategorySSD):
		return GetProductSSDByID(baseProduct.ID)

	case string(models.CategoryHDD):
		return GetHDDByID(baseProduct.ID)

	default:
		return nil, fmt.Errorf("unknown product category: %s", baseProduct.Category)
	}
}

