package util

import(
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

/// 去除填充位
func PKCSUnPadding(orig []byte) []byte {
	length := len(orig)
	unpadding := int(orig[length-1])
	if unpadding > length {
		return []byte{}
	}
	return orig[:length-unpadding]
}

/// 填充补码
func PKCSPadding(orig []byte, blockSize int) []byte {
	/// 计算需要补多少位数
	padding := blockSize - len(orig)%blockSize
	paddingtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(orig, paddingtext...)
}

// 3DES解密，key的长度必须是24byte
func TriDESEncrypt(plain, key []byte) ([]byte, error) {
	/// 获取block块
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	plain = PKCSPadding(plain, block.BlockSize())
	// 设置加密方式为3DES，使用3条56位的密钥对数据进行三次加密
	blockMode := cipher.NewCBCEncrypter(block, key[:8])
	// 创建明文长度的数组
	crypted := make([]byte, len(plain))
	blockMode.CryptBlocks(crypted, plain)
	return crypted, nil
}

// 3DES解密,key的长度必须是24byte
func TriDESDecrypt(crypted, key []byte) (content []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("decrypt error %s", e))
		}
	}()
	// 获取block块

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	// 创建切片
	context := make([]byte, len(crypted))
	// 设置解密方式
	blockMode := cipher.NewCBCDecrypter(block, key[:8])
	// 解密密文到数组
	blockMode.CryptBlocks(context, crypted)
	// 去补码
	context = PKCSUnPadding(context)

	return context, nil
}

func Sha1(data []byte) string {
	s := sha1.New()
	s.Reset()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}

func Sha256(data []byte) string {
	s := sha256.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}

func LoadRsaPulicbKeyByData(data []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("fail to parse PEM block containing the public key")
	}
	pub, e := x509.ParsePKIXPublicKey(block.Bytes)
	if e != nil {
		return nil, errors.New("fail to parse public key from PEM data")
	}
	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("unexpected public key type")
	}
	return rsaPub, nil
}

func LoadRsaPrivateKeyByData(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("fail to parse PEM block containing the private key")
	}
	priv, e := x509.ParsePKCS1PrivateKey(block.Bytes)
	if e != nil {
		return nil, e
	}
	return priv, nil
}

/// 公钥加密
func RsaEncryptFromKeyData(data []byte, key []byte) ([]byte, error) {
	pub, err := LoadRsaPulicbKeyByData(key)
	if err != nil {
		return nil, err
	}
	return rsa.EncryptPKCS1v15(rand.Reader, pub, data)
}

/// 私钥解密
func RsaDecryptFromKeyData(data []byte, key []byte) ([]byte, error) {
	priv, err := LoadRsaPrivateKeyByData(key)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, data)
}

func LoadRsaPulicbKey(filename string) (*rsa.PublicKey, error) {
	data, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	return LoadRsaPulicbKeyByData(data)
}

func LoadRsaPrivateKey(filename string) (*rsa.PrivateKey, error) {
	data, e := ioutil.ReadFile(filename)
	if e != nil {
		return nil, e
	}
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("fail to parse PEM block containing the private key")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

/// 将打乱的Key恢复正常
func LoadMixedKey(keyMixed string) ([]byte, error) {
	key, err := base64.StdEncoding.DecodeString(keyMixed)
	if err != nil {
		return nil, err
	}
	for index := 0; index+5 < len(key); index += 7 {
		key[index+3], key[index+5] = key[index+5], key[index+3]
	}
	for index, _ := range key {
		key[index] = byte(key[index] - 100)
	}
	key = bytes.Replace(key, []byte("RSA_PVB"), []byte("-----BEGIN RSA PRIVATE KEY-----"), -1)
	key = bytes.Replace(key, []byte("RSA_PVE"), []byte("-----END RSA PRIVATE KEY-----"), -1)
	key = bytes.Replace(key, []byte("RSA_PUB"), []byte("-----BEGIN PUBLIC KEY-----"), -1)
	key = bytes.Replace(key, []byte("RSA_PUE"), []byte("-----END PUBLIC KEY-----"), -1)

	return key, nil
}

/// 打乱Private Key
func MixKey(key []byte) string {
	key = bytes.Replace(key, []byte("-----BEGIN RSA PRIVATE KEY-----"), []byte("RSA_PVB"), -1)
	key = bytes.Replace(key, []byte("-----END RSA PRIVATE KEY-----"), []byte("RSA_PVE"), -1)
	key = bytes.Replace(key, []byte("-----BEGIN PUBLIC KEY-----"), []byte("RSA_PUB"), -1)
	key = bytes.Replace(key, []byte("-----END PUBLIC KEY-----"), []byte("RSA_PUE"), -1)

	/// 混肴处理
	for index, _ := range key {
		key[index] = byte(key[index] + 100)
	}

	for index := 0; index+5 < len(key); index += 7 {
		key[index+3], key[index+5] = key[index+5], key[index+3]
	}

	return base64.StdEncoding.EncodeToString(key)
}

func AESEncryptCFB(plain []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	cipherText := make([]byte, aes.BlockSize+len(plain))
	iv := cipherText[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(cipherText[aes.BlockSize:], plain)
	return cipherText, nil
}

func AESDecryptCFB(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(encrypted) < aes.BlockSize {
		return nil, errors.New("cipher text too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(encrypted, encrypted)
	return encrypted, nil
}