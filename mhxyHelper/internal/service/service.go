package service

import (
	"bufio"
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/models"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/pkg/utils"
	"os"
	"strings"
)

func BuildDictByStr(dictStr string) {

}

// 传入商品信息字符串 构建数据和日志
func BuildStuffByStr(stuffArr []string) error {

	stuffs := make([]models.Stuff, len(stuffArr))

	// 读取本地文件增加到本次处理商品信息中
	tempStuffs, err := readCSVFromStuffData()
	if err != nil {
		// TODO log
		fmt.Printf("BuildStuffByStr-readCSVFromStuffData err: %v\n", err)
		return err
	}
	stuffArr = append(stuffArr, tempStuffs...)

	for _, stuf := range stuffArr {

		var (
			qName  string
			name   string
			vMH    float32
			vRM    float32
			order  int
			region int
		)

		splits := strings.Split(stuf, ",")

		// 必填字段  没有则报错
		qName, err = utils.ArrGetWithCheck(splits, 0)
		if err != nil {
			return err
		}

		// 必填字段  没有则报错
		name, err = utils.ArrGetWithCheck(splits, 1)
		if err != nil {
			return err
		}

		// 非必填字段
		vMHStr, _ := utils.ArrGetWithCheck(splits, 2)
		vMH, err = utils.ConvertStr2Float32(vMHStr)
		if err != nil {
			return err
		}

		// 非必填字段
		vRMStr, _ := utils.ArrGetWithCheck(splits, 3)
		vRM, err = utils.ConvertStr2Float32(vRMStr)
		if err != nil {
			return err
		}

		orderStr, _ := utils.ArrGetWithCheck(splits, 4)
		order, err = utils.ConvertStr2Int(orderStr)
		if err != nil {
			return err
		}

		regionStr, _ := utils.ArrGetWithCheck(splits, 5)
		region, err = utils.ConvertStr2Int(regionStr)
		if err != nil {
			return err
		}

		temp := models.Stuff{
			QName:    qName,
			Name:     name,
			ValMH:    vMH,
			ValRM:    vRM,
			Order:    order,
			RegionID: region,
		}

		temp, err = buildVal(temp)
		if err != nil {
			fmt.Printf("buildVal [temp: %v] err: %v \n", temp, err)
			continue
		}
		stuffs = append(stuffs, temp)
	}

	fmt.Printf("stuffs: %+v \n", stuffs)

	return nil
}

func buildVal(s models.Stuff) (models.Stuff, error) {

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
