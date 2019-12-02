package aesutil

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "encoding/base64"
    "strings"
)

const key = `qimin_api_v1_key`

func Decrypt(crypted string) (string, error) {


    defer func() {
        recover()
    }()


    crypted = strings.Replace(crypted, "\r\n", "", -1)
    crypted = strings.Replace(crypted, "%2b", "+", -1)

    decode, _ := base64.StdEncoding.DecodeString(crypted)

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return crypted, err
    }
    blockMode := NewECBDecrypter(block)
    origData := make([]byte, len(decode))
    blockMode.CryptBlocks(origData, decode)
    origData = PKCS5UnPadding(origData)

    return string(origData), nil
}

func Encrypt(src string) string {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return src
    }
    if src == "" {
        return src
    }
    ecb := NewECBEncrypter(block)
    content := []byte(src)
    content = PKCS5Padding(content, block.BlockSize())
    crypted := make([]byte, len(content))
    ecb.CryptBlocks(crypted, content)

    result := base64.StdEncoding.EncodeToString(crypted)
    result = strings.Replace(result, "\r\n", "", -1)
    result = strings.Replace(result, "+", "%2b", -1)

    return result
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
    padding := blockSize - len(ciphertext)%blockSize
    text := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, text...)
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

type ecb struct {
    b         cipher.Block
    blockSize int
}

func newECB(b cipher.Block) *ecb {
    return &ecb{
        b:         b,
        blockSize: b.BlockSize(),
    }
}

type ecbEncrypter ecb

func NewECBEncrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbEncrypter)(newECB(b))
}
func (x *ecbEncrypter) BlockSize() int { return x.blockSize }
func (x *ecbEncrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Encrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}

type ecbDecrypter ecb

func NewECBDecrypter(b cipher.Block) cipher.BlockMode {
    return (*ecbDecrypter)(newECB(b))
}
func (x *ecbDecrypter) BlockSize() int { return x.blockSize }
func (x *ecbDecrypter) CryptBlocks(dst, src []byte) {
    if len(src)%x.blockSize != 0 {
        panic("crypto/cipher: input not full blocks")
    }
    if len(dst) < len(src) {
        panic("crypto/cipher: output smaller than input")
    }
    for len(src) > 0 {
        x.b.Decrypt(dst, src[:x.blockSize])
        src = src[x.blockSize:]
        dst = dst[x.blockSize:]
    }
}
