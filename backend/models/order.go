package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID         uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey;" json:"orderId"`
	MenuItem   []MenuItem `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE" json:"menuItem"`
	TotalPrice int        `json:"totalprice"`
	CreateTime time.Time
}

//紀錄 餐點數量、價格
type MenuItem struct {
	ID       int       `gorm:"primaryKey;AUTO_INCREMENT" json:"menuItemId"`
	OrderID  uuid.UUID `gorm:"type:uuid" json:"-"`
	MenuID   int       `gorm:"not null" json:"-"`
	Menu     Menu      `gorm:"constraint:OnDelete:CASCADE"`
	Quantity int       `gorm:"type:integer;not null" json:"quantity"`
	Subtotal int       `gorm:"type:integer;not null" json:"subtotal"`
}
