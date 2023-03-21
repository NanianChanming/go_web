package test

import "testing"

/*
gotest_test.go 这是单元测试文件，有下面这些原则：
·文件名必须是以_test.go结尾的，这样在执行go test的时候才会执行到相应的代码
·你必须import testing这个包
·所有的测试用例函数必须是Test开头
·测试用例会按照源代码中写的顺序依次执行
·测试函数TestXxx()的参数是testing.T,我们可以使用该类型来记录错误或者是测试状态
·测试格式：func TestXxx(t *testing.T), Xxx部分可以为任意的字母数字组合，但是首字母不能是小写
·函数通过调用testing.T的Error，Errorf，FailNow，Fatal，FatalIf方法，说明测试不通过，调用Log方法用来记录测试信息
下面是测试用例的代码
*/
func Test_Division_1(t *testing.T) {
	if i, err := Division(6, 2); i != 3 || err != nil {
		t.Error("除法函数测试未通过")
	} else {
		t.Log("第一个测试通过")
	}
}
