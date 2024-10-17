package models

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

// 物品信息 无特殊属性的通用物品 如：宝石、兽决、灵饰指南、书铁、珍珠 等
type Stuff struct {
	gorm.Model
	QName    string  `gorm:"column:q_name"`    // 搜索名 一类商品总名称 如：月亮石
	Name     string  `gorm:"column:name"`      // 实际商品名 TODO: name字段添加全表唯一索引
	Order    int     `gorm:"column:order"`     // 顺序
	ValMH    float32 `gorm:"column:val_mh"`    // MH W为单位
	ValRM    float32 `gorm:"column:val_rm"`    // RM yuan为单位
	RegionID int     `gorm:"column:region_id"` // 所在区 通过 map 翻译即可
}

type StuffLog struct {
	Stuff
}

// 搜索名：月亮石/月亮/月
// 月亮石  1级    8W 4元
// 元宵   攻击   64W ?元
// 低兽决 连击  100W  ?元
// 高兽决 高连击 500W ?元
// 灵饰指南 100级佩饰 55W ?元

func (pVal Stuff) ExistByQName(ctx context.Context, db *gorm.DB, qName string) (bool, uint, error) {

	stuff, err := pVal.FindByQName(ctx, db)
	if err == gorm.ErrRecordNotFound {
		return false, 0, nil
	}
	if err != nil {
		return false, 0, err
	}

	if stuff.ID > 0 {
		return true, stuff.ID, nil
	}

	return false, 0, nil
}

// 查询单个物品信息
func (pVal Stuff) FindByQName(ctx context.Context, db *gorm.DB) (Stuff, error) {
	res := Stuff{}
	if err := db.
		WithContext(ctx).
		Model(Stuff{}).
		Where("q_name = ?", pVal.QName).
		First(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return Stuff{}, nil
		}
		return Stuff{}, fmt.Errorf("find one stuff info by qname err: %v", err)
	}

	return res, nil
}

// 创建商品价格信息
func (pVal Stuff) Create(ctx context.Context, db *gorm.DB) (uint, error) {
	if err := db.
		WithContext(ctx).
		Create(&pVal).Error; err != nil {
		return 0, fmt.Errorf("create stuff info err: %v", err)
	}
	return pVal.ID, nil
}

// 更新商品信息
func (pVal Stuff) Update(ctx context.Context, db *gorm.DB) (uint, error) {

	updateMap := map[string]interface{}{}

	if pVal.ValMH > 0 {
		updateMap["val_mh"] = pVal.ValMH
	}

	if pVal.ValRM > 0 {
		updateMap["val_rm"] = pVal.ValRM
	}

	if err := db.WithContext(ctx).
		Model(Stuff{}).
		Where("id = ?", pVal.ID).
		Error; err != nil {
		return 0, fmt.Errorf("update product value by updates:[%s] err: %v", updateMap, err)
	}
	return pVal.ID, nil
}

// 获取列表 目前仅提供通过名称查询
func (pVal Stuff) List(ctx context.Context, db *gorm.DB, offset, limit int) (int64, []Stuff, error) {
	vals := make([]Stuff, 0)
	var total int64
	db = db.WithContext(ctx).
		Model(Stuff{})

	// 组名称查询
	if len(pVal.QName) > 0 {
		db = db.Where("q_name = ?", pVal.QName)
	}

	// 唯一名称查询
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

func (pVal Stuff) CreateStuffLog(ctx context.Context, db *gorm.DB) (uint, error) {
	if err := db.WithContext(ctx).
		Table("stuff_log").
		Create(&pVal).
		Error; err != nil {
		return 0, fmt.Errorf("create stuff log err: %v", err)
	}
	return pVal.ID, nil
}
