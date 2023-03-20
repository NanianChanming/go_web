package security

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/*
有时候需要把一些敏感数据加密后存储起来，在将来某个时候，随需要将它们解密出来，
此时我们应该在选用对称加密算法来满足我们的需求。
*/
/*
base64加解密
如果web应用足够简单，数据的安全性没有那么高的要求，那么可以采用一种比较简单的加解密方法是base64，
这种方式实现起来比较简单，Go语言的base64包已经很好的支持了这个
*/
func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

/*
高级加解密
Go语言的crypto里面支持对称加密的高级加解密包有：
·crypto/aes包：AES（Advanced Encryption Standard）, 又称Rijindael加密法，是美国联邦政府采用的一种区块加密标准
·crypto/des包：DES(Data Encryption Standard),是一种对称加密标准，是目前使用最广泛的密钥系统，
特别是在保护金融数据的安全中，曾是美国联邦政府的加密标准，但现在已经被AES所替代。
因为这两种算法使用方法类似，所以我们仅用aes为例
*/
var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func aesEncode() {
	// 需要加密的字符串
	plaintext := []byte("hello world")
	// aes的加密字符串
	key_text := "astaxie12798akljzmknm.ahkjkljl;k"
	// 创建加密算法
	c, err := aes.NewCipher([]byte(key_text))
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		return
	}

	// 加密字符串
	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	fmt.Printf("%s => %x\n", plaintext, ciphertext)

	// 解密字符串
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	fmt.Printf("%x => %s\n", ciphertext, plaintextCopy)
}
