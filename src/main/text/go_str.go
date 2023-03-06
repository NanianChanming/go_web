package text

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

/*
StrHandler
字符串操作
go标准库strings和strconv两个包
*/
func StrHandler(w http.ResponseWriter, r *http.Request) {
	// Contains 字符串中是否包含substr, 返回bool值
	/*fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))*/

	// func Join(a []string, sep string) string
	// Join 字符串连接，把slice a通过sep连接起来
	s := []string{"关羽", "张飞", "赵云", "黄忠", "马超"}
	fmt.Println(strings.Join(s, "|"))

	// func Index(s, sep string) int
	// 在字符串s中查找sep所在位置，返回位置值，找不到返回-1
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	// func Repeat(s string, count int) string
	// 重复s字符串count次，最后返回重复的字符串
	fmt.Println(strings.Repeat("no", 3))

	// func Replace(s, old, new string, n int) string
	// 在s字符串中，把old字符串替换为new字符串，n表示替换的次数，小于0表示全部替换
	fmt.Println(strings.Replace("oink oink oink", "in", "", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))

	// func Split(s, sep string) []string
	// 把s字符串按照sep分割，返回slice
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo 0'Higgins"))

	// func Trim(s string, cutset string) string
	// 在s字符串的头部和尾部去除cutset指定的字符串
	fmt.Printf("[%q]\n", strings.Trim(" !!! Achtung !!! ", "! "))

	// func Fields(s string) []string
	// 去除s字符串的空格符，并且按照空格分割返回slice
	fmt.Printf("Fields are %q\n", strings.Fields("  foo bar  baz"))

	// 字符串转换
	// Append 系列函数将整数等转换为字符串后，添加到现有的字节数组中
	str := make([]byte, 0, 100)
	str = strconv.AppendInt(str, 4567, 10)
	str = strconv.AppendBool(str, false)
	str = strconv.AppendQuote(str, "abcdefg")
	str = strconv.AppendQuoteRune(str, '单')
	fmt.Println(string(str))

	// Format系列函数把其他类型的转换为字符串
	a := strconv.FormatBool(false)
	b := strconv.FormatFloat(123.32, 'g', 12, 64)
	c := strconv.FormatInt(1234, 10)
	d := strconv.FormatUint(12345, 10)
	e := strconv.Itoa(1023)
	fmt.Println(a, b, c, d, e)

	// Parse 系列函数把字符串转换为其他类型
	a1, err := strconv.ParseBool("false")
	checkError(err)
	b1, err := strconv.ParseFloat("123.23", 64)
	checkError(err)
	c1, err := strconv.ParseInt("1234", 10, 64)
	checkError(err)
	d1, err := strconv.ParseUint("12345", 10, 64)
	checkError(err)
	e1, err := strconv.Atoi("1023")
	checkError(err)
	fmt.Println(a1, b1, c1, d1, e1)
}

func checkError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
