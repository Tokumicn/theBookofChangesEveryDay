package service

import (
	"fmt"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/models"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/pkg/utils"
	"strings"
)

func BuildDictByStr(dictStr string) {

}

// 传入商品信息字符串 构建数据和日志
func BuildStuffByStr(stuffArr []string) error {

	stuffs := make([]models.Stuff, len(stuffArr))
	for _, stuf := range stuffArr {

		splits := strings.Split(stuf, ",")
		vMH, err := utils.ConvertStr2Float32(splits[2])
		if err != nil {
			// TODO log
			fmt.Printf("BuildStuffByStr-ConvertStr2Float32 str: %s, err: %v \n",
				splits[2], err)
			return err
		}

		vRM, err := utils.ConvertStr2Float32(splits[3])
		if err != nil {
			// TODO log
			fmt.Printf("BuildStuffByStr-ConvertStr2Float32 str: %s, err: %v \n",
				splits[3], err)
			return err
		}

		order, err := utils.ConvertStr2Int(splits[4])
		if err != nil {
			// TODO log
			fmt.Printf("BuildStuffByStr-ConvertStr2Float32 str: %s, err: %v \n",
				splits[3], err)
			return err
		}

		region, err := utils.ConvertStr2Int(splits[5])
		if err != nil {
			// TODO log
			fmt.Printf("BuildStuffByStr-ConvertStr2Float32 str: %s, err: %v \n",
				splits[3], err)
			return err
		}

		temp := models.Stuff{
			QName:    splits[0],
			Name:     splits[1],
			ValMH:    vMH,
			ValRM:    vRM,
			Order:    order,
			RegionID: region,
		}

		temp, err = buildVal(temp)
		if err != nil {
			// TODO log
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
