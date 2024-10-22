package domain

import "time"

type Roles struct {
	ID   int64  `gorm:"primaryKey;autoIncrement:true"`
	Name string `gorm:"size:255"`
	
	CreatedAt time.Time `gorm:"autoCreateTime"`
    UpdatedAt time.Time `gorm:"autoUpdateTime"`
}