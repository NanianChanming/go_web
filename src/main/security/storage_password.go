package security

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

/*
很多网站遭遇过用户密码泄露事件，因为人们往往习惯在不同网站使用相同的密码，所以一旦泄露则全部泄露
作为web应用开发者，在选择密码存储方案时，该如何避免。
*/

/*
加密存储普通方案
目前用的最多的密码存储方案是将明文密码做单向哈希后存储，
单向哈希算法有一个特征:无法通过哈希后的摘要（digest）恢复原始数据，这也是“单向”二字的来源。
常用的单向哈希算法包括SHA-256，SHA-1，MD5等。
单向哈希有两个特性：
·同一个密码进行单向哈希，得到的总是唯一确定的摘要
·计算速度快。随着技术进步，一秒钟能够完成数十亿次单向哈希计算。
针对上面两个特点，考虑到多数人所使用的密码为常见的组合，攻击者可以将所有密码的常见组合进行单向哈希，得到一个摘要组合，
然后与数据库中的摘要进行比对即可获得对应的密码。
这个摘要组合也被称为rainbow table
因此通过单向加密之后存储的数据，和明文存储没有多大区别。
Go语言对这三种加密算法实现如下
*/
func encryptStr() {
	hash := sha256.New()
	io.WriteString(hash, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("%x", hash.Sum(nil))

	hash = sha1.New()
	io.WriteString(hash, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("%x", hash.Sum(nil))

	hash = md5.New()
	io.WriteString(hash, "需要加密的密码")
	fmt.Printf("%x", hash.Sum(nil))
}

/*
进阶方案
通过上面介绍我们知道可以用rainbow table来破解hash密码，很大程度上是因为加密使用的哈希算法是公开的，
如果不知道加密的hash算法是什么，那么就无从下手了
一个直接的解决办法是，自己设计一个哈希算法，然而，一个好的哈希算法是很难设计的，既要避免碰撞，又不能有明显的规律，
做到这两点要比想象中的要困难很多，因此实际应用中更多的是利用已有的哈希算法进行多次哈希。
单纯的多次哈希，依然阻挡不住黑客，两次MD5、三次MD5之类的方法，我们能想到，黑客自然也能想到。
特别是对于一些开源代码，这样哈希更是相当于直接把算法告诉了黑客。
现在安全性比较好的网站，都会用一种叫做加盐的方式来存储密码，也就是说salt，他们通常的做法是：
先将用户输入的密码进行一次MD5（或其他哈希算法进行加密），将得到的MD5值前后加上一些只有管理员自己知道的随机串，
再进行一次MD5加密，这个随机串中可以包括某些固定的串，也可以包括用户名（用来保证每个用户加密使用的密钥都不一样）
实现如下
*/
func encryptStr1() {
	// 假设用户名 abc, 密码 123456
	hash := md5.New()
	io.WriteString(hash, "需要加密的密码")
	pwmd5 := fmt.Sprintf("%x", hash.Sum(nil))
	// 指定两个salt
	salt1 := "@#$%"
	salt2 := "^&*()"
	// salt1 + 用户名 + salt2 + MD5拼接
	io.WriteString(hash, salt1)
	io.WriteString(hash, "abc")
	io.WriteString(hash, salt2)
	io.WriteString(hash, pwmd5)

	fmt.Sprintf("%x", hash.Sum(nil))
}

/*
进阶方案在几年前也许是足够安全的方案，因为攻击者没有足够的资源建立这么多的rainbow table。
但是时至今日，因为并行计算能力的提升，这种攻击已经完全可行。
如何解决这个问题？
只要时间和资源允许，没有破译不了的密码，所以方案是：故意增加密码计算所需耗费的资源和时间，
使得任何人都不可获得足够的资源建立所需的rainbow table。

这类方案有一个特点，算法中都有个因子，用于指明计算密码摘要所需要的资源和时间，也就是计算强度。
计算强度越大，攻击者建立rainbow table越困难，以至于不可继续。
*/
