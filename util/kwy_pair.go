package util

import (
	"encoding/base64"
	"fmt"
	"time"
)

type KeyPairEncrypt struct {
	Keys   []string
	useAES bool
}

func NewKeyPairEncrypt(keys []string) *KeyPairEncrypt {
	return &KeyPairEncrypt{
		Keys: keys,
	}
}

func NewKeyPairAESEncrypt(keys []string) *KeyPairEncrypt {
	return &KeyPairEncrypt{
		Keys:   keys,
		useAES: true,
	}
}

/// 第一个返回值为加密的密文，第二个返回值为时间戳
func (enc *KeyPairEncrypt) Encrypt(keySeed int, text string) (string, int64, error) {
	if len(enc.Keys) == 0 {
		return "", 0, fmt.Errorf("didnot have any keys")
	}

	if len(text) == 0 {
		return "", 0, fmt.Errorf("text is empty")
	}

	timestamp := time.Now().Unix()
	key, keySeedNew := enc.GetKey(keySeed)
	insertPos := (keySeedNew >> 2) % len(text)
	text = text[0:insertPos] + fmt.Sprintf("%d", timestamp) + text[insertPos:]

	var encryptBytes []byte
	var err error
	if enc.useAES {
		encryptBytes, err = AESEncryptCFB([]byte(text), key)
	} else {
		encryptBytes, err = TriDESEncrypt([]byte(text), key)
	}
	if err != nil {
		return "", 0, err
	}

	return base64.StdEncoding.EncodeToString(encryptBytes), timestamp, nil
}

/// 获取密钥
func (enc *KeyPairEncrypt) GetKey(keySeed int) ([]byte, int) {
	keySeed += 7
	if keySeed < 0 {
		keySeed = -keySeed + 5
	}
	keySeed = keySeed % len(enc.Keys)

	key := make([]byte, 24)
	for i := 0; i < 24; i++ {
		keySeed += 3
		row := keySeed % len(enc.Keys)
		keySeed += 5
		col := keySeed % len(enc.Keys[row])

		key[i] = byte((int(enc.Keys[row][col]) + 7) % 256)
	}
	return key, keySeed
}

/// 解密
func (enc *KeyPairEncrypt) Decrypt(keySeed int, timestamp int64, validSecs int, encryptText string) (string, error) {
	/// 判断时间超出了要求的时间长度
	if validSecs <= 0 {
		return "", fmt.Errorf("error valid seconds")
	}

	curTimestamp := time.Now().Unix()
	if timestamp+int64(validSecs) < curTimestamp {
		return "", fmt.Errorf("error timestamp")
	}

	/// 解密内容
	if len(enc.Keys) == 0 {
		return "", fmt.Errorf("didnot have any keys")
	}

	if len(encryptText) == 0{
		return "",fmt.Errorf("text is empty")
	}
	key,keySeedNew := enc.GetKey(keySeed)
	timestampStr := fmt.Sprintf("%d",timestamp)
	encryptBytes, err := base64.StdEncoding.DecodeString(encryptText)
	if err != nil {
		return "", fmt.Errorf("base64 decode error")
	}

	var plainBytes []byte

	if enc.useAES {
		plainBytes, err = AESDecryptCFB(encryptBytes, key)
	} else {
		plainBytes, err = TriDESDecrypt(encryptBytes, key)
	}
	if err != nil {
		return "", fmt.Errorf("decrypt error")
	}

	plainText := string(plainBytes)

	if len(plainText) <= len(timestampStr) {
		return "", fmt.Errorf("plain error")
	}
	insertPos := (keySeedNew >> 2) % (len(plainText) - len(timestampStr))
	lastPos := insertPos + len(timestampStr)

	if len(plainText) < lastPos {
		return "", fmt.Errorf("plain length error")
	}
	if plainText[insertPos:lastPos] != timestampStr {
		return "", fmt.Errorf("timestamp cannot match.")
	}

	return plainText[:insertPos] + plainText[lastPos:], nil
}