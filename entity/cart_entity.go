package entity

type Cart struct {
	ID     uint       `gorm:"primaryKey" json:"-"`
	Items  []CartItem `json:"items"`
	UserID uint       `json:"-"`
	Total  float64    `json:"total_price"`
}

type CartItem struct {
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
	Quantity  int     `gorm:"default:1" json:"quantity"`
	Total     float64 `json:"total"`
	ProductID uint    `json:"-"`
	CartID    uint    `json:"-"`
}

type AddProduct struct {
	ProductID   uint `json:"product_id"`
	IsCheckList bool `json:"is_checklist"`
}

type DeleteProductFromCart struct {
	ProductID uint `json:"product_id"`
}
