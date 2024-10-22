package domain

import "time"

type Invoices struct {
	ID           int64         `gorm:"primaryKey;autoIncrement:true"`
	CommissionID int64         `gorm:"index"`
	Commission   Commissions   `gorm:"foreignKey:CommissionID"`
	Status       InvoiceStatus `gorm:"type:InvoiceStatus"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
}

type InvoiceStatus string

const (
	BelumBayar   InvoiceStatus = "belumbayar"
	SudahBayarDP InvoiceStatus = "sudahbayardp"
	SudahLunas   InvoiceStatus = "sudahlunas"
)
