package common

const (
	DateTime = "2006-01-02 15:04:05"
	DateOnly = "2006-01-02"
	TimeOnly = "15:04:05"
)

var NumberAlias = [...]string{
	"零", "一", "二", "三", "四",
	"五", "六", "七", "八", "九",
}

var LunarMonthAlias = [...]string{
	"正", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "冬", "腊",
}

var SolarMonthAlias = [...]string{
	"一", "二", "三", "四", "五", "六",
	"七", "八", "九", "十", "十一", "十二",
}

var DateAlias = [...]string{
	"初", "十", "廿", "卅",
}

// 先天卦序
var XianTianIndexMap = map[int]string{
	1: "乾",
	2: "兑",
	3: "离",
	4: "震",
	5: "巽",
	6: "坎",
	7: "艮",
	8: "坤",
}

// 后天卦序
// 坎一、坤二、震三、巽四、五屮宫、乾六、七兑、八艮、离九
var HouTianIndexMap = map[int]string{
	1: "坎",
	2: "坤",
	3: "震",
	4: "巽",
	5: "中宫",
	6: "乾",
	7: "兑",
	8: "艮",
	9: "离",
}
