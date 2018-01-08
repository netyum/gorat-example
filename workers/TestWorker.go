package workers

import "fmt"


//define Worker Args
//定义Worker的参数
type TestWorker struct {
	Str string
	Int int
	Float float64
	Bool bool
	Lint int64
	Map map[string]string
	TdMap map[string]map[string]string
	Slice []int
	TdSlice []interface{}
}

//实现方法
//feature method
func (w *TestWorker)Run() {
	fmt.Println("In TestWorker")
	fmt.Println("  String :", w.Str)
	fmt.Println("  Int :", w.Int)
	fmt.Println("  Float:", w.Float)
	fmt.Println("  Bool :", w.Bool)
	fmt.Println("  Lint int64 :", w.Lint)
	fmt.Println("  Map :", w.Map)
	fmt.Println("  Td Map :", w.TdMap)
	fmt.Println("  Slice :", w.Slice)
	fmt.Println("  Td Slice :", w.TdSlice)
}

