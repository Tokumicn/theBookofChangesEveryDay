package service

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/database"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/pkg/common"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/pkg/utils"
	"os"
	"strings"
)

func BuildDictByStr(dictStr string) {

}

// 查询商品信息
func QueryStuff(inStr string) (int64, []database.Stuff, error) {
	var (
		ctx    = context.Background()
		qStuff database.Stuff
		err    error
		offset int
		limit  = 50
	)

	qStuff, err = buildQuery(inStr)
	if err != nil {
		return 0, nil, err
	}

	// TODO log
	fmt.Printf("QueryStuff-buildQuery qStuff: %+v", qStuff)

	return qStuff.List(ctx, offset, limit)
}

// 根据输入字符串构建更为精准的查询条件  数据量不大该过程可以通过map映射完成
func buildQuery(inStr string) (database.Stuff, error) {
	res := database.Stuff{}

	qNameStr, ok := common.QueryQNameMap[inStr]
	if ok {
		res.QName = qNameStr
	}

	nameStr, ok := common.QueryNameMap[inStr]
	if ok {
		res.Name = nameStr
	}

	if len(nameStr) == 0 && len(qNameStr) == 0 {
		fmt.Printf("[优化内容经常查看] buildQuery [inStr: %s] 出现无法映射的用户输入, 请查看后收录进查询映射表\n", inStr)
		res.Name = inStr // 都无法命中的尝试用name字段查询
	}

	return res, nil
}

// 传入商品信息字符串 构建数据和日志
func BuildStuffByStr(stuffArr []string) error {
	ctx := context.Background()

	stuffs := make([]database.Stuff, len(stuffArr))

	// 读取本地文件增加到本次处理商品信息中
	tempStuffs, err := readCSVFromStuffData()
	if err != nil {
		// TODO log
		fmt.Printf("BuildStuffByStr-readCSVFromStuffData err: %v\n", err)
		return err
	}
	stuffArr = append(stuffArr, tempStuffs...)

	for _, stuf := range stuffArr {

		temp, err := str2Stuff(stuf)
		if err != nil {
			// TODO log
			fmt.Printf("buildStuff [temp: %v] err: %v \n", temp, err)
			continue
		}

		temp, err = buildStuffVal(temp)
		if err != nil {
			// TODO log
			fmt.Printf("buildVal [temp: %v] err: %v \n", temp, err)
			continue
		}
		stuffs = append(stuffs, temp)
	}

	// 存储数据
	err = saveStuffs(ctx, stuffs)
	if err != nil {
		// TODO log
		return err
	}

	return nil
}

// 存储Stuff信息，根据Name判断是否已经存放，该段为全库表唯一
func saveStuffs(ctx context.Context, list []database.Stuff) error {

	for _, s := range list {
		isExist, id, err := s.ExistByQName(ctx)
		if err != nil {
			// TODO log
			return err
		}

		if isExist { // 更新
			s.ID = id
			_, err = s.Update(ctx)
			if err != nil {
				// TODO log
				return err
			}
		} else {
			_, err = s.Create(ctx)
			if err != nil {
				// TODO log
				return err
			}
		}
	}

	return nil
}

// 将字符串转换为对象
func str2Stuff(stuf string) (database.Stuff, error) {

	var (
		err    error
		qName  string
		name   string
		vMH    float32
		vRM    float32
		order  int
		region int
		empty  database.Stuff
	)

	splits := strings.Split(stuf, ",")

	// 必填字段  没有则报错
	qName, err = utils.ArrGetWithCheck(splits, 0)
	if err != nil {
		return empty, err
	}

	// 必填字段  没有则报错
	name, err = utils.ArrGetWithCheck(splits, 1)
	if err != nil {
		return empty, err
	}

	// 非必填字段
	vMHStr, _ := utils.ArrGetWithCheck(splits, 2)
	vMH, err = utils.ConvertStr2Float32(vMHStr)
	if err != nil {
		return empty, err
	}

	// 非必填字段
	vRMStr, _ := utils.ArrGetWithCheck(splits, 3)
	vRM, err = utils.ConvertStr2Float32(vRMStr)
	if err != nil {
		return empty, err
	}

	orderStr, _ := utils.ArrGetWithCheck(splits, 4)
	order, err = utils.ConvertStr2Int(orderStr)
	if err != nil {
		return empty, err
	}

	regionStr, _ := utils.ArrGetWithCheck(splits, 5)
	region, err = utils.ConvertStr2Int(regionStr)
	if err != nil {
		return empty, err
	}

	temp := database.Stuff{
		QName:    qName,
		Name:     name,
		ValMH:    vMH,
		ValRM:    vRM,
		Order:    order,
		RegionID: region,
	}

	return temp, nil
}

// 填充对象内容
func buildStuffVal(s database.Stuff) (database.Stuff, error) {

	if (s.ValMH == 0 && s.ValRM == 0) || (s.ValMH != 0 && s.ValRM != 0) {
		// log 无需转换
		return s, nil
	}

	if s.ValMH == 0 {
		valMH, err := utils.RM2MH(s.ValRM)
		if err != nil {
			// TODO log
			fmt.Printf("RM2MH[ValRM: %f] err: %v\n", s.ValRM, err)
			return s, err
		}
		s.ValMH = valMH
	}

	if s.ValRM == 0 {
		valRM, err := utils.MH2RM(s.ValMH)
		if err != nil {
			// TODO log
			fmt.Printf("MH2RM[ValMH: %f] err: %v\n", s.ValMH, err)
			return s, err
		}
		s.ValRM = valRM
	}

	return s, nil
}

// 从stuff数据文件中读取数据
func readCSVFromStuffData() ([]string, error) {
	res := make([]string, 0)
	f, err := os.Open("/Users/zhangrui/Workspace/goSpace/src/Tokumicn/theBookofChangesEveryDay/mhxyHelper/config/stuff_data.csv")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	r.ReadLine() // 丢弃第一行
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		res = append(res, string(line))
	}
	return res, nil
}
