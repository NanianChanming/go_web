package test

import (
	"errors"
	"testing"
)

/*
如何编写压力测试
压力测试用来检测函数的性能，和编写单元功能测试的方法类似，需要注意以下几点
·压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
func BenchmarkXxx(b *testing.B)
·go test不会默认执行压力测试的函数, 如果执行压力测试需要带上参数 -bench,
语法：go test -bench="test_name_regex"
例如：go test -bench=".*" 表示测试全部的压力测试函数
·在压力测试用例中，请记得在循环体内使用testing.B.N以使测试可以正常运行
·文件名也必须以 _test.go结尾
*/
func Benchmark_Division(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division1(4, 5)
	}
}

func Benchmark_TimeConsumingFunction(b *testing.B) {
	b.StopTimer() //调用该函数停止压力测试的时间计数
	// 做一些初始化的工作，例如读取文件数据，数据库连接之类的工作
	// 这样这些时间不影响我们测试函数本身的性能
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division1(4, 5)
	}
}

func Division1(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
