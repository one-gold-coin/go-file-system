package main


const (
	//运行环境
	AppEnv = "prod"
	//文件存放路径
	FilePath = "./file"
	//可上传文件类型
	FileType = "bmp,jpeg,png,gif"
	//设置服务端口
	Port = 9080
	//文件访问path
	StaticPath = "static"
	//图片访问代理
	StaticProxyHost = "127.0.0.1:9080"
)

var (
	//设置服务端口
	FlagPort int
	//文件存放路径
	FlagFilePath string
	//可上传文件类型
	FlagFileType string
	//文件访问path
	FlagStaticPath string
	//图片访问代理
	FlagStaticProxyHost string
	//运行环境
	FlagAppEnv string
)

