package common

type Yao struct {
	Image     string `json:"image"`     // 爻象
	Text      string `json:"text"`      // 爻辞
	ImageText string `json:"imageText"` // 象辞
}

// 八卦
type BaGua struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Char       string `json:"char"`
	YinYangArr []int  `json:"yinYangArr"`
	Tips       string `json:"tips"`
}

// 卦象： 六十四卦
type GuaImage struct {
	Index   int         `json:"index"`   // 卦序
	Wai     BaGua       `json:"wai"`     // 上卦，外卦
	Nei     BaGua       `json:"nei"`     // 下卦，内卦
	Name    string      `json:"name"`    // 卦名
	DuYin   string      `json:"duYin"`   // 读音
	Text    string      `json:"text"`    // 卦辞
	Extra   string      `json:"extra"`   // 额外信息
	YongYao Yao         `json:"yongYao"` // 如用九、用六
	Short   string      `json:"short"`   // 卦简介
	Desc    string      `json:"desc"`    // 介绍
	Yao     [GuaLen]Yao `json:"yao"`     // 六爻
}

var (
	BaGuaMap = map[int]BaGua{
		Qian: {
			Id:         Qian,
			Name:       "乾(天)",
			Char:       "☰",
			YinYangArr: []int{Yang, Yang, Yang},
			Tips:       "乾三连",
		},
		Kun: {
			Id:         Kun,
			Name:       "坤(地)",
			Char:       "☷",
			YinYangArr: []int{Yin, Yin, Yin},
			Tips:       "坤六断",
		},
		Zhen: {
			Id:         Zhen,
			Name:       "震(雷)",
			Char:       "☳",
			YinYangArr: []int{Yin, Yin, Yang},
			Tips:       "震仰盂", // 下实上虚，形似口朝上的钵盂
		},
		Gen: {
			Id:         Gen,
			Name:       "艮(山)",
			Char:       "☶",
			YinYangArr: []int{Yang, Yin, Yin},
			Tips:       "艮覆碗", // 上实，下虚，形似扣着的碗
		},
		Li: {
			Id:         Li,
			Name:       "離(火)",
			Char:       "☲",
			YinYangArr: []int{Yang, Yin, Yang},
			Tips:       "离中虚",
		},
		Kan: {
			Id:         Kan,
			Name:       "坎(水)",
			Char:       "☵",
			YinYangArr: []int{Yin, Yang, Yin},
			Tips:       "坎中满",
		},
		Dui: {
			Id:         Dui,
			Name:       "兌(澤)",
			Char:       "☱",
			YinYangArr: []int{Yin, Yang, Yang},
			Tips:       "兑上缺",
		},
		Xun: {
			Id:         Xun,
			Name:       "巽xùn(風)",
			Char:       "☴",
			YinYangArr: []int{Yang, Yang, Yin},
			Tips:       "巽下断",
		},
	}
)
