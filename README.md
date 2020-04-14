
go-learn 学习笔记

## compare 包

在Go中比较两个Slice的大小，可以使用`reflect.DeepEqual`，但此方法性能较低。
因此自定义包方法用于比较Slice，在已知slice的具体值类型时使用该包性能得到大大提高。



## 使用

- 拉取项目代码
```
go get -u -v github.com/gushasha/go-learn
```
- 使用示例
```
import (
    "fmt"
    "github.com/gushasha/go-learn/compare"
)
func test() {
    s1 := []string{"h", "e", "l", "l", "o"}
    s2 := []string{"h", "e", "l", "l", "o"}
    if compare.Strings(s1, s2) {
        fmt.Println("equal..")
    } else {
        fmt.Println("not equal!")
    }
}
```

## 性能分析

以`compare.Strings`为例，对比 `reflect.DeepEqual` Benchmark测试结果如下：

```
BenchmarkStrings-4            	 2000000	       702 ns/op
BenchmarkRelfectDeepEqual-4   	 1000000	      1316 ns/op
```
