package text

import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"io/ioutil"
	"net/http"
	"os"
)

/**
解析json
go的json包中有如下函数
func Unmarshal(data []byte, v interface{}) error
*/

// Server 解析到结构体
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

type Serverslice struct {
	Servers []Server `json:"servers"`
}

func ParseJson(w http.ResponseWriter, r *http.Request) {
	var s Serverslice
	file, err := os.Open("./file/config.json")
	if err != nil {
		fmt.Printf("文件读取错误，error : %v\n", err)
		return
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &s)
	fmt.Println(s)
}

/**
解析到interface{}
当不知道json数据结构的情况下，可以用interface{}来接收任意数据类型的对象。
json包中采用map[string]interface{}和[]interface{}结构来存储任意的json对象和数组
go类型和json类型对应关系如下
·bool代表json booleans
·float64代表json numbers
·string代表json strings
·nil代表json null
*/

func ParseUnknownJson(w http.ResponseWriter, r *http.Request) {
	// 假设有如下json数据
	bytes := []byte(`{"Name":"ZhangSan","Age":6,"Parents": ["ZhangFei", "GuanYu"]}`)
	// 现在在不知道结构的情况下，把它解析到interface{}里
	var f interface{}
	json.Unmarshal(bytes, &f)
	// 这时候f里存储了一个map类型，他们的key是string，值存储在空的interface{}里
	//f = map[string]interface{}{
	//	"Name":    "ZhangSan",
	//	"Age":     6,
	//	"Parents": []interface{}{"ZhangFei", "GuanYu"},
	//}
	// 可以通过断言的方式访问
	m := f.(map[string]interface{})
	// 断言之后可以通过下面的方式访问里面的数据
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", v)
		case int:
			fmt.Println(k, "is int", v)
		case float64:
			fmt.Println(k, "is float64", v)
		case []interface{}:
			fmt.Println(k, "is array", v)
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}

/**
以上为go官方json包实现，另外有bitly公司开源的一个叫做simplejson的包，
在处理未知结构体的json时相当方便，示例如下
*/
func SimpleJsonDemo(w http.ResponseWriter, r *http.Request) {
	newJson, _ := simplejson.NewJson([]byte(`{
		"test": {
				"array": [1, "2", 3],
				"int": 10,
				"float": 5.150,
				"bignum": 9223372036854775807,
				"string": "simplejson",
				"bool": true
			}
		}`))
	array, _ := newJson.Get("test").Get("array").Array()
	i, _ := newJson.Get("test").Get("int").Int()
	i2, _ := newJson.Get("test").Get("bignum").Int64()
	s, _ := newJson.Get("test").Get("string").String()
	fmt.Println(array)
	fmt.Println(i)
	fmt.Println(i2)
	fmt.Println(s)
}

/**
json生成
json包里通过Marshal函数来处理，函数定义如下
func Marshal (v interface{})([]byte, error)
*/

func GenerateJson(w http.ResponseWriter, r *http.Request) {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("generate json error : ", err)
		return
	}
	fmt.Println(string(b))
}

/**
可以看到上面输出的字段名首字母都是大写，如果想用小写的首字母，可以通过struct tag实现
*/
func GenerateJsonDemo(w http.ResponseWriter, r *http.Request) {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("generate json error : ", err)
		return
	}
	fmt.Println(string(b))
}

/**
针对json的输出，定义struct tag的时候需要注意的几点是
·字段的tag是‘—’，那么这个字段不会输出到json
·tag中带有自定义名称，那么这个自定义名称会出现在json的字段名中
·tag中如果带有'omitempty'选项，那么如果该字段值为空，就不会输出到json串中
·如果字段类型是bool，string，int，int64等，而tag中带有“,.string”选项，
那么这个字段在输出到json的时候会把该字段对应的值转换成json字符串
*/

type User struct {
	// id字段不会输出到json中
	ID int `json:"-"`
	// 会以字符串格式输出
	Age int `json:"age,string"`
	// 如果为空则不输出
	Name string `json:"name,omitempty"`
}

func GenerateJsonDemo2(w http.ResponseWriter, r *http.Request) {
	u := User{
		ID:   1,
		Age:  18,
		Name: "Zhangsan",
	}
	bytes, _ := json.Marshal(u)
	fmt.Println(string(bytes))
}

/**
Marshal函数只有在转换成功的时候才会返回数据，在转换过程中需要注意几点
·json对象只支持string作为key，所以要编码一个map，那么必须是map[string]T这种类型的数据
·Channel,complex和function是不能被编码成json的
·嵌套的数据是不能编码的，不然会死循环
·指针在编码的时候会输出指针指向的内容，而空指针会输出nil
*/
