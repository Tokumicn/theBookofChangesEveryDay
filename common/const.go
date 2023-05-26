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
