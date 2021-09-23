package util

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
)

//获取文件地址
func GetFilePath(fileMd5 string) string {
	fileHashDir := HashMakeDir(fileMd5) //获取文件存储路径
	return fileHashDir + "/" + fileMd5
}

//获取文件存放目录 目录打散hash算法
func HashMakeDir(fileMd5 string) (fileDir string) {
	var hashNum int
	hashNum = int(crc32.ChecksumIEEE([]byte(fileMd5)))
	if -hashNum >= 0 {
		hashNum = -hashNum
	}
	dir := strconv.Itoa(hashNum & 0xf)           //生成1～16的数字
	subDir := strconv.Itoa((hashNum & 0xf) >> 4) //得到名为1~16的下下及文件夹
	return dir + "/" + subDir
}

//创建目录
func MkdirAll(filePath string) (err error) {
	//如果文件夹不存在 则不创建
	if _, mkdirErr := os.Stat(filePath); os.IsNotExist(mkdirErr) {
		// 必须分成两步：先创建文件夹、再修改权限
		if err = os.MkdirAll(filePath, os.ModePerm); err == nil {
			err := os.Chmod(filePath, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return
	}
	return
}

//获取文件MD5值
func GetFileMd5(file *multipart.FileHeader) (md5Str string) {
	src, _ := file.Open()
	defer src.Close()
	md5h := md5.New()
	_, err := io.Copy(md5h, src)
	if err != nil {
		return ""
	}
	md5Str = fmt.Sprintf("%x", md5h.Sum([]byte(""))) //md5
	return
}

func CheckFileType(fileType, fileSuffix string) bool {
	fileSuffixList := strings.Split(fileSuffix, ",")
	fileType = strings.Replace(fileType, ".", "", 1)
	for _, val := range fileSuffixList {
		if fileType == val {
			return true
		}
	}
	return false
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func GetScaleImgPath(path, width, height, quality string) string {
	return path + "_scale_" + width + "*" + height + "_" + quality
}
