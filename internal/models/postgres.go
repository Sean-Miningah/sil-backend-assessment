package models

type Category struct {
	ID       uint       `json:"id" gorm:"primaryKey"`
	Name     string     `json:"name"`
	ParentID *uint      `json:"parent_id"`
	Children []Category `json:"children" gorm:"foreignKey:ParentID"`
}

type Product struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	CategoryID uint    `json:"category_id"`
}
