package util

import (
	"fmt"
	// "os/exec"

	"os"
	"strings"
)

func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func ReadFile(path string) string {
	contents, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	result := strings.Replace(string(contents), "\n", "", 1)
	return string(result)
}

func RemoveFile(fileName string) {
	if Exists(fileName) {
		err := os.Remove(fileName)
		if err != nil {
			fmt.Printf("remove file %s error %s ", fileName, err.Error())
		}
	} else {
		fmt.Printf("No file[%s] to remove ", fileName)
	}
}

func RemoveDir(dirName string) error {
	if Exists(dirName) {
		return os.RemoveAll(dirName)
	}
	return nil

}

func MakeDir(path string) error {
	if !Exists(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// func ReadFile(path string) string {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}
// 	defer file.Close()

// 	fileinfo, err := file.Stat()
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}

// 	filesize := fileinfo.Size()
// 	buffer := make([]byte, filesize)

// 	_, err = file.Read(buffer)
// 	if err != nil {
// 		fmt.Println(err)
// 		return ""
// 	}

// 	fmt.Printf("read str: %s", string(buffer))
// 	return string(buffer)
// }
