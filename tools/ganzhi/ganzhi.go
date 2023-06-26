package ganzhi

import (
	"fmt"
	"time"

	"github.com/Tokumicn/theBookofChangesEveryDay/tools/animal"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/gan"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/solarterm"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/utils"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/zhi"
)

// Ganzhi 干支历
type Ganzhi struct {
	t             *time.Time
	YearGan       *gan.Gan
	YearZhi       *zhi.Zhi
	MonthGan      *gan.Gan
	MonthZhi      *zhi.Zhi
	DayGan        *gan.Gan
	DayZhi        *zhi.Zhi
	HourGan       *gan.Gan
	HourZhi       *zhi.Zhi
	PrevSolarterm *solarterm.Solarterm
	NextSolarterm *solarterm.Solarterm
}

// NewGanzhi 创建干支历对象
func NewGanzhi(t *time.Time) *Ganzhi {
	year := int64(t.Year())
	if !isSupported(year) {
		return nil
	}
	if t.Unix() < solarterm.SpringTimestamp(year) {
		year--
	}
	yearGan := gan.NewGan(utils.YearOrderMod(year-3, 10))
	yearZhi := zhi.NewZhi(utils.YearOrderMod(year-3, 12))

	p, n := solarterm.CalcSolarterm(t)

	i := p.Index()
	if n.Index()-p.Index() == 2 {
		i++
	}
	i = utils.YearOrderMod((i%24)/2, 12)

	monthZhi := zhi.NewZhi(utils.YearOrderMod(i+2, 12))
	monthGan := gan.NewGan(utils.YearOrderMod(i+yearGan.Order()*2, 10))

	begin := time.Date(solarterm.SolartermFromYear, 1, 1, 0, 0, 0, 0, time.Local)
	seconds := t.Sub(begin).Seconds()
	dayOrder := utils.YearOrderMod(int64(seconds/86400)+31, 60)

	dayGan := gan.NewGan(utils.YearOrderMod(dayOrder, 10))
	dayZhi := zhi.NewZhi(utils.YearOrderMod(dayOrder, 12))

	hourZhi := zhi.NewZhi(utils.YearOrderMod(int64(((t.Hour()+1)/2)+1), 12))
	hourGan := gan.NewGan(utils.YearOrderMod(hourZhi.Order()-2+dayGan.Order()*2, 10))

	return &Ganzhi{
		t:             t,
		YearGan:       yearGan,
		YearZhi:       yearZhi,
		MonthGan:      monthGan,
		MonthZhi:      monthZhi,
		DayGan:        dayGan,
		DayZhi:        dayZhi,
		HourGan:       hourGan,
		HourZhi:       hourZhi,
		PrevSolarterm: p,
		NextSolarterm: n,
	}
}

// Animal 返回年份生肖
func (gz *Ganzhi) Animal() *animal.Animal {
	return animal.NewAnimal(gz.YearZhi.Order())
}

// YearGanzhiAlias 年干支(甲子, 乙丑...)
func (gz *Ganzhi) YearGanzhiAlias() string {
	return gz.YearGan.Alias() + gz.YearZhi.Alias()
}

// MonthGanzhiAlias 月干支(甲子, 乙丑...)
func (gz *Ganzhi) MonthGanzhiAlias() string {
	return gz.MonthGan.Alias() + gz.MonthZhi.Alias()
}

// DayGanzhiAlias 日干支(甲子, 乙丑...)
func (gz *Ganzhi) DayGanzhiAlias() string {
	return gz.DayGan.Alias() + gz.DayZhi.Alias()
}

// HourGanzhiAlias 时干支(甲子, 乙丑...)
func (gz *Ganzhi) HourGanzhiAlias() string {
	return gz.HourGan.Alias() + gz.HourZhi.Alias()
}

// YearGanzhiOrder 年干支六十甲子序数(1,2...)
func (gz *Ganzhi) YearGanzhiOrder() int64 {
	return ganzhiOrder(gz.YearGan.Order(), gz.YearZhi.Order())
}

// MonthGanzhiOrder 月干支六十甲子序数(1,2...)
func (gz *Ganzhi) MonthGanzhiOrder() int64 {
	return ganzhiOrder(gz.MonthGan.Order(), gz.MonthZhi.Order())
}

// DayGanzhiOrder 日干支六十甲子序数(1,2...)
func (gz *Ganzhi) DayGanzhiOrder() int64 {
	return ganzhiOrder(gz.DayGan.Order(), gz.DayZhi.Order())
}

// HourGanzhiOrder 时干支六十甲子序数(1,2...)
func (gz *Ganzhi) HourGanzhiOrder() int64 {
	return ganzhiOrder(gz.HourGan.Order(), gz.HourZhi.Order())
}

// Equals 返回两个对象是否相同
func (gz *Ganzhi) Equals(b *Ganzhi) bool {
	return gz.YearGanzhiOrder() == b.YearGanzhiOrder() &&
		gz.MonthGanzhiOrder() == b.MonthGanzhiOrder() &&
		gz.DayGanzhiOrder() == b.DayGanzhiOrder() &&
		gz.HourGanzhiOrder() == b.HourGanzhiOrder()
}

func (gz *Ganzhi) ToString() string {
	return fmt.Sprintf("[ %s年 %s月 %s日 %s时 ]",
		gz.YearGanzhiAlias(),
		gz.MonthGanzhiAlias(),
		gz.DayGanzhiAlias(),
		gz.HourGanzhiAlias())
}

func isSupported(year int64) bool {
	return solarterm.SolartermFromYear <= year && year < solarterm.SolartermToYear
}

func ganzhiOrder(ganOrder, zhiOrder int64) int64 {
	return utils.YearOrderMod(((ganOrder+10-zhiOrder)%10)/2*12+zhiOrder, 60)
}
