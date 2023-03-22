package deploy

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

/*
Go中提供了一个简易的log包，我们使用该包可以方便的实现日志记录的功能，
这些日志都是基于fmt包的打印再结合panic之类的函数来进行一般的打印、抛出错误处理。
Go目前标准包只是包含了简单的功能，如果我们想把我们的日志保存到文件，然后又能结合日志实现很多复杂的功能，
可以使用第三方的日志系统logrus和seelog,他们实现了很强大的日志功能，可以根据自己的项目选择
*/
/*
logrus介绍
logrus是使用Go语言实现的一个日志系统，与标准库log完全兼容且核心API很稳定，是Go语言目前最活跃的日志库
首先安装logrus
go get -u github.com/sirupsen/logrus
*/
func init() {
	// 日志格式化为json，而不是默认的ASCII
	log.SetFormatter(&log.JSONFormatter{})
	// 输出 stdout 而不是默认的stderr, 也可以是一个文件
	log.SetOutput(os.Stdout)
	// 只记录严重或以上的警告
	log.SetLevel(log.DebugLevel)

	http.HandleFunc("/log", LogDemo)
	http.HandleFunc("/logFormat", LogFormat)
	log.Info("--log包加载--")
}

func LogDemo(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{"animal": "walrus"}).Info("A walrus appears")
}

func LogFormat(w http.ResponseWriter, r *http.Request) {
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously")

	//log.WithFields(log.Fields{
	//	"omg":    true,
	//	"number": 100,
	//}).Fatal("The ice breaks")

	// 通过日志语句重用字段
	// logrus.Entry 返回自 WithFields()
	contextLogger := log.WithFields(log.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log.WithTime(time.Now()).Info("this is time log")
}
