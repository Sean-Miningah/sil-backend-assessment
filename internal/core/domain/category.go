package domain

type Category struct {
	ID       uint       `json:"id"`
	Name     string     `json:"name"`
	ParentID *uint      `json:"parent_id"`
	Children []Category `json:"children" gorm:"foreingKey:ParentID"`
}
