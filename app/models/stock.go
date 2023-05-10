package models

import "time"

type Stock struct {
	InsertDt      time.Time `json:"insertDt"`
	ItemCode      string    `json:"itemCode"`
	ItemName      string    `json:"itemName"`
	FactoryCode   string    `json:"factoryCode"`
	FactoryName   string    `json:"factoryName"`
	CustomerCode  string    `json:"customerCode"`
	CustomerName  string    `json:"customerName"`
	StockQuantity float64   `json:"stockQuantity"`
}
