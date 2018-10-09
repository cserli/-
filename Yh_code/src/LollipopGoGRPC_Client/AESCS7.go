package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	//	"crypto/rand"
	//	"encoding/base64"
	//	"fmt"
	//	"io"
)

func Encrypt(plantText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key) //选择加密算法
	if err != nil {
		return nil, err
	}
	plantText = PKCS7Padding(plantText, block.BlockSize())

	blockModel := cipher.NewCBCEncrypter(block, key)

	ciphertext := make([]byte, len(plantText))

	blockModel.CryptBlocks(ciphertext, plantText)
	return ciphertext, nil
}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

/*CBC加密 按照golang标准库的例子代码
不过里面没有填充的部分,所以补上
*/

////使用PKCS7进行填充，IOS也是7
//func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
//	padding := blockSize - len(ciphertext)%blockSize
//	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
//	fmt.Println("ciphertext:-----------", string(ciphertext))
//	fmt.Println("padtext:-----------", string(padtext))
//	//return append(ciphertext, padtext...)
//	return ciphertext
//}

//func PKCS7UnPadding(origData []byte) []byte {
//	length := len(origData)
//	unpadding := int(origData[length-1])
//	return origData[:(length - unpadding)]
//}

////aes加密，填充秘钥key的16位，24,32分别对应AES-128, AES-192, or AES-256.
//func AesCBCEncrypt(rawData, key []byte) ([]byte, error) {
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}
//	iv := []byte(ivParameter)
//	//填充原文
//	blockSize := block.BlockSize()
//	rawData = PKCS7Padding(rawData, blockSize)
//	sddddd := rawData
//	//初始向量IV必须是唯一，但不需要保密
//	cipherText := make([]byte, blockSize+len(rawData))
//	// cipherText := make([]byte, len(rawData))
//	//block大小 16
//	//iv := cipherText[:blockSize]

//	//if _, err := io.ReadFull(rand.Reader, iv); err != nil {
//	//	panic(err)
//	//}

//	fmt.Println("iv:-----------", string(iv))

//	//block大小和初始向量大小一定要一致
//	mode := cipher.NewCBCEncrypter(block, iv)
//	fmt.Println("blockSize:-----------", blockSize)
//	mode.CryptBlocks(cipherText[blockSize:], rawData)
//	fmt.Println("rawData:-----------", string(rawData))
//	fmt.Println("cipherText:-----------", string(cipherText))
//	fmt.Println("sddddd:-----------", string(sddddd))
//	return cipherText, nil
//	//return rawData, nil
//}

//func AesCBCDncrypt(encryptData, key []byte) ([]byte, error) {
//	//func AesCBCDncrypt(encryptData []byte) ([]byte, error) {

//	//key := []byte(sKey)
//	iv := []byte(ivParameter)
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		panic(err)
//	}

//	blockSize := block.BlockSize()

//	if len(encryptData) < blockSize {
//		panic("ciphertext too short")
//	}
//	//iv := encryptData[:blockSize]
//	encryptData = encryptData[blockSize:]

//	// CBC mode always works in whole blocks.
//	if len(encryptData)%blockSize != 0 {
//		panic("ciphertext is not a multiple of the block size")
//	}

//	mode := cipher.NewCBCDecrypter(block, iv)

//	// CryptBlocks can work in-place if the two arguments are the same.
//	mode.CryptBlocks(encryptData, encryptData)
//	//解填充
//	encryptData = PKCS7UnPadding(encryptData)
//	return encryptData, nil
//}

//// []byte(sKey)
//func Encrypt(rawData, key []byte) (string, error) {
//	data, err := AesCBCEncrypt(rawData, key)
//	if err != nil {
//		return "", err
//	}
//	return base64.StdEncoding.EncodeToString(data), nil
//}

//func Dncrypt(rawData string, key []byte) (string, error) {
//	data, err := base64.StdEncoding.DecodeString(rawData)
//	if err != nil {
//		return "", err
//	}
//	dnData, err := AesCBCDncrypt(data, key)
//	if err != nil {
//		return "", err
//	}
//	return string(dnData), nil
//}
