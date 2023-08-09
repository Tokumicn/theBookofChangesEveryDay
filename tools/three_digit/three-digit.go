package three_digit

import (
	"github.com/Tokumicn/theBookofChangesEveryDay/common"
	"github.com/Tokumicn/theBookofChangesEveryDay/tools/utils"
)

// 将数转换为卦
func TransformGuaByThreeDigit(shang, xia int) string {
	shangGua := common.XianTianIndexMap[shang]
	xiaGua := common.XianTianIndexMap[xia]

	return utils.FormatGua(shangGua, xiaGua)
}
