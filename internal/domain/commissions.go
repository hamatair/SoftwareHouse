package domain

import "time"

type Commissions struct {
	ID        int64            `gorm:"primaryKey;autoIncrement:true"`
	UserID    int64            `gorm:"index"`
	User      Users            `gorm:"foreignKey:UserID"`
	Status    CommissionStatus `gorm:"type:commissionStatus"`
	CreatedAt time.Time        `gorm:"autoCreateTime"`
	UpdatedAt time.Time        `gorm:"autoUpdateTime"`
}

type CommissionStatus string

const (
	Pending   CommissionStatus = "pending"
	Approved  CommissionStatus = "approved"
	Completed CommissionStatus = "completed"
)
