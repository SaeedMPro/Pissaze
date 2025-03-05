package repositories

import (
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetCompatibleByID(productID int) ([]models.Product, error) {
	db := storage.GetDB()

	query := `
        SELECT p.id, p.brand, p.model, p.current_price, 
               p.stock_count, p.category, p.product_image
        FROM (
            -- CPU-Cooler compatibility
            SELECT cooler_id AS compatible_id FROM compatible_cc_socket WHERE cpu_id = $1
            UNION
            SELECT cpu_id FROM compatible_cc_socket WHERE cooler_id = $1
            
            UNION ALL
            
            -- CPU-Motherboard compatibility
            SELECT motherboard_id FROM compatible_mc_socket WHERE cpu_id = $1
            UNION
            SELECT cpu_id FROM compatible_mc_socket WHERE motherboard_id = $1
            
            UNION ALL
            
            -- RAM-Motherboard compatibility
            SELECT motherboard_id FROM compatible_rm_slot WHERE ram_id = $1
            UNION
            SELECT ram_id FROM compatible_rm_slot WHERE motherboard_id = $1
            
            UNION ALL
            
            -- GPU-PowerSupply compatibility
            SELECT power_supply_id FROM compatible_gp_connector WHERE gpu_id = $1
            UNION
            SELECT gpu_id FROM compatible_gp_connector WHERE power_supply_id = $1
            
            UNION ALL
            
            -- SSD-Motherboard compatibility
            SELECT motherboard_id FROM compatible_sm_slot WHERE ssd_id = $1
            UNION
            SELECT ssd_id FROM compatible_sm_slot WHERE motherboard_id = $1
            
            UNION ALL
            
            -- GPU-Motherboard compatibility
            SELECT motherboard_id FROM compatible_gm_slot WHERE gpu_id = $1
            UNION
            SELECT gpu_id FROM compatible_gm_slot WHERE motherboard_id = $1
        ) AS compat
        JOIN product p ON p.id = compat.compatible_id`

	rows, err := db.Query(query, productID)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ID, &product.Brand, &product.Model, &product.CurrentPrice,
			&product.StockCount, &product.Category, &product.ProductImage,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return products, nil
} 

func GetAllCompatibleCPUwithCoolerBySocket() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			pc.id, pc.brand, pc.model, pc.current_price, pc.stock_count, pc.category, pc.product_image,
			c.generation, c.microarchitecture, c.num_cores, c.num_threads, c.base_frequency, c.boost_frequency, c.max_memory_limit, c.wattage,
			pcl.id, pcl.brand, pcl.model, pcl.current_price, pcl.stock_count, pcl.category, pcl.product_image,
			cl.cooling_method, cl.fan_size, cl.max_rotational_speed, cl.wattage, cl.depth, cl.height, cl.width
		FROM compatible_cc_socket mc
		JOIN product_cpu c ON mc.cpu_id = c.product_id
		JOIN product_cooler cl ON mc.cooler_id = cl.product_id
		JOIN product pc ON pc.id = mc.cpu_id
		JOIN product pcl ON pcl.id = mc.cooler_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var cpu models.ProductCPU
		var cooler models.ProductCooler

		err := rows.Scan(
			&cpu.ID, &cpu.Brand, &cpu.Model, &cpu.CurrentPrice, &cpu.StockCount, &cpu.Category, &cpu.ProductImage,
			&cpu.Generation, &cpu.Microarchitecture, &cpu.NumCores, &cpu.NumThreads, &cpu.BaseFrequency, &cpu.BoostFrequency, &cpu.MaxMemoryLimit, &cpu.Wattage,

			&cooler.ID, &cooler.Brand, &cooler.Model, &cooler.CurrentPrice, &cooler.StockCount, &cooler.Category, &cooler.ProductImage,
			&cooler.CoolingMethod, &cooler.FanSize, &cooler.MaxRotationalSpeed, &cooler.Wattage, &cooler.Depth, &cooler.Height, &cooler.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  cpu,
			SecondProduct: cooler,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}

func GetAllCompatibleCPUwithMotherboardBySocket() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			pc.id, pc.brand, pc.model, pc.current_price, pc.stock_count, pc.category, pc.product_image,
			c.generation, c.microarchitecture, c.num_cores, c.num_threads, c.base_frequency, c.boost_frequency, c.max_memory_limit, c.wattage,
			pm.id, pm.brand, pm.model, pm.current_price, pm.stock_count, pm.category, pm.product_image,
			m.chipset_name, m.num_memory_slots, m.memory_speed_range, m.wattage, m.depth, m.height, m.width
		FROM compatible_mc_socket mc
		JOIN product_cpu c ON mc.cpu_id = c.product_id
		JOIN product_motherboard m ON mc.motherboard_id = m.product_id
		JOIN product pc ON pc.id = mc.cpu_id
		JOIN product pm ON pm.id = mc.motherboard_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var cpu models.ProductCPU
		var motherboard models.ProductMotherboard

		err := rows.Scan(
			&cpu.ID, &cpu.Brand, &cpu.Model, &cpu.CurrentPrice, &cpu.StockCount, &cpu.Category, &cpu.ProductImage,
			&cpu.Generation, &cpu.Microarchitecture, &cpu.NumCores, &cpu.NumThreads, &cpu.BaseFrequency, &cpu.BoostFrequency, &cpu.MaxMemoryLimit, &cpu.Wattage,

			&motherboard.ID, &motherboard.Brand, &motherboard.Model, &motherboard.CurrentPrice, &motherboard.StockCount, &motherboard.Category, &motherboard.ProductImage,
			&motherboard.ChipsetName, &motherboard.NumMemorySlots, &motherboard.MemorySpeedRange, &motherboard.Wattage, &motherboard.Depth, &motherboard.Height, &motherboard.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  cpu,
			SecondProduct: motherboard,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}

func GetAllCompatibleRAMstickWithMotherboardBySlot() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			prm.id, prm.brand, prm.model, prm.current_price, prm.stock_count, prm.category, prm.product_image,
			rs.generation, rs.capacity, rs.frequency, rs.wattage, rs.depth, rs.height, rs.width,
			pm.id, pm.brand, pm.model, pm.current_price, pm.stock_count, pm.category, pm.product_image,
			m.chipset_name, m.num_memory_slots, m.memory_speed_range, m.wattage, m.depth, m.height, m.width
		FROM compatible_rm_slot rm
		JOIN product_ram_stick rs ON rm.ram_id = rs.product_id
		JOIN product_motherboard m ON rm.motherboard_id = m.product_id
		JOIN product prm ON prm.id = rm.ram_id
		JOIN product pm ON pm.id = rm.motherboard_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var ramStick models.ProductRAMStick
		var motherboard models.ProductMotherboard

		err := rows.Scan(
			&ramStick.ID, &ramStick.Brand, &ramStick.Model, &ramStick.CurrentPrice, &ramStick.StockCount, &ramStick.Category, &ramStick.ProductImage,
			&ramStick.Generation, &ramStick.Capacity, &ramStick.Frequency, &ramStick.Wattage, &ramStick.Depth, &ramStick.Height, &ramStick.Width,

			&motherboard.ID, &motherboard.Brand, &motherboard.Model, &motherboard.CurrentPrice, &motherboard.StockCount, &motherboard.Category, &motherboard.ProductImage,
			&motherboard.ChipsetName, &motherboard.NumMemorySlots, &motherboard.MemorySpeedRange, &motherboard.Wattage, &motherboard.Depth, &motherboard.Height, &motherboard.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  ramStick,
			SecondProduct: motherboard,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}

func GetAllCompatibleGPUwithPowerSupplyByConnector() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			pg.id, pg.brand, pg.model, pg.current_price, pg.stock_count, pg.category, pg.product_image,
			g.ram_size, g.clock_speed, g.num_fans, g.wattage, g.depth, g.height, g.width,
			pps.id, pps.brand, pps.model, pps.current_price, pps.stock_count, pps.category, pps.product_image,
			ps.supported_wattage, ps.depth, ps.height, ps.width
		FROM compatible_gp_connector gp
		JOIN product_gpu g ON gp.gpu_id = g.product_id
		JOIN product_power_supply ps ON gp.power_supply_id = ps.product_id
		JOIN product pg ON pg.id = gp.gpu_id
		JOIN product pps ON pps.id = gp.power_supply_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var gpu models.ProductGPU
		var powerSupply models.ProductPowerSupply

		err := rows.Scan(
			&gpu.ID, &gpu.Brand, &gpu.Model, &gpu.CurrentPrice, &gpu.StockCount, &gpu.Category, &gpu.ProductImage,
			&gpu.RAMSize, &gpu.ClockSpeed, &gpu.NumFans, &gpu.Wattage, &gpu.Depth, &gpu.Height, &gpu.Width,

			&powerSupply.ID, &powerSupply.Brand, &powerSupply.Model, &powerSupply.CurrentPrice, &powerSupply.StockCount, &powerSupply.Category, &powerSupply.ProductImage,
			&powerSupply.SupportedWattage, &powerSupply.Depth, &powerSupply.Height, &powerSupply.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  gpu,
			SecondProduct: powerSupply,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}

func GetAllCompatibleSSDwithMotherboardBySlot() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			pg.id, pg.brand, pg.model, pg.current_price, pg.stock_count, pg.category, pg.product_image,
			g.ram_size, g.clock_speed, g.num_fans, g.wattage, g.depth, g.height, g.width,
			pm.id, pm.brand, pm.model, pm.current_price, pm.stock_count, pm.category, pm.product_image,
			m.chipset_name, m.num_memory_slots, m.memory_speed_range, m.wattage, m.depth, m.height, m.width
		FROM compatible_gm_slot gm
		JOIN product_gpu g ON gm.gpu_id = g.product_id
		JOIN product_motherboard m ON gm.motherboard_id = m.product_id
		JOIN product pg ON pg.id = gm.gpu_id
		JOIN product pm ON pm.id = gm.motherboard_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var ssd models.ProductSSD
		var motherboard models.ProductMotherboard

		err := rows.Scan(
			&ssd.ID, &ssd.Brand, &ssd.Model, &ssd.CurrentPrice, &ssd.StockCount, &ssd.Category, &ssd.ProductImage,
			&ssd.Capacity, &ssd.Wattage,

			&motherboard.ID, &motherboard.Brand, &motherboard.Model, &motherboard.CurrentPrice, &motherboard.StockCount, &motherboard.Category, &motherboard.ProductImage,
			&motherboard.ChipsetName, &motherboard.NumMemorySlots, &motherboard.MemorySpeedRange, &motherboard.Wattage, &motherboard.Depth, &motherboard.Height, &motherboard.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  ssd,
			SecondProduct: motherboard,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}

func GetAllCompatibleGPUwithMotherboardBySlot() ([]models.Compatible, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			ps.id, ps.brand, ps.model, ps.current_price, ps.stock_count, ps.category, ps.product_image,
			s.capacity, s.wattage,
			pm.id, pm.brand, pm.model, pm.current_price, pm.stock_count, pm.category, pm.product_image,
			m.chipset_name, m.num_memory_slots, m.memory_speed_range, m.wattage, m.depth, m.height, m.width
		FROM compatible_sm_slot sm
		JOIN product_ssd s ON sm.ssd_id = s.product_id
		JOIN product_motherboard m ON sm.motherboard_id = m.product_id
		JOIN product ps ON ps.id = sm.ssd_id
		JOIN product pm ON pm.id = sm.motherboard_id`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	var compatibles []models.Compatible
	for rows.Next() {
		var gpu models.ProductGPU
		var motherboard models.ProductMotherboard

		err := rows.Scan(
			&gpu.ID, &gpu.Brand, &gpu.Model, &gpu.CurrentPrice, &gpu.StockCount, &gpu.Category, &gpu.ProductImage,
			&gpu.RAMSize, &gpu.ClockSpeed, &gpu.NumFans, &gpu.Wattage, &gpu.Depth, &gpu.Height, &gpu.Width,

			&motherboard.ID, &motherboard.Brand, &motherboard.Model, &motherboard.CurrentPrice, &motherboard.StockCount, &motherboard.Category, &motherboard.ProductImage,
			&motherboard.ChipsetName, &motherboard.NumMemorySlots, &motherboard.MemorySpeedRange, &motherboard.Wattage, &motherboard.Depth, &motherboard.Height, &motherboard.Width,
		)

		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		compatibles = append(compatibles, models.Compatible{
			FirstProduct:  gpu,
			SecondProduct: motherboard,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return compatibles, nil
}