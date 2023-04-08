package deploy

/*
seelog是用Go语言实现的一个日志系统，它提供了一些简单的函数来实现复杂的日志分配、过滤和格式化，
主要有如下特性
·XML的动态配置，可以不用重新编译程序而动态的加载配置信息
·支持热更新，能够动态改变配置而不需要重启应用
·支持多输出流，能够同时把日志输出到多种流中、例如文件流、网络流等
·支持不同的日志输出
	命令行输出、文件输出、缓存输出、支持log rotate、SMTP邮件
上面只列举了部分特性，seelog是一个特别强大的日志处理系统，详细内容参考官方wiki，
下面介绍如何使用：
首先安装seelog
go get -u github.com/cihub/seelog
*/

/*var Logger log.LoggerInterface

func init() {
	seelogDemo()
	DisableLog()
	loadLogConfig()
	http.HandleFunc("/seelog", SeeLogDemo)
}

func seelogDemo() {
	defer log.Flush()
	log.Info("Hello from Seelog")
}

func loadLogConfig() {
	fileInfo, err := os.Stat("./file/seelog_config.xml")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	bytes := make([]byte, fileInfo.Size())
	file, err := os.Open("./file/seelog_config.xml")
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	file.Read(bytes)
	logger, err := log.LoggerFromConfigAsBytes(bytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}
	UseLogger(logger)
}

func DisableLog() {
	Logger = log.Disabled
}

func UseLogger(newLogger log.LoggerInterface) {
	Logger = newLogger
}*/

/*
上面主要实现了三个函数
DisableLog
初始化全局变量Logger为seelog的禁用状态，主要为了防止Logger被多次初始化
loadAppConfig
根据配置文件初始化seelog的配置信息，这里读取xml文件。里面的配置说明如下
·seelog
	minlevel参数可选，如果被配置，高于或等于此级别的日志会被记录，maxlevel同理
·outputs
	输出信息的目的地，这里分成了两份数据，一份记录到logrotate文件里面，另一份设置了filter，如果这个错误级别是critical，那么将发送报警邮件
·formats
	定义了各种日志的格式
UseLogger
设置当前的日志器为相应的日志处理
*/

/*
SeeLogDemo
定义日志处理包之后，下面是使用示例
*/
/*func SeeLogDemo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		log.Flush()
	}()
	Logger.Info("start log")
}*/
