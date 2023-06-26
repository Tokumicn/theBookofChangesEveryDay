package Suan

import (
	"bufio"
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/calendar"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/ganzhi"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/utils"
	"io"
	"os"
	"strings"
	"time"
)

// 算
type Suan struct {
	User          User               // 用户信息
	Calendar      *calendar.Calendar // 日历计算
	PeiShu        ganzhi.PeiShu      // 干支配数
	XianTianGua   string             // 先天卦
	HouTianGua    string             // 后天卦
	gua64File2Map map[string]string  // 64卦映射
	gua64File2Arr []GuaKV            // 64卦顺序映射
}

func NewSuan(user User) Suan {
	// 计算年龄
	user.Age = time.Now().Year() - user.Birthday.Year()

	// 构建万年历
	calend := calendar.BySolar(user.Birthday)
	// 构建干支配数
	peiShu := ganzhi.NewPeiShu(*calend.Ganzhi, user.Sex)

	// 读取文件并初始化
	gua64Map, gua64Arr := init64GuaFile()

	return Suan{
		User:          user,
		Calendar:      calend,
		PeiShu:        peiShu,
		gua64File2Map: gua64Map,
		gua64File2Arr: gua64Arr,
	}
}

type GuaKV struct {
	Key string // 乾 - 乾
	Val string // 乾为天（乾卦）自强不息
}

func init64GuaFile() (map[string]string, []GuaKV) {

	resultMap := map[string]string{}
	resultArr := []GuaKV{}

	//打开文件
	file, _ := os.Open(`/Users/tomhxiao/gopath/src/github.com/TokumiCN/theBookofChangesEveryDay/core/db/64gua.txt`)
	defer file.Close()

	//创建文件的缓冲读取器
	reader := bufio.NewReader(file)

	tempGua := GuaKV{}

	for {
		//逐行读取
		lineBytes, _, err := reader.ReadLine()
		//读到文件末尾时退出循环
		if err == io.EOF {
			break
		}
		//当前行内容
		oneLineStr := string(lineBytes)

		// 第一行
		if strings.HasPrefix(oneLineStr, "第") {
			tempLineBytes, _, err := reader.ReadLine()
			//读到文件末尾时退出循环
			if err == io.EOF {
				break
			}
			tempGua.Val = string(tempLineBytes)

		} else if strings.Contains(oneLineStr, "（上") {
			subIndex := strings.Index(oneLineStr, "（上")
			subTag := oneLineStr[subIndex+6 : subIndex+15]
			splits := strings.Split(subTag, "上")
			key := fmt.Sprintf("%s - %s", splits[0], splits[1])
			tempGua.Key = key
			fmt.Println(key)

		} else if strings.Contains(oneLineStr, "（下") {
			subIndex := strings.Index(oneLineStr, "（下")
			subTag := oneLineStr[subIndex+6 : subIndex+15]

			seps := []string{"上", "下"}
			sepSplitMap := utils.SplitEnhance(subTag, seps)
			for _, sep := range seps {
				tempArr := sepSplitMap[sep]
				if len(tempArr) == 2 {
					tempKey := fmt.Sprintf("%s - %s", tempArr[0], tempArr[1])
					tempGua.Key = tempKey

					// 解析内容填充到map
					saveGua64Map(tempGua, resultMap)
					resultArr = append(resultArr, tempGua)
					tempGua = GuaKV{}
					break
				}
			}
		}
	}

	return resultMap, resultArr
}

func saveGua64Map(kv GuaKV, gua64Map map[string]string) {
	gua64Map[kv.Key] = kv.Val
}

func (s *Suan) GetGua64Arr() []GuaKV {
	return s.gua64File2Arr
}

// 算
func (s *Suan) Do() {
	// 获取天地数
	tian, di := s.PeiShu.GetTianDiShu()
	// 获取先天卦
	gua := s.PeiShu.GetXianTianGua(tian, di)
	s.XianTianGua = s.transformGua(gua)

	// 计算元堂爻

	// 获取后天卦
	s.HouTianGua = "TODO"

	// 计算流年
}

func (s *Suan) transformGua(gua string) string {
	//guaNameMap := map[string]string{
	//	"乾 - 乾": "乾为天 （乾卦） 自强不息",     // 1
	//	"坤 - 坤": "坤为地 （坤卦） 厚德载物",     // 2
	//	"坎 - 震": "水雷屯 （屯卦） 起始维艰",     // 3
	//	"艮 - 坎": "山水蒙 （蒙卦） 启蒙奋发",     // 4
	//	"坎 - 乾": "水天需 （需卦） 守正待机",     // 5
	//	"乾 - 坎": "天水讼 （讼卦） 慎争戒讼",     // 6
	//	"坤 - 坎": "地水师 （师卦） 行险而顺",     // 7
	//	"坎 - 坤": "水地比 （比卦） 诚信团结",     // 8
	//	"巽 - 乾": "风天小畜 （小畜卦） 蓄养待进", // 9
	//	"乾 - 兑": "天泽履（履卦）脚踏实地",       // 10
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     // 11
	//	"乾 - 坤": "天地否（否卦）不交不通",     // 12
	//	"乾 - 离": "天火同人（同人卦）上下和同",   // 13
	//	"离 - 乾": "火天大有（大有卦）顺天依时",   // 14
	//	"坤 - 艮": "地山谦（谦卦）内高外低",     // 15
	//	"震 - 坤": "雷地豫（豫卦）顺时依势",     // 16
	//	"兑 - 震": "泽雷随（随卦）随时变通",     // 17
	//	"艮 - 巽": "山风蛊（蛊卦）振疲起衰",     // 18
	//	"坤 - 兑": "地泽临（临卦）教民保民",     // 19
	//	"巽 - 坤": "风地观（观卦）观下瞻上",     // 20
	//	"离 - 震": "火雷噬嗑（噬嗑卦）刚柔相济",   // 21
	//	"艮 - 离": "山火贲（贲卦）饰外扬质",     // 22
	//	"艮 - 坤": "山地剥（剥卦）顺势而止",     //23
	//	"坤 - 震": "地雷复（复卦）寓动于顺",     //24
	//	"乾 - 震": "天雷无妄（无妄卦）无妄而得",   // 25
	//	"艮 - 乾": "山天大畜（大畜卦）止而不止",   // 26
	//	"艮 - 震": "山雷颐（颐卦）纯正以养",     // 27
	//	"兑 - 巽": "泽风大过（大过卦）非常行动",   // 28
	//	"坎 - 坎": "坎为水（坎卦）行险用险",     // 29
	//	"离 - 离": "离为火（离卦）附和依托",     // 30
	//	"坤 - 乾": "泽山咸（咸卦）相互感应",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //31
	//	"坤 - 乾": "地天泰（泰卦）应时而变",     //3

	return s.gua64File2Map[gua]
}
