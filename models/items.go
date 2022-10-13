package models

type Item struct {
	Item_id     uint   `gorm:"primaryKey" json:"id"`
	Item_code   string `gorm:"not null; unique;type varchar(191)" json:"itemCode"`
	Description string `gorm:"type text" json:"description"`
	Quantity    uint   `gorm:"not null; type int" json:"quantity"`
	OrderID     uint
}
