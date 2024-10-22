package domain

import "time"

type Users struct {
	ID int64 `gorm:"primaryKey;autoIncrement:true"`
	Username string `gorm:"size:255"` 
	Email string `gorm:"size:255"`
	Password string `gorm:"size:255"`
	Address string `gorm:"size:255"`
	Company string `gorm:"size:255"`
  
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
