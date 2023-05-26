package tools

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/calendar"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/lunar"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/solar"
	"testing"
	"time"
)

func TestTools(t *testing.T) {
	ti := time.Now()
	// 1. ByTimestamp
	// 时间戳
	c := calendar.ByTimestamp(ti.Unix())

	bytes, err := c.ToJSON()
	if err != nil {
		t.Log(err.Error())
	}

	fmt.Println(string(bytes))

	so := solar.NewSolar(&ti)
	lu := lunar.NewLunar(&ti)

	fmt.Println("SolarTimeFormat: " + so.Format(""))
	fmt.Println("LunarTimeFormat: " + lu.Format(""))
	fmt.Println("SolarToString: " + so.ToString())
	fmt.Println("LunarToString: " + lu.ToString())
	fmt.Println("干支: " + c.Ganzhi.ToString())
}
