package main

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/core/Suan"
	"github.com/Tokumicn/theBookofChangesEveryDay/core/db"
	"time"
)

func main() {

	// initGua64DB()
}

// initGua64DB
func initGua64DB() {

	// 获取当前程序所在目录路径
	//dir, err := os.Getwd()
	//if err != nil {
	//	fmt.Println("获取当前目录失败：", err)
	//	return
	//}
	//fmt.Println("当前程序所在目录路径：", dir)

	err := db.InitDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	suan := Suan.NewSuan(Suan.User{
		Id:       1,
		Name:     "Test",
		Age:      31,
		Sex:      true,
		Birthday: time.Date(1992, 3, 29, 4, 30, 0, 0, time.Local),
	})

	for _, gua := range suan.GetGua64Arr() {
		err = db.InsertGua64(gua)
		if err != nil {
			e := fmt.Errorf("InsertGua64 [%+v] err: %s", gua, err.Error())
			fmt.Println(e.Error())
		}
	}
}

func clearGua64DB() {

}
