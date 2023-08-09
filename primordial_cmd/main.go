package main

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/core/Suan"
	"time"
)

func main() {
	// 提供生日以及性别
	u1 := Suan.User{
		Id:       1,
		Name:     "Test",
		Sex:      true,
		Birthday: time.Date(1992, 3, 29, 4, 30, 0, 0, time.Local),
	}

	Qiu(u1)
	//fmt.Println("=============================================")

	// 2
	//u2 := Suan.User{
	//	Id:       2,
	//	Name:     "Test2",
	//	Sex:      false,
	//	Birthday: time.Date(1962, 11, 10, 18, 00, 0, 0, time.Local),
	//}
	//
	//Qiu(u2)
}

func Qiu(u Suan.User) {
	suan := Suan.NewSuan(u)
	// 干支配数
	// 求得天数地数
	suan.Do()

	// 获取阳历
	solarStr := suan.Calendar.Solar.ToString()
	fmt.Println("阳历: ", solarStr)

	// 获取阴历
	lunarStr := suan.Calendar.Lunar.ToString()
	fmt.Println("阴历: ", lunarStr)

	// 干支
	ganzhiStr := suan.Calendar.Ganzhi.ToString()
	fmt.Println("干支: ", ganzhiStr)

	// 获取配数
	peishuStr := suan.PeiShu.PeiShu2String()
	fmt.Println("配数: ", peishuStr)

	// 获得先天卦
	fmt.Println("先天卦: ", suan.XianTianGua)

	// 获取元堂爻

	// 获取后天卦

	// 获取大运卦

	// 获取流年卦

	// 三数定卦
	suan.ThreeDigit([]int{280, 980, 324})
	fmt.Println("三数卦: ", suan.SanShuGua)
}
