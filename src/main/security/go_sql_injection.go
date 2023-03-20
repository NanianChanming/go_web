package security

/*
sql 注入攻击（SQL Injection）简称注入攻击，是web开发中最常见的一种安全漏洞。
可以用它来从数据库获取敏感信息，或者利用数据库的特性执行添加用户，导出文件等一些列恶意操作，
甚至有可能获取数据库乃至系统用户最高权限。
造成sql注入的原因是因为程序没有有效过滤用户的输入，使攻击者成功向服务器提交恶意的sql查询代码，
程序在接收后错误的将攻击者的输入作为了查询语句的一部分执行，导致原始的查询逻辑被改变
额外的执行了攻击者精心构造的恶意代码。
*/

/*
通过一些案例来解释sql注入的方式
*/
func sqlInjectionDemo(username, password string) {
	sql := "select * from user where username = " + "'" + username + "'and password = '" + password + ""
	// 假定用户输入的用户名如下，密码随意
	// username = myuser' or 'foo' = 'foo' --
	// 那么拼接后的sql就变成了如下所示
	sql = "select * from user where username = 'myuser' or 'foo' = 'foo' -- and password = 'xxx'"
	// --在sql里使注释标记，所以查询语句会在此终端，这就让攻击者在不知道任何合法用户名和密码的情况下成功登录
	// 对于mysql还有更危险的一种注入方式，就是控制系统
	var prod string
	sql = "select * from products where name like '%'" + prod + "'%'"
	Db.Exec(sql)
	// 如果攻击提交 a%' exec master..xp_cmdshell 'net user test testpass/ADD' ，那么整条sql将变成
	sql = "select * from products where name like '%a%' exec master..xp_cmdsheel 'net user test testpass /ADD' --%'"
	// mssql服务器会执行这条sql语句，包括它后面那个用于向系统添加新用户的命令。
	// 如果这个服务器是以sa运行而mssqlserver服务又有足够的权限的话，攻击者就可以获得一个系统账号来访问主机了。
}

/*
如何防止sql注入
永远不要相信外界输入的数据，特别是来自用户的数据，包括选择框、表单隐藏域和cookie。
下面一些建议或许对防治sql注入有一定帮助
·严格限制web应用的数据库操作权限，给此用户提供仅仅能够满足其工作的最低权限，从而最大限度减少注入攻击对数据库的危害。
·检查输入的数据是否有所期望的数据格式，严格限制变量的类型，例如使用regexp包进行一些匹配处理，或者使用strconv包对字符串转化成其他基本类型的数据进行判断。
·对进入数据库的特殊字符进行转义处理，或编码转换。
go的 text/template 包里面的HTMLEscapeString函数可以对字符串进行转义处理。
·所有的查询语句建议使用数据库提供的参数化查询接口，参数化的语句使用参数而不是将用户变量嵌入到sql中，即不要直接拼接sql语句。
·在应用发布之前建议使用专业的SQL注入检测工具进行检测，以及及时修补被发现的SQL注入漏洞。
·避免在网站打印出SQL错误信息，比如类型错误，字段不匹配等，把代码里的sql语句暴露出来，以防止攻击者利用这些错误信息进行注入攻击。
*/
