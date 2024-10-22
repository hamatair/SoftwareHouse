package domain

type UserRoles struct {
	UserID int64 `gorm:"primaryKey"`
	RoleID int64 `gorm:"primaryKey"`
	User   Users `gorm:"foreignKey:UserID"`
	Role   Roles  `gorm:"foreignKey:RoleID"`
}
