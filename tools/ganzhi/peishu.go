package ganzhi

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/common"
)

type PeiShu struct {
	PeiShuArr      []int  // 干支配数
	GanZhi         Ganzhi // 干支
	Sex            bool   // 男(true) 女(false)
	YearGanYinYang bool   // 年干阴(false) 阳(true)
}

// 天干配数
var ganOrderMap = map[string]int{
	"甲": 6,
	"乙": 2,
	"丙": 8,
	"丁": 7,
	"戊": 1,
	"己": 9,
	"庚": 3,
	"辛": 4,
	"壬": 6,
	"癸": 2,
}

// 地支配数
var zhiOrderMap = map[string][]int{
	"亥": {1, 6}, "子": {1, 6},
	"寅": {3, 8}, "卯": {3, 8},
	"巳": {2, 7}, "午": {2, 7},
	"申": {4, 9}, "酉": {4, 9},
	"辰": {5, 10}, "戌": {5, 10}, "丑": {5, 10}, "未": {5, 10},
}

// 天地数对应卦
var ganzhiOrderGuaMap = map[int]string{
	1:  "坎卦",
	2:  "坤卦",
	3:  "震卦",
	4:  "巽卦",
	6:  "乾卦",
	7:  "兑卦",
	8:  "艮卦",
	9:  "离卦",
	51: "艮卦", // 上元 男
	52: "坤卦", // 上元 女
	53: "艮卦", // 中元 阳男阴女
	54: "坤卦", // 中元 阴男阳女
	55: "离卦", // 下元 男
	56: "兑卦", // 下元 女
}

func NewPeiShu(ganzhi Ganzhi, sex bool) PeiShu {

	// 获取年干阴阳
	yearIsYang := getYearYinYang(ganzhi.YearGan.Alias())
	// 构建配数列表
	peiShuArr := buildPeiShu(ganzhi)

	return PeiShu{
		PeiShuArr:      peiShuArr,
		GanZhi:         ganzhi,
		Sex:            sex,
		YearGanYinYang: yearIsYang,
	}
}

func (gz *Ganzhi) PeiShu2String() string {
	return "TODO"
}

// 获取天数和地数
func (ps *PeiShu) GetTianDiShu() (int, int) {
	tianTotal, diTotal := 0, 0

	// 干支数，全部单数相加为天数，全部双数相加为地数
	for _, n := range ps.PeiShuArr {
		if common.IsEven(n) {
			diTotal += n
		} else {
			tianTotal += n
		}
	}

	tian := calc(tianTotal, 25)
	di := calc(diTotal, 30)
	return tian, di
}

// 计算天地数
// 天数: 25
// 地数: 30
func calc(total, seed int) int {
	//不足数，则除十不用，只用零位之数
	if total <= seed {
		// 特殊情况
		if total%10 == 0 {
			return total / 10
		}
		return total % 10
	} else {
		// 超过数，则除天地数不用，只用零位之数
		return (total % seed) % 10
	}
}

// 获取后天卦
func (ps *PeiShu) GetHouTianGua(tian, di int) string {

	tian, di = ps.GetTianDiShu()

	tianGua := getGua(tian, ps.Sex, ps.YearGanYinYang, ps.GanZhi.t.Year())
	diGua := getGua(di, ps.Sex, ps.YearGanYinYang, ps.GanZhi.t.Year())

	if ps.Sex {
		return fmt.Sprintf("%s - %s", tianGua, diGua)
	} else {
		return fmt.Sprintf("%s - %s", diGua, tianGua)
	}
}

const (
	ShangYuan = 111 // 上元
	ZhongYuan = 222 // 中元
	XiaYuan   = 333 // 下元
)

// 构建配数数组
func buildPeiShu(ganzhi Ganzhi) []int {
	peishuArr := make([]int, 0)

	// 年干
	peishuArr = append(peishuArr, ganOrderMap[ganzhi.YearGan.Alias()])
	// 年支
	peishuArr = append(peishuArr, zhiOrderMap[ganzhi.YearZhi.Alias()]...)

	// 月干
	peishuArr = append(peishuArr, ganOrderMap[ganzhi.MonthGan.Alias()])
	// 月支
	peishuArr = append(peishuArr, zhiOrderMap[ganzhi.MonthZhi.Alias()]...)

	// 日干
	peishuArr = append(peishuArr, ganOrderMap[ganzhi.DayGan.Alias()])
	// 日支
	peishuArr = append(peishuArr, zhiOrderMap[ganzhi.DayZhi.Alias()]...)

	// 时干
	peishuArr = append(peishuArr, ganOrderMap[ganzhi.HourGan.Alias()])
	// 时支
	peishuArr = append(peishuArr, zhiOrderMap[ganzhi.HourZhi.Alias()]...)

	return peishuArr
}

// 根据天地数获取卦
func getGua(number int, sexMan bool, yearYang bool, year int) string {
	// 5有特殊处理
	if number == 5 {
		return ganzhiOrderGuaMap[processOrder5(sexMan, yearYang, year)]
	} else {
		return ganzhiOrderGuaMap[number]
	}
}

// 获取年干阴阳
func getYearYinYang(yearGan string) bool {
	//六阳支：子寅辰午申戌，
	//六阴支：丑卯巳未酉亥.
	yearGanYinYangMap := map[string]bool{
		"子": true, "寅": true, "辰": true, "午": true, "申": true, "戌": true,
		"丑": false, "卯": false, "巳": false, "未": false, "酉": false, "亥": false,
	}

	return yearGanYinYangMap[yearGan]
}

// 处理天地数为5的特殊情况
func processOrder5(sexMan bool, yearYang bool, year int) int {
	guaOrder := 0

	// 年份 求元
	yuanTag := 0
	if year > 1864 && year <= 1923 {
		yuanTag = ShangYuan
	} else if year > 1924 && year <= 1983 {
		yuanTag = ZhongYuan
	} else if year > 1984 && year <= 2043 {
		yuanTag = XiaYuan
	}

	switch yuanTag {
	case ShangYuan: // 上元: 男艮卦 女坤卦
		if sexMan {
			guaOrder = 51
		} else {
			guaOrder = 52
		}

	case ZhongYuan: // 中元
		if (sexMan && yearYang) || ((!sexMan) && !yearYang) {
			// 阳男阴女[艮卦]
			guaOrder = 53
		} else if (sexMan && (!yearYang)) || ((!sexMan) && yearYang) {
			// 阴男阳女[坤卦]
			guaOrder = 54
		}

	case XiaYuan: // 下元: 男离卦 女兑卦
		if sexMan {
			guaOrder = 55
		} else {
			guaOrder = 56
		}
	}

	return guaOrder
}
