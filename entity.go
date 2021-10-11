package main

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
