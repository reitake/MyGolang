package main

import (
	"errors"
	"fmt"
	"os"
	"time"
)

const (
	IMAGESFOLDER     = "source/_images"     // 图片所在文件夹名
	IMAGESFOLDERDONE = "source/_imagesDONE" // 处理完的文件要放去的文件夹名
)

var err error

func main() {
	e := movepic()
	fmt.Println("err in main:", e)
}

func movepic() error {

	// 获取文件夹路径
	tempPath, _ := os.Getwd()
	imagesFolderPath := tempPath + "./" + IMAGESFOLDER
	imagesFolderDonePath := tempPath + "./" + IMAGESFOLDERDONE

	// 判断是否存在放图片的文件夹
	// 若不存在或出错，则返回错误
	// 若存在，则进入文件夹
	if exist, err := isPathExists(imagesFolderPath); err != nil {
		fmt.Println("Error:", err)
		return err
	} else if !exist {
		fmt.Printf("未找到存放图片的指定文件夹：\"%s\"\n", IMAGESFOLDER)
		err = errors.New("存放原图片的文件夹不存在")
		return err
	}

	// 判断是否存在转移图片的文件夹
	// 若不存在，则创建文件夹
	if exist, err := isPathExists(imagesFolderDonePath); err != nil {
		fmt.Println("Error:", err)
		return err
	} else if !exist {
		err = os.Mkdir(imagesFolderDonePath, 0666)
		if err != nil {
			return err
		}
		fmt.Println("不存在转移图片的目标文件夹，但已成功创建！")
	}

	file, _ := os.Open(imagesFolderPath)
	fmt.Println(file)

	filenameto := time.Now().Format("2006-01-02-15:04:05")
	fmt.Println(filenameto)
	return nil
}

func isPathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
