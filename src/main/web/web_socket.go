package web

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/*
WebSocket是HTML5的重要特性，它实现了基于浏览器的远程socket，它使浏览器和服务器可以进行全双工通信，
在websocket出现之前，为了实现即时通信，采用的技术都是"轮询"，
即在特定的时间间隔内，由浏览器对服务器发出http request，服务器在接收到请求后，返回最新的数据给浏览器刷新，
“轮询”使得浏览器需要对服务器不断发出请求，这样会占用大量带宽。
websocket采用了特殊的报头，使得浏览器和服务器只需要做一个握手的动作，就可以在浏览器和服务器之间建立一条连接通道。
且此连接会保持在活动状态，你可以使用javascript来向连接中写入或接收数据，
就像在使用一个常规的TCP Socket一样。
它解决了web实时化的问题，相比较传统http有如下好处：
·一个web客户端只建立一个tcp连接
·websocket服务端可以推送数据到web客户端
·有更轻量级的头，减少数据传送量
websocket url的起始输入是ws://或是wss://(在SSL上)
*/
/*
Web Socket原理
websocket的协议颇为简单，在第一次handshake通过后，连接便建立成功，其后的通讯数据都是以"\x00"开头，以"\xFF"结尾。
在客户端，这是透明的，websocket组件会自动将原始数据“掐头去尾”。
浏览器发出websocket连接请求，然后服务器发出回应，然后建立连接成功，这个过程通常成为握手
在请求中的"Sec-WebSocket-Key"是随机的，这是一个经过base64编码后的数据。
服务端收到这个请求之后需要把这个字符串连接上一个固定的字符串：
258EAFA5-E914-47DA-95CA-C5AB0DC85B11
即f7cb4ezEAl6C3wRaU6JORA==(示例请求中参数，此处并未列出)连接上上述那一串字符串，生成这样一个字符串：
f7cb4ezEAl6C3wRaU6JORA==258EAFA5-E914-47DA-95CA-C5AB0DC85B11
对该字符串先用sha1安全散列算法计算出二进制的值，然后用base64对其进行编码，既可以得到握手后的字符串：
rE91AJhfC+6JdVcVXOGJEADEJdQ=  将之作为响应头Sec-WebSocket-Key-Accept的值反馈给客户端。
*/
/*
Go 实现WebSocket
Go语言标准包里没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持，可以通过以下命令获取：
go get golang.org/x/net/websocket
但是此文中将以 gorilla/websocket 包来实现，获取方式如下：
go get github.com/gorilla/websocket
测试页面以postman代替
*/

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

/*
WebSocketHandle
Conn 类型表示WebSocket连接。服务器应用程序从HTTP请求处理程序调用
Upgrader.Upgrade方法以获取*Conn
*/
func WebSocketHandle(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// use conn to send and receive messages
	// 调用连接的WriteMessage和ReadMessage方法以字节切片的形式发送和接收消息
	for {
		// 读取客户端消息 message是一个[]byte, messageType是一个int类型websocket.BinaryMessage或websocket.TextMessage的值
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(message))
		// 调用WriteMessage给客户端写回消息
		if err = conn.WriteMessage(messageType, message); err != nil {
			log.Println(err)
			return
		}
	}
}
