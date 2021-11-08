package storage_util

import (
	"fmt"
	"io"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

// 写入文件,文件不存在则创建,如在则追加内容
func WriteFile(path string, str string) error {
	_, b := IsFile(path)
	var f *os.File
	var err error

	if b {
		if err = os.Remove(path); err != nil {
			return err
		}
	}

	f, err = os.Create(path)
	if err != nil {
		return err
	}

	//使用完毕，需要关闭文件
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println("err = ", err)
		}
	}()

	_, err = f.WriteString(str)
	return err
}

// 判断路径是否存在
func IsExists(path string) (os.FileInfo, bool) {
	f, err := os.Stat(path)
	return f, err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func IsDir(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && f.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) (os.FileInfo, bool) {
	f, flag := IsExists(path)
	return f, flag && !f.IsDir()
}

func main() {
	path := "./demo.txt"
	str := "abcd\r\nefg我爱学习"
	WriteFile(path, str)
}
