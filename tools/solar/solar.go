package solar

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/common"
	"strings"
	"time"

	"github.com/Tokumicn/theBookofChangesEveryDay/tools/animal"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/constellation"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/solarterm"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/utils"
)

// Solar 公历
type Solar struct {
	t                *time.Time
	CurrentSolarterm *solarterm.Solarterm
	PrevSolarterm    *solarterm.Solarterm
	NextSolarterm    *solarterm.Solarterm
}

var weekAlias = [...]string{
	"日", "一", "二", "三", "四", "五", "六",
}

// NewSolar 创建公历对象
func NewSolar(t *time.Time) *Solar {
	var c *solarterm.Solarterm
	p, n := solarterm.CalcSolarterm(t)
	if n.Index()-p.Index() == 1 {
		if p.IsInDay(t) {
			c = p
			p = p.Prev()
		}
		if n.IsInDay(t) {
			c = n
			p = c.Prev()
			n = c.Next()
		}
	}
	return &Solar{
		t:                t,
		CurrentSolarterm: c,
		PrevSolarterm:    p,
		NextSolarterm:    n,
	}
}

// IsLeep 是否为闰年
func (solar *Solar) IsLeep() bool {
	year := solar.t.Year()
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// WeekNumber 返回当前周次(周日为0, 周一为1...)
func (solar *Solar) WeekNumber() int64 {
	return int64(solar.t.Weekday())
}

// WeekAlias 返回当前周次(日, 一...)
func (solar *Solar) WeekAlias() string {
	return weekAlias[solar.WeekNumber()]
}

// Animal 返回年份生肖
func (solar *Solar) Animal() *animal.Animal {
	return animal.NewAnimal(utils.YearOrderMod(int64(solar.t.Year()-3), 12))
}

// Constellation 返回星座
func (solar *Solar) Constellation() *constellation.Constellation {
	return constellation.NewConstellation(solar.t)
}

// GetYear 年
func (solar *Solar) GetYear() int64 {
	return int64(solar.t.Year())
}

// GetMonth 月
func (solar *Solar) GetMonth() int64 {
	return int64(solar.t.Month())
}

// GetDay 日
func (solar *Solar) GetDay() int64 {
	return int64(solar.t.Day())
}

// GetHour 时
func (solar *Solar) GetHour() int64 {
	return int64(solar.t.Hour())
}

// GetMinute 分
func (solar *Solar) GetMinute() int64 {
	return int64(solar.t.Minute())
}

// GetSecond 秒
func (solar *Solar) GetSecond() int64 {
	return int64(solar.t.Second())
}

// GetNanosecond 毫秒
func (solar *Solar) GetNanosecond() int64 {
	return int64(solar.t.Nanosecond())
}

func (solar *Solar) Format(layout string) string {
	if len(layout) <= 0 {
		layout = common.DateTime
	}
	return solar.t.Format(layout)
}

func (solar *Solar) ToString() string {
	return fmt.Sprintf("%s [%s年 %s %s]",
		solar.t.Format(common.DateTime),
		solar.YearAlias(),
		solar.MonthAlias(),
		solar.DayAlias())
}

// YearAlias 汉字表示年(二零一八)
func (solar *Solar) YearAlias() string {
	s := fmt.Sprintf("%d", solar.t.Year())
	for i, replace := range common.NumberAlias {
		s = strings.Replace(s, fmt.Sprintf("%d", i), replace, -1)
	}
	return s
}

// MonthAlias 汉字表示月(八月, 闰六月)
func (solar *Solar) MonthAlias() string {
	pre := ""
	//if lunar.monthIsLeap {
	//	pre = "闰"
	//}
	return pre + common.SolarMonthAlias[solar.t.Month()-1] + "月"
}

// DayAlias 汉字表示日(初一, 初十...)
func (solar *Solar) DayAlias() (alias string) {
	switch solar.t.Day() {
	case 10:
		alias = "初十"
	case 20:
		alias = "二十"
	case 30:
		alias = "三十"
	default:
		alias = common.DateAlias[(int)(solar.t.Day()/10)] + common.NumberAlias[solar.t.Day()%10]
	}
	return
}

// Equals 返回两个对象是否相同
func (solar *Solar) Equals(b *Solar) bool {
	return solar.GetYear() == b.GetYear() &&
		solar.GetMonth() == b.GetMonth() &&
		solar.GetDay() == b.GetDay() &&
		solar.GetHour() == b.GetHour() &&
		solar.GetMinute() == b.GetMinute() &&
		solar.GetSecond() == b.GetSecond() &&
		solar.GetNanosecond() == b.GetNanosecond()
}
