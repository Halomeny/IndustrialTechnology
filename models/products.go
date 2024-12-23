package models

type Product struct {
	Name         string  `json:"Название" gorm:"column:name";primary_key`
	Price        float32 `json:"Цена" gorm:"column:price"`
	Description  string  `json:"Описание" gorm:"column:description";`
	DeletionMark bool    `json:"Пометка удаления" gorm:"column:deletionmark"`
}

func (Product) TableName() string {
	return "products"
}
