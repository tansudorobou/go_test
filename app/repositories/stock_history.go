package repositories

import (
	"app/models"
	"database/sql"
)

// ユーザーデータの読み出し
func GetStocks(db *sql.DB) ([]*models.Stock, error) {

	// currentTime := time.Now()
	// date := currentTime.Format("2006-01-02")

	query := `
	SELECT
	insert_dt,
	item_code,
	item_name,
	factory_code,
	factory_name,
	customer_code,
	customer_name,
	SUM(stock_quantity) 
	FROM stock_history
	GROUP BY insert_dt, item_code, item_name, factory_code, factory_name,customer_code,customer_name
	`
	// WHERE insert_dt = :$1
	// クエリの実行
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var stocks []*models.Stock

	for rows.Next() {
		stock := &models.Stock{}
		err := rows.Scan(
			&stock.InsertDt,
			&stock.ItemCode,
			&stock.ItemName,
			&stock.FactoryCode,
			&stock.FactoryName,
			&stock.CustomerCode,
			&stock.CustomerName,
			&stock.StockQuantity,
		)
		if err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stocks, nil
}
