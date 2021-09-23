package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-file-system/library"
	"go-file-system/util"
	"net/http"
	"net/http/httputil"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

const (
	//文件存放路径
	FilePath = "./file"
	//可上传文件类型
	FileType = "bmp,jpeg,png,gif"
	//设置服务端口
	Port = 9080
	//文件访问path
	StaticPath = "static"
)

var (
	FlagPort     int    //端口
	FlagFilePath string //文件存放路径
)

//接口返回结构
type ResultStruct struct {
	Code int         `json:"code"` // 0-正常 非0 异常
	Data interface{} `json:"data"` // 返回数据
	Msg  string      `json:"msg"`  //返回信息
}

//文件信息
type FileInfo struct {
	FileName string `json:"file_name"` //文件名称
	FileKey  string `json:"file_key"`  //文件唯一key用作访问文件
}

func main() {
	router := gin.Default()

	//初始化参数
	initFlag()

	//上传示例
	router.Static("/demo", "./demo")

	router.Static("/"+StaticPath, FlagFilePath)

	//文件上传
	router.POST("/upload", func(c *gin.Context) {
		fileInfoList := make([]FileInfo, 0)
		result := ResultStruct{
			Data: new(interface{}),
		}
		// Multipart form
		form, err := c.MultipartForm()
		if err != nil {
			result.Code = -1
			result.Msg = err.Error()
			c.JSON(http.StatusOK, result)
			return
		}
		files := form.File["files"]
		if len(files) == 0 {
			result.Code = -1
			result.Msg = "必须传文件"
			c.JSON(http.StatusOK, result)
			return
		}
		for _, file := range files {
			filename := filepath.Base(file.Filename) //获取文件名称

			fileSuffix := path.Ext(filename)
			//判断可上传文件类型
			if !util.CheckFileType(fileSuffix, FileType) {
				result.Code = -1
				result.Msg = "非法文件类型"
				c.JSON(http.StatusOK, result)
				return
			}

			fileMd5 := util.GetFileMd5(file) //获取文件MD5值
			fileInfoList = append(fileInfoList, FileInfo{FileName: filename, FileKey: fileMd5})

			fileHashDir := util.HashMakeDir(fileMd5) //获取文件存储路径

			fileDir := FlagFilePath + "/" + fileHashDir
			err = util.MkdirAll(fileDir) //生成文件目录
			if err != nil {
				result.Code = -1
				result.Msg = err.Error()
				c.JSON(http.StatusOK, result)
				return
			}

			filePath := FlagFilePath + "/" + util.GetFilePath(fileMd5)

			//文件是否存在
			exists, _ := util.PathExists(util.GetFilePath(fileMd5))
			if exists {
				fmt.Println(" file exist: ", filePath)
				continue
			}

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				result.Code = -1
				result.Msg = err.Error()
				c.JSON(http.StatusOK, result)
				return
			}
		}

		result.Data = fileInfoList
		result.Msg = "success"
		c.JSON(http.StatusOK, result)
	})

	//访问文件支持所有文件
	//根据文件md5值判断文件路径
	//如果不存在则404
	//如果存在则转发
	router.GET("/file/:file_key", func(c *gin.Context) {
		fileMd5 := c.Param("file_key")
		filePath := util.GetFilePath(fileMd5)
		//文件是否存在
		exists, _ := util.PathExists(FlagFilePath + "/" + filePath)
		if !exists {
			c.Redirect(http.StatusNotFound, "")
			return
		}
		//转发
		reverseProxy(c, filePath)
		return
	})

	//原始图片访问
	router.GET("/img/original/:file_key", func(c *gin.Context) {
		fileMd5 := c.Param("file_key")
		filePath := util.GetFilePath(fileMd5)
		//文件是否存在
		exists, _ := util.PathExists(FlagFilePath + "/" + filePath)
		if !exists {
			c.Redirect(http.StatusNotFound, "")
			return
		}
		//转发
		reverseProxy(c, filePath)
		return
	})

	//压缩图片访问
	router.GET("/img/scale/:file_key/:width/:height/:quality", func(c *gin.Context) {
		fileMd5 := c.Param("file_key")
		width := c.Param("width")
		height := c.Param("height")
		quality := c.Param("quality")
		filePath := util.GetFilePath(fileMd5)
		//文件是否存在
		p := filePath
		newP := util.GetScaleImgPath(p, width, height, quality)
		exists, _ := util.PathExists(FlagFilePath + "/" + newP)
		if !exists {
			fIn, err := os.Open(FlagFilePath + "/" + p)
			if err != nil {
				panic(err)
			}
			defer fIn.Close()
			fOut, _ := os.Create(FlagFilePath + "/" + newP)
			defer fOut.Close()
			widthInt, _ := strconv.Atoi(width)
			heightInt, _ := strconv.Atoi(height)
			qualityInt, _ := strconv.Atoi(quality)
			err = library.Scale(fIn, fOut, widthInt, heightInt, qualityInt)
			if err != nil {
				panic(err)
			}
		}
		//转发
		reverseProxy(c, newP)
		return
	})

	//启动服务
	err := router.Run(":" + strconv.Itoa(FlagPort))
	if err != nil {
		fmt.Println(" server err: ", err.Error())
		return
	}
}

// go run main.go -f ./file -p 9081

func initFlag() {
	flag.StringVar(&FlagFilePath, "f", FilePath, "文件存储路径,默认:"+FilePath)
	flag.IntVar(&FlagPort, "p", Port, "端口,默认:"+strconv.Itoa(Port))
	//从arguments中解析注册的flag
	flag.Parse()
}

//转发
func reverseProxy(c *gin.Context, filePath string) {
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL.Scheme = "http"
			request.URL.Host = "127.0.0.1:" + strconv.Itoa(FlagPort)
			request.URL.Path = "/" + StaticPath + "/" + filePath
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
	c.Abort()
}
