package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	IMAGESFOLDER     = "source\\_images"     // 图片所在文件夹名
	IMAGESFOLDERDONE = "source\\_imagesDONE" // 处理完的文件要放去的文件夹名
	PTHSEP           = string(os.PathSeparator)
)

var err error

func main() {
	for i := 0; i < 20; i++ {
		e := Movepic()
		fmt.Println("err in main:", e, "i =", i)
		time.Sleep(1e9)
	}

}

func Movepic() error {

	// 获取文件夹路径
	tempPath, _ := os.Getwd()
	imagesFolderPath := tempPath + PTHSEP + IMAGESFOLDER
	imagesFolderDonePath := tempPath + PTHSEP + IMAGESFOLDERDONE

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
		fmt.Printf("不存在转移图片的目标文件夹，但已成功创建！路径：%s", imagesFolderDonePath)
	}

	// 读取目录下文件列表
	images, err := ioutil.ReadDir(imagesFolderPath)
	if err != nil {
		return err
	}

	_ = images
	for i := 0; i < len(images); {
		if strings.HasSuffix(strings.ToLower(images[i].Name()), ".jpg") ||
			strings.HasSuffix(strings.ToLower(images[i].Name()), ".png") ||
			strings.HasSuffix(strings.ToLower(images[i].Name()), ".jpeg") ||
			strings.HasSuffix(strings.ToLower(images[i].Name()), ".gif") {
			i++
		} else {
			images = append(images[:i], images[i+1:]...)
		}
	}
	if len(images) == 0 {
		fmt.Printf("\"%s\"文件夹中未找到 .png/.jpg/.jpeg/.gif 后缀的图片！\n", IMAGESFOLDER)
		err = errors.New("图片库文件夹中未发现可用图片！")
		return err
	}

	// 把第一张图片改名，并转移到DONE文件夹
	renameImage(imagesFolderPath, imagesFolderDonePath, images[0].Name())
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

func renameImage(folderPath string, tofolderPath string, imageName string) {
	oldpath := folderPath + PTHSEP + imageName
	var newpath string
	switch {
	case strings.HasSuffix(strings.ToLower(imageName), ".png"):
		newpath = tofolderPath + PTHSEP + time.Now().Format("2006-01-02-15-04-05") + ".png"
	case strings.HasSuffix(strings.ToLower(imageName), ".jpg"):
		newpath = tofolderPath + PTHSEP + time.Now().Format("2006-01-02-15-04-05") + ".jpg"
	case strings.HasSuffix(strings.ToLower(imageName), ".jpeg"):
		newpath = tofolderPath + PTHSEP + time.Now().Format("2006-01-02-15-04-05") + ".jpeg"
	case strings.HasSuffix(strings.ToLower(imageName), ".gif"):
		newpath = tofolderPath + PTHSEP + time.Now().Format("2006-01-02-15-04-05") + ".gif"
	}
	err := os.Rename(oldpath, newpath)
	if err != nil {
		fmt.Println("图片重命名并转移时出错：", err)
	}
	fmt.Printf("图片\"%s\"重命名及转移成功!新文件路径：%s\n", imageName, newpath)
	return
}
