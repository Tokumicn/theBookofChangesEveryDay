package models

import (
	"context"
	"fmt"
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

// 创建日志
func (plog *ProductLog) Create(ctx context.Context, db *gorm.DB) (uint, error) {
	if err := db.
		WithContext(ctx).
		Model(ProductLog{}).
		Create(plog).Error; err != nil {
		return 0, fmt.Errorf("create product log err: %v", err)
	}
	return plog.ID, nil
}

// 获取列表 目前仅提供通过名称查询
func (plog *ProductLog) ListByName(ctx context.Context, db *gorm.DB, offset, limit int) (int64, []ProductLog, error) {
	logs := make([]ProductLog, 0)
	var total int64
	db = db.WithContext(ctx).
		Model(ProductLog{}).
		Where("name = ?", plog.Name)

	if err := db.Count(&total).Error; err != nil {
		return -1, nil, fmt.Errorf("get all product logs count err: %v", err)
	}

	if err := db.
		WithContext(ctx).
		Where("name = ?", plog.Name).
		Offset(offset).
		Limit(limit).
		Find(&logs).Error; err != nil {
		return -1, nil, fmt.Errorf("get all product logs err: %v", err)
	}
	return total, logs, nil
}

// 创建商品价格信息
func (pVal *ProductValV1) Create(ctx context.Context, db *gorm.DB) (uint, error) {
	if err := db.
		WithContext(ctx).
		Model(ProductValV1{}).
		Create(pVal).Error; err != nil {
		return 0, fmt.Errorf("create product value err: %v", err)
	}
	return pVal.ID, nil
}

// 更新商品信息
func (pVal *ProductValV1) Update(ctx context.Context, db *gorm.DB) (uint, error) {

	updateMap := map[string]interface{}{}

	if len(pVal.Name) > 0 {
		updateMap["name"] = pVal.Name
	}

	if pVal.MinVal > 0 {
		updateMap["min_val"] = pVal.MinVal
	}

	if pVal.MaxVal > 0 {
		updateMap["max_val"] = pVal.MaxVal
	}

	if err := db.WithContext(ctx).
		Model(ProductValV1{}).
		Where("id = ?", pVal.ID).Error; err != nil {
		return 0, fmt.Errorf("update product value by updates:[%s] err: %v", updateMap, err)
	}
	return pVal.ID, nil
}

// 获取列表 目前仅提供通过名称查询
func (pVal *ProductValV1) List(ctx context.Context, db *gorm.DB, offset, limit int) (int64, []ProductValV1, error) {
	vals := make([]ProductValV1, 0)
	var total int64
	db = db.WithContext(ctx).
		Model(ProductLog{})

	// 名称查询
	if len(pVal.Name) > 0 {
		db = db.Where("name = ?", pVal.Name)
	}

	// id查询
	if pVal.ID > 0 {
		db = db.Where("id = ?", pVal.ID)
	}

	if err := db.Count(&total).Error; err != nil {
		return -1, nil, fmt.Errorf("get list product value count err: %v", err)
	}

	if err := db.
		WithContext(ctx).
		Offset(offset).
		Limit(limit).
		Find(&vals).Error; err != nil {
		return -1, nil, fmt.Errorf("get list product value  err: %v", err)
	}
	return total, vals, nil
}
