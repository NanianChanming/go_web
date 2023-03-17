package security

/*
现在的Web应用都含有大量的动态内容以提高用户体验。
所谓动态内容，就是应用程序能够根据用户环境和用户请求，输出相应的内容。
动态站点会受到一种名为“跨站脚本攻击”（Cross Site Scripting，安全专家通常缩写为XSS）的威胁，
而静态站点则完全不受其影响。
*/

/**
如何预防XSS
很简单，坚决不要相信用户的任何输入，并过滤掉输入中的所有特殊字符，这样就可以消灭绝大部分的XSS攻击
目前防御XSS主要有如下几种方式：
·过滤特殊字符
避免XSS的方法之一主要就是将用户所提供的内容进行过滤，Go语言提供了HTML的过滤函数：
text/template包下面的HTMLEscapeString、JSEscapeString等函数
·使用HTTP指定头类型
`w.Header().Set("Content-Type", "text/javascript")`
这样就可以让浏览器解析javascript代码，而不是html输出
*/
