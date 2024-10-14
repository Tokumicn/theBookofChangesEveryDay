package models

// Stuff 物品
type Stuff struct {
	QName  string  // 搜索名 一类商品总名称 如：月亮石
	Name   string  // 实际商品名
	Order  int     // 顺序
	MaxVal ValueV1 // 最高值
	MinVal ValueV1 // 最高值
}

type ValueV1 struct {
	MHValue int     // MH W为单位
	RMValue float32 // RM yuan为单位
}

// 搜索名：月亮石/月亮/月
// 月亮石  1级    8W 4元
// 元宵   攻击   64W ?元
// 低兽决 连击  100W  ?元
// 高兽决 高连击 500W ?元
// 灵饰指南 100级佩饰 55W ?元
