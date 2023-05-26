package main

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/core/Suan"
	"time"
)

func main() {
	// 提供生日以及性别
	u := Suan.User{
		Id:       1,
		Name:     "Test",
		Age:      31,
		Sex:      true,
		Birthday: time.Date(1992, 3, 29, 4, 30, 0, 0, time.Local),
	}

	suan := Suan.NewSuan(u)
	// 干支配数
	// 求得天数地数
	suan.Do()

	// 获取农历干支
	ganzhiStr := suan.Calendar.Lunar.ToString()
	fmt.Println("干支: ", ganzhiStr)

	// 获得先天卦
	fmt.Println("先天卦: ", suan.XianTianGua)

	// 获取元堂爻

	// 获取后天卦

	// 获取大运卦

	// 获取流年卦
}
