package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

const DEFAULT_CHARSET = "UTF-8"

// EncryptPublicKey RSA公钥加密
func EncryptPublicKey(data string, publicKey string) (string, error) {
	// 1. 解析公钥
	pubKey, err := parsePublicKey(publicKey)
	if err != nil {
		return "", err
	}

	// 2. 加密数据
	encryptedData, err := rsa.EncryptPKCS1v15(nil, pubKey, []byte(data))
	if err != nil {
		return "", err
	}

	// 3. Base64编码
	encodedData := base64.StdEncoding.EncodeToString(encryptedData)
	return encodedData, nil
}

// 解析PEM格式的公钥
func parsePublicKey(publicKey string) (*rsa.PublicKey, error) {
	// 解码PEM块
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return nil, errors.New("failed to parse PEM block containing the public key")
	}

	// 解析公钥
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言为RSA公钥
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	return rsaPub, nil
}
