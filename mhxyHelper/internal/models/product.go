package models

import (
	"gorm.io/gorm"
)

type ProductValV1 struct {
	gorm.Model
	Name   string `gorm:"name"`     // 物品名
	MaxVal int    `gorm:"max_val"`  // 最高价
	MinVal int    `gorm:"min_val" ` // 最低价
}

type ProductLog struct {
	gorm.Model
	Name  string `gorm:"name"`  // 物品名
	Value int    `gorm:"value"` // 价值(MH币)
}
