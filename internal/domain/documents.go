package domain

import "time"

type Documents struct {
	ID           int64         `gorm:"primaryKey;autoIncrement:true"`
	Type         DocumentTypes `gorm:"type:document_types"` // Enum
	CommissionID int64         `gorm:"index"`               // Relasi ke Commission
	Commission   Commissions   `gorm:"foreignKey:CommissionID"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
}

type DocumentTypes string

const (
	MOU       DocumentTypes = "mou"
	BAST      DocumentTypes = "bast"
	DP        DocumentTypes = "dp"
	Pelunasan DocumentTypes = "pelunasan"
)
