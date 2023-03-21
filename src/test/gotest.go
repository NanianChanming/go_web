package test

import "errors"

/*
Go语言中自带有一个轻量级测试框架testing和自带的go test命令来实现单元测试和性能测试
testing框架和其他语言中的测试框架类似，你可以基于这个框架写针对相应函数的测试用例，
也可以基于该框架写相应的压力测试用例，
另外建议安装gotests插件自动生成测试代码
*/
/*
如何编写测试用例
由于go test命令只能在一个相应的目录下执行所有文件，所以我们新建了一个项目目录test，
这样我们所有的代码和测试代码都在这个目录下。
然后我们在该目录下创建两个文件：gotest.go和gotest_test.go
1.gotest.go: 这个文件里创建了一个包，里面有一个函数实现了除法运算
*/

func Division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
