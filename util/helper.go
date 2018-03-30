package util

import (
	"fmt"
	"io"
	"bytes"
    "crypto/rand"
	"compress/zlib"
)

const (
    letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
)

//随机产生一个字符串，长度自定义
func GenerateRandomString(n int) (string, error) {
    bytes, err := generateRandomBytes(n)
    if err != nil {
        return "", err
    }
    length := byte(len(letters))
    for i, b := range bytes {
        bytes[i] = letters[b%length]
    }
    return string(bytes), nil
}

func Compress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	defer w.Close()
	w.Write(data)
	return in.Bytes()
}

func Decompress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		fmt.Printf("error decompress: %v\n", err)
		return []byte{}
	}
	var out bytes.Buffer
	io.Copy(&out, r)
	return out.Bytes()
}

//=======================================================
func generateRandomBytes(n int) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    // Note that err == nil only if we read len(b) bytes.
    if err != nil {
        return nil, err
    }

    return b, nil
}
