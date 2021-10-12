package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/one-gold-coin/getenv"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

func upload(ctx *gin.Context) {
	fileInfoList := make([]FileInfo, 0)
	result := ResultStruct{
		Data: new(interface{}),
	}
	// Multipart form
	form, err := ctx.MultipartForm()
	if err != nil {
		result.Code = -1
		result.Msg = err.Error()
		ctx.JSON(http.StatusOK, result)
		return
	}
	files := form.File["files"]
	if len(files) == 0 {
		result.Code = -1
		result.Msg = "必须传文件"
		ctx.JSON(http.StatusOK, result)
		return
	}
	for _, file := range files {
		filename := filepath.Base(file.Filename) //获取文件名称

		fileSuffix := path.Ext(filename)
		//判断可上传文件类型
		if !CheckFileType(fileSuffix, getenv.GetVal("FILE_TYPE").String()) {
			result.Code = -1
			result.Msg = "非法文件类型"
			ctx.JSON(http.StatusOK, result)
			return
		}

		fileMd5 := GetFileMd5(file) //获取文件MD5值
		fileInfoList = append(fileInfoList, FileInfo{FileName: filename, FileKey: fileMd5})

		fileHashDir := HashMakeDir(fileMd5) //获取文件存储路径

		fileDir := getenv.GetVal("FILE_PATH").String() + "/" + fileHashDir
		err = MkdirAll(fileDir) //生成文件目录
		if err != nil {
			result.Code = -1
			result.Msg = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}

		filePath := getenv.GetVal("FILE_PATH").String() + "/" + GetFilePath(fileMd5)

		//文件是否存在
		exists, _ := PathExists(GetFilePath(fileMd5))
		if exists {
			fmt.Println(" file exist: ", filePath)
			continue
		}

		if err := ctx.SaveUploadedFile(file, filePath); err != nil {
			result.Code = -1
			result.Msg = err.Error()
			ctx.JSON(http.StatusOK, result)
			return
		}
	}

	result.Data = fileInfoList
	result.Msg = "success"
	ctx.JSON(http.StatusOK, result)
	return
}

//访问文件支持所有文件
//根据文件md5值判断文件路径
//如果不存在则404
//如果存在则转发
func file(ctx *gin.Context) {
	fileMd5 := ctx.Param("file_key")
	if fileMd5 == "" {
		ctx.Redirect(http.StatusNotFound, "")
		return
	}

	filePath := GetFilePath(fileMd5)
	//文件是否存在
	exists, _ := PathExists(getenv.GetVal("FILE_PATH").String() + "/" + filePath)
	if !exists {
		ctx.Redirect(http.StatusNotFound, "")
		return
	}
	//转发
	reverseProxy(ctx, filePath)
	return
}

func imgOriginal(ctx *gin.Context) {
	fileMd5 := ctx.Param("file_key")
	if fileMd5 == "" {
		ctx.Redirect(http.StatusNotFound, "")
		return
	}
	filePath := GetFilePath(fileMd5)
	//文件是否存在
	exists, _ := PathExists(getenv.GetVal("FILE_PATH").String() + "/" + filePath)
	if !exists {
		ctx.Redirect(http.StatusNotFound, "")
		return
	}
	//转发
	reverseProxy(ctx, filePath)
	return
}

func imgScale(ctx *gin.Context) {
	fileMd5 := ctx.Param("file_key")
	if fileMd5 == "" {
		ctx.Redirect(http.StatusNotFound, "")
		return
	}
	width := ctx.Param("width")
	height := ctx.Param("height")
	quality := ctx.Param("quality")
	filePath := GetFilePath(fileMd5)
	//文件是否存在
	p := filePath
	newP := GetScaleImgPath(p, width, height, quality)
	exists, _ := PathExists(getenv.GetVal("FILE_PATH").String() + "/" + newP)
	if !exists {
		fIn, err := os.Open(getenv.GetVal("FILE_PATH").String() + "/" + p)
		if err != nil {
			panic(err)
		}
		defer fIn.Close()
		fOut, _ := os.Create(getenv.GetVal("FILE_PATH").String() + "/" + newP)
		defer fOut.Close()
		widthInt, _ := strconv.Atoi(width)
		heightInt, _ := strconv.Atoi(height)
		qualityInt, _ := strconv.Atoi(quality)
		err = Scale(fIn, fOut, widthInt, heightInt, qualityInt)
		if err != nil {
			panic(err)
		}
	}
	//转发
	reverseProxy(ctx, newP)
	return
}
