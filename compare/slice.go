package compare

type Interface interface {
	// len方法返回集合中的元素个数
	Len() int
	// item方法根据索引返回集合的值
	Item(idx int) interface{}
}

type IntSlice []int

func (p IntSlice) Len() int                 { return len(p) }
func (p IntSlice) Item(idx int) interface{} { return p[idx] }
func Ints(a, b []int) bool {
	return Compare(IntSlice(a), IntSlice(b))
}

type Int64Slice []int64

func (p Int64Slice) Len() int                 { return len(p) }
func (p Int64Slice) Item(idx int) interface{} { return p[idx] }
func Int64s(a, b []int64) bool {
	return Compare(Int64Slice(a), Int64Slice(b))
}

type Int32Slice []int32

func (p Int32Slice) Len() int                 { return len(p) }
func (p Int32Slice) Item(idx int) interface{} { return p[idx] }
func Int32s(a, b []int32) bool {
	return Compare(Int32Slice(a), Int32Slice(b))
}

type Float64Slice []float64

func (p Float64Slice) Len() int                 { return len(p) }
func (p Float64Slice) Item(idx int) interface{} { return p[idx] }
func Float64s(a, b []float64) bool {
	return Compare(Float64Slice(a), Float64Slice(b))
}

type Float32Slice []float32

func (p Float32Slice) Len() int                 { return len(p) }
func (p Float32Slice) Item(idx int) interface{} { return p[idx] }
func Float32s(a, b []float32) bool {
	return Compare(Float32Slice(a), Float32Slice(b))
}

type StringSlice []string

func (p StringSlice) Len() int                 { return len(p) }
func (p StringSlice) Item(idx int) interface{} { return p[idx] }
func Strings(a, b []string) bool {
	return Compare(StringSlice(a), StringSlice(b))
}

// Compare 循环遍历比较两个集合
// 先比较两个集合的长度是否相等
// 再循环遍历每一个元素进行比较
func Compare(a, b Interface) bool {
	if a.Len() != b.Len() {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for i := 0; i < a.Len(); i++ {
		if a.Item(i) != b.Item(i) {
			return false
		}
	}
	return true
}
