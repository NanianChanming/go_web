package text

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 解析file包下config.xml
/**
可以通过xml包的Unmarshal函数来达到目的
func Unmarshal(data []byte, v interface{}) error
data接收的是xml数据流，v是需要输出的结构，定义为interface，也就是可以把xml转换为任意的格式。
这里主要介绍struct的转换，因为struct和xml都有类似树结构的特征。

Unmarshal 解析的时候 XML 元素和字段怎么对应起来的呢？
这是有一个优先级读取流程的，首先会读取 struct tag，
如果没有，那么就会对应字段名。必须注意一点的是解析的时候 tag、字段名、XML 元素都是大小写敏感的，所以必须一一对应字段。

解析XML到struct的时候遵循如下的规则:
1.如果struct的一个字段是string或者[]byte类型且它的tag含有“,innerxml”,
Unmarshal将会将此字段所对应的元素内所有内嵌的原始xml累加到此字段上。
2.如果 struct 中有一个叫做 XMLName，且类型为 xml.Name 字段(根节点),
那么在解析的时候就会保存这个 element 的名字到该字段，如上面例子中的 servers。
3.如果某个struct字段的tag定义中含有XML结构的element的名称，
那么解析的时候就会把相应的element值赋值给该字段，如servername和serverip定义
4.如果某个struct字段的tag定义了中含有“.attr”,
那么解析的时候就会将该结构对应的element的与字段同名的属性的值赋值给该字段，如version定义
5.如果某个struct字段的tag定义型如“a>b>c”，则解析的时候，会将xml结构a下面的c元素的值赋给该字段
6.如果某个struct字段的tag定义了“—”，那么不会为该字段解析匹配任何xml数据
7.如果struct字段后面的tag定义了“,any”,如果他的子元素在不满足其他的规则的时候就会匹配到这个字段
8.如果某个xml元素包含一条或者多条注释，那么这些注释将被累加到第一个tag含有“.comments”的字段上，
这个字段的类型可能是[]byte或string，如果没有这样的字段存在，那么注释将会被抛弃
*/

// Recurlyservers 定义struct
type Recurlyservers struct {
	XMLName     xml.Name `xml:"servers"`
	Version     string   `xml:"version,attr"`
	Encoding    string   `xml:"encoding,attr"`
	Svs         []server `xml:"server"`
	Description string   `xml:",innerxml"`
}

type server struct {
	XMLName    xml.Name `xml:"server"`
	ServerName string   `xml:"serverName"`
	ServerIP   string   `xml:"serverIP"`
}

func ParseXML(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("./file/config.xml")
	if err != nil {
		fmt.Fprintf(w, "文件读取失败")
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Fprintf(w, "文件内容读取失败")
		return
	}
	v := Recurlyservers{}
	err = xml.Unmarshal(data, &v)
	if err != nil {
		fmt.Fprintf(w, "xml解析失败")
		return
	}
	fmt.Println(v)
}

/**
输出xml
如果需要生成xml文件，xml包中提供了Marshal和MarshalIndent两个函数，
这两个函数主要的区别是第二个函数会增加前缀和缩进，函数定义如下：
func Marshal(v interface{})([]byte, error)
func MarshalIndent(v interface{}, prefix, indent string)([]byte, error)
这两个函数第一个参数是用来生成xml的结构定义类型数据，都是返回生成的xml数据流
os.Stdout.Writer([]byte(xml.Header)) 这句代码的出现是因为xml.MarshalIndent或者xml.Marshal输出的信息都是不带XML头的，
为了生成正确的xml文件，我们使用了xml包预定义的Header变量。
可以发现，Marshal函数接收的参数v是interface{}类型的，即它可以接受任意类型的参数，
那么xml包则是根据以下规则来生成相应的XML文件的：
1.如果v是array或者slice，那么输出每一个元素，类似value
2.如果v是指针，那么Marshal会输出指针指向的内容，如果指针为空，则什么都不输出
3.如果v是interface，那么就处理interface所包含的数据
4.如果v是其他数据类型，就会输出这个数据类型所拥有的字段信息
生成xml文件中的element的名字按照如下优先级从struct中获取
1.如果 v 是 struct，XMLName 的 tag 中定义的名称
2.类型为 xml.Name 的名叫 XMLName 的字段的值
3.通过 struct 中字段的 tag 来获取
4.通过 struct 的字段名用来获取
5.marshal 的类型名称
*/

func GenerateXML(w http.ResponseWriter, r *http.Request) {
	v := &Recurlyservers{Version: "1.0", Encoding: "utf-8"}
	v.Svs = append(v.Svs, server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	v.Svs = append(v.Svs, server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	output, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Printf("xml generate error: %v,\n", err)
		return
	}
	os.Stdout.Write([]byte(xml.Header))
	os.Stdout.Write(output)
}
