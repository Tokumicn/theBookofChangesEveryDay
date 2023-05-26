package Suan

import (
	"time"
)

type AdminUser struct {
	Id   int64  // 唯一ID
	Name string // 姓名
}

type User struct {
	Id            int64     // 唯一ID
	Name          string    // 姓名
	Age           int       // 年龄
	Sex           bool      // 男(true) 女(false)
	Birthday      time.Time // 阳历生日 YYYY-MM-DD HH:mm
	LunarBirthday time.Time // 农历生日 YYYY-MM-DD HH:mm
}
