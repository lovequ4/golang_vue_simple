package models

type Menu struct {
	ID          int    `gorm:"primaryKey;AUTO_INCREMENT" json:"menuId"`
	MenuName    string `gorm:"type:varchar;unique;not null" json:"menuname"`
	Description string `gorm:"type:varchar;unique;not null" json:"description"`
	Price       int    `gorm:"type:integer;not null" json:"price"`
	MenuImg     string `gorm:"type:varchar;unique;not null" json:"menuimg"`
}
