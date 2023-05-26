package Suan

import (
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/calendar"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/ganzhi"
)

// 算
type Suan struct {
	User        User               // 用户信息
	Calendar    *calendar.Calendar // 日历计算
	PeiShu      ganzhi.PeiShu      // 干支配数
	XianTianGua string             // 先天卦
	HouTianGua  string             // 后天卦
}

func NewSuan(user User) Suan {

	// 构建万年历
	calend := calendar.BySolar(user.Birthday)
	// 构建干支配数
	peiShu := ganzhi.NewPeiShu(*calend.Ganzhi, user.Sex)

	return Suan{
		User:     user,
		Calendar: calend,
		PeiShu:   peiShu,
	}
}

// 算
func (s *Suan) Do() {
	// 获取天地数
	tian, di := s.PeiShu.GetTianDiShu()
	// 获取先天卦
	gua := s.PeiShu.GetHouTianGua(tian, di)
	s.XianTianGua = gua

	// 计算元堂爻

	// 获取后天卦
	s.HouTianGua = "TODO"

	// 计算流年
}
