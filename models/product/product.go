package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	WishId    int            `json:"wishId"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
