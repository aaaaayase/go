package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
)

const SecretKey = "11111"

// HmacSHA1Encrypt 使用 HMAC-SHA1 签名方法对 encryptText 进行签名
func HmacSHA1Encrypt(encryptText, encryptKey string) (string, error) {
	key := []byte(encryptKey)
	h := hmac.New(sha1.New, key)
	_, err := h.Write([]byte(encryptText))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// HmacSHA1EncryptBytes 使用 HMAC-SHA1 签名方法对 encryptData 进行签名
func HmacSHA1EncryptBytes(encryptData []byte, encryptKey string) (string, error) {
	key := []byte(encryptKey)
	h := hmac.New(sha1.New, key)
	_, err := h.Write(encryptData)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
