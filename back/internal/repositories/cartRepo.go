package repositories

import (
	"fmt"

	"github.com/pissaze/internal/models"
	"github.com/pissaze/internal/storage"
)

func GetShoppingCartByClientID(clientID int) ([]models.ShoppingCart, error) {
	db := storage.GetDB()

	query := `
		SELECT s.cart_number, s.client_id, s.cart_status
		FROM shopping_cart s
		WHERE s.client_id = $1`

	rows, err := db.Query(query, clientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var shoppingCarts []models.ShoppingCart
	for rows.Next() {
		var shoppingCart models.ShoppingCart

		err := rows.Scan(
			&shoppingCart.CartNumber, &shoppingCart.ClientID, &shoppingCart.CartStatus,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		shoppingCarts = append(shoppingCarts, shoppingCart)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return shoppingCarts, nil
}

func GetLockedShoppingCartByClientID(clientID int, limit int) ([]models.LockedShoppingCart, error) {
	db := storage.GetDB()

	query := `
		SELECT 
			ls.locked_number, ls.cart_number, ls.client_id, ls.time_stamp
		FROM locked_shopping_cart ls
		JOIN issued_for i ON 
			ls.locked_number = i.locked_number AND
			ls.cart_number = i.cart_number AND
			ls.client_id = i.client_id
		WHERE ls.client_id = $1
		ORDER BY ls.time_stamp
		LIMIT $2`

	rows, err := db.Query(query, clientID, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lockedShoppingCarts []models.LockedShoppingCart
	for rows.Next() {
		var lockedShoppingCart models.LockedShoppingCart

		err := rows.Scan(
			&lockedShoppingCart.LockedCartNumber, &lockedShoppingCart.CartNumber,
			&lockedShoppingCart.ClientID, &lockedShoppingCart.TimeStamp,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		lockedShoppingCarts = append(lockedShoppingCarts, lockedShoppingCart)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration: %v", err)
	}

	return lockedShoppingCarts, nil
}

func GetProductsInLockedShoppingCart(ls *models.LockedShoppingCart) error {
	db := storage.GetDB()
	
	query := `
        SELECT 
            p.id, p.brand, p.model, p.current_price, p.stock_count, p.category, a.quantity, a.cart_price
        FROM added_to a
        JOIN product p ON a.product_id = p.id
        WHERE a.client_id = $1
          AND a.cart_number = $2
          AND a.locked_number = $3`

	rows, err := db.Query(query, ls.ClientID, ls.CartNumber, ls.LockedCartNumber)
	if err != nil {
		return fmt.Errorf("failed to query products: %w", err)
	}
	defer rows.Close()

	var products []models.ProductShoppingCart
	totalPrice := 0.0
	for rows.Next() {
		var product models.ProductShoppingCart

		err := rows.Scan(
			&product.Product.ID, &product.Product.Brand, &product.Product.Model, &product.Product.CurrentPrice,
			&product.Product.StockCount, &product.Product.Category, &product.Quantity, &product.CartPrice,
		)

		if err != nil {
			return fmt.Errorf("failed to scan product: %w", err)
		}

		totalPrice += float64(product.Quantity) * product.CartPrice 
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return fmt.Errorf("row iteration error: %w", err)
	}

	ls.TotalPrice = totalPrice
	ls.Products = products
	return nil
}

func GetCurrentMonthVIPProfit(clientID int) (float64, error) {
    db := storage.GetDB()

    query := `
        SELECT COALESCE(SUM(adt.cart_price) * 0.15, 0)
        FROM vip_client vc
        JOIN issued_for ifo ON vc.client_id = ifo.client_id
        JOIN transaction t ON ifo.tracking_code = t.tracking_code
        JOIN added_to adt ON ifo.client_id = adt.client_id 
            AND ifo.cart_number = adt.cart_number 
            AND ifo.locked_number = adt.locked_number
        WHERE vc.client_id = $1
          AND t.transaction_status = 'Successful'
          AND t.time_stamp >= DATE_TRUNC('month', CURRENT_DATE)
          AND t.time_stamp <= NOW()
          AND vc.expiration_time >= t.time_stamp`

    var cashback float64
    err := db.QueryRow(query, clientID).Scan(&cashback)
    
    if err != nil {
        return 0, fmt.Errorf("failed to calculate VIP profit: %w", err)
    }
    
    return cashback, nil
}