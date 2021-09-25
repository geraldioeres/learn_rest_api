package wish

import (
	"learn_api/models/product"
	"time"

	"gorm.io/gorm"
)

type Wish struct {
	Id        int            `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`

	Product []product.Product `json:"product" gorm:"foreignKey:WishId"`
}
