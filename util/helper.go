package util

import (
	"fmt"
	"io"
	"os"
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
	w.Write(data)
	w.Close()	//flush数据
	return in.Bytes()
}

func Decompress(data []byte) []byte {
	if len(data) == 0 {
		return []byte{}
	}
	fmt.Println("1")
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		fmt.Printf("error decompress: %v\n", err)
		return []byte{}
	}
	defer r.Close()
	var out bytes.Buffer
	io.Copy(&out, r)
	fmt.Println("2")
	return out.Bytes()
}

func Decompress2(data []byte) {
	if len(data) == 0 {
		return
	}
	b := bytes.NewReader(data)
	r, err := zlib.NewReader(b)
	if err != nil {
		fmt.Printf("error decompress: %v\n", err)
		return
	}
	defer r.Close()
	io.Copy(os.Stdout, r)
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
