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

		temp = buildVal(temp)

		stuffs = append(stuffs, temp)
	}

	fmt.Printf("stuffs: %+v \n", stuffs)

	return nil
}

func buildVal(s models.Stuff) (models.Stuff, error) {

	if s.ValMH == 0 && s.ValRM == 0 {
		return
	}
}
