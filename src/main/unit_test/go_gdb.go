package unit_test

/*
Go内部已经内置了GDB，所以，我们可以通过GDB来进行调试
另外，纯go代码建议使用delve可以很好的进行Go代码调试
*/
/*
GDB调试简介
GDB是FSF（自由软件基金会）发布的一个强大的类UNIX系统下的程序调试工具，
使用GDB可以做如下事情：
·启动程序，可以按照开发者的自定义要求运行程序
·可让被调试的程序在开发者设定的调置的断点处停住，（断点可以是条件表达式）
·当程序被停住时，可以检查此时程序中所发生的事
·动态的改变当前程序的执行环境
目前支持调试的Go程序的GDB版本必须大于7.1。
编译Go程序的时候需要注意以下几点
·传递参数 -ldflags "-s",忽略debug的打印信息
·传递 -gcflags "-N -l",这样可以忽略Go内部做的一些优化，聚合变量和函数等优化，
这样对于GDB调试来说非常困难，所以在编译的时候加入这两个参数避免这些优化。
*/
/*
常用命令
·list
简写命令 l, 用来显示源代码，默认显示十行，后面可以带上参数的具体行，
例如：list 15, 显示十行代码，其中第15行在显示的十行里面的中间
·break
简写命令 b, 用来设置断电，后面跟上参数设置断点的行数，
例如：b 10在第十行设置断点
·delete
简写命令 d, 用来删除断点，后面跟上断点设置的序号，这个序号可以通过 info breakpos 获取相应的断点序号
·backtrace
简写命令 bt，用来打印执行的代码过程
·info
info命令用来显示信息，后面有几种参数，我们常用的有如下几种：
info locals 显示当前执行的程序中的变量值
info breakpoints 显示当前设置的断点列表
info goroutines 显示当前执行的goroutine列表，如下所示，带 * 的表示当前执行的
	* 1  running runtime.gosched
	* 2  syscall runtime.entersyscall
	  3  waiting runtime.gosched
	  4 runnable runtime.gosched
·print
简写命令p, 用来打印变量或者其他信息，后面跟上需要打印的变量名，
还有一些很有用的函数$len() 和 $cap(), 用来返回当前string、slices或者maps的长度和容量
·whatis
用来显示当前变量的名称，后面跟上变量名，
例如：whatis msg
显示如下：
type = struct string
·next
简写命令 n, 用来单步调试，跳到下一步，当有断点之后，可以输入n跳转到下一步继续执行
·continue
简写命令 c, 用来跳出当前断点处，后面可以跟参数N，跳过多少次断点
·set variable
该命令用来改变运行过程中的变量值，格式如： set variable <var>=<value>
*/
/*
demo 示例位于go_demo项目中
*/
