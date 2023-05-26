package ganzhi

import "testing"

func TestPeiShu(t *testing.T) {

	tests := []struct {
		total  int
		seed   int
		answer int
	}{
		// 天数计算
		{17, 25, 7},
		{20, 25, 2},
		{23, 25, 3},
		{25, 25, 5},
		{29, 25, 4},
		{38, 25, 3},

		// 地数计算
		{15, 30, 5},
		{26, 30, 6},
		{30, 30, 3},
		{38, 30, 8},
		{45, 30, 5},
	}

	for _, test := range tests {
		tempAn := calc(test.total, test.seed)
		if tempAn != test.answer {
			t.Errorf("测试用例: {total: %d, seed: %d, 期望结果(%d) != 错误结果(%d)}", test.total, test.seed, test.answer, tempAn)
		}
	}

}
