package animal

// Animal 生肖
type Animal struct {
	yearOrder int64
}

var animalAlias = [...]string{
	"鼠", "牛", "虎", "兔", "龙", "蛇",
	"马", "羊", "猴", "鸡", "狗", "猪",
}

// NewAnimal 创建生肖对象
func NewAnimal(yearOrder int64) *Animal {
	if !checkOrder(yearOrder) {
		return nil
	}
	return &Animal{yearOrder: yearOrder}
}

// Alias 返回生肖名称(鼠牛虎...)
func (animal *Animal) Alias() string {
	return animalAlias[(animal.yearOrder-1)%12]
}

func checkOrder(order int64) bool {
	return 1 <= order && order <= 12
}
