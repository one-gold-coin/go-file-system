# go-file-system
Golang 简单文件系统

通过目录打散hash算法，让文件相对平均到每个文件夹

目前支持上传文件类型：bmp,jpeg,png,gif

如果需要上传其他文件类型，自行修改

项目自支持文件转发机制，访问文件是，可以修改成Nginx

# 启动项目
-f 指定配置文件


```text
go run *.go -f .env.example
```

# 示例

1、上传示例
```text
127.0.0.1:9001/demo
```

# 接口
```text
1、上传接口：127.0.0.1:9001/upload

接口返回：
{"code":0,"data":[{"file_name":"golang.jpeg","file_key":"f187d470506a19c3f0bb54fbb84ee6f2"}],"msg":"success"}

2、访问文件接口：127.0.0.1:9001/file/:file_key
file_key：文件key；接口1返回

访问地址：127.0.0.1:9001/file/f187d470506a19c3f0bb54fbb84ee6f2

3、图片类型，原始图片访问：127.0.0.1:9001/img/:file_key
file_key：文件key；接口1返回

访问地址：127.0.0.1:9001/img/f187d470506a19c3f0bb54fbb84ee6f2

4、图片类型、压缩图片访问，可自定义尺寸：127.0.0.1:9001/img/:file_key/scale/:width/:height/:quality
file_key：文件key；接口1返回
width: 图片长度，单位像素
height: 图片高度，单位像素
quality: 图片质量，0～100指定，传0则为100

访问地址：127.0.0.1:9001/img/f187d470506a19c3f0bb54fbb84ee6f2/scale/200/200/0

```