package text

import (
	"fmt"
	"net/http"
	"os"
)

/**
在计算机设备中，文件都是必须的对象，
go中的文件操作大多函数都在os包里面，下面列举几个目录操作：
·func Mkdir(name string, perm FileMode) error
创建名称为name的目录，权限设置是perm, 例如0777
·func MkdirAll(path string, perm FileMode) error
根据path创建多级子目录，例如test/test1/test2
·func Remove(name string) error
删除名称为name的目录，当目录下有文件或者其他目录时会出错
·func RemoveAll(path string) error
根据path删除多级子目录，如果path是单个名称，那么该目录下的子目录全部删除
*/

func FileHandler(w http.ResponseWriter, r *http.Request) {
	//DirDemo()
	//FileCreate()
	//FileOpen()
	//FileWrite()
	//FileRead()
	FileDelete()
}

/*
DirDemo
目录操作
*/
func DirDemo() {
	//os.Mkdir("dir-test", 0777)
	//os.MkdirAll("dir-test/test1/test2", 0777)
	err := os.Remove("dir-test")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("dir-test")
}

/*
FileCreate
文件操作
新建文件可以通过如下两个方法
·func Create(name string) (file *File, err Error)
根据提供的文件名创建新的文件，返回一个文件对象，默认权限是0666的文件，返回的文件对象是可读写的
·func newFile(fd uintptr, name string) *File
根据文件描述符创建相应的文件，返回一个文件对象
*/
func FileCreate() {
	os.Create("./file/filetest.txt")
	// 使用newfile创建的文件不会真正被保存
	os.NewFile(0, "./file/filetest-1.txt")
}

/*
FileOpen
可以通过下面两个方法打开文件
·func Open(name string)(file *File, err Error)
该方法打开一个名称为name的文件,但是是只读方式，内部实现其实调用了OpenFile
·func OpenFile(name string, flag int, perm unit32)(file *File, err Error)
打开名为name的文件，flag是打开的方式，只读、读写等，perm是权限
*/
func FileOpen() {
	file, err := os.Open("./file/filetest.txt")
	if err != nil {
		fmt.Println("打开文件失败, err: ", err)
		return
	}
	file.Close()

	file, err = os.OpenFile("./file/filetest.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("打开文件失败, err: ", err)
		return
	}
	file.Close()
}

/*
FileWrite
写入文件函数
·func (file *File) Write(b []byte) (n int, err Error)
写入 byte 类型的信息到文件
·func (file *File) WriteAt(b []byte, off int64) (n int, err Error)
在指定位置开始写入 byte 类型的信息
·func (file *File) WriteString(s string) (ret int, err Error)
写入 string 信息到文件
*/
func FileWrite() {
	file, err := os.OpenFile("./file/filetest.txt", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("准备写入-打开文件失败, err: ", err)
		return
	}
	defer file.Close()
	file.Write([]byte("这是一首简单的小情歌\r\n"))
	file.WriteString("唱着人们心头的白鸽\r\n")
}

/*
FileRead
·func (file *File) Read(b []byte) (n int, err Error)
读取数据到 b 中
·func (file *File) ReadAt(b []byte, off int64) (n int, err Error)
从 off 开始读取数据到 b 中
*/
func FileRead() {
	file, _ := os.Open("./file/filetest.txt")
	defer file.Close()
	bytes := make([]byte, 1024)
	for {
		n, _ := file.Read(bytes)
		if n == 0 {
			break
		}
		os.Stdout.Write(bytes[:n])
	}
}

/*
FileDelete
Go 语言里面删除文件和删除文件夹是同一个函数
·func Remove(name string) Error
调用该函数就可以删除文件名为 name 的文件
*/
func FileDelete() {
	os.Remove("./file/filetest.txt")
}
